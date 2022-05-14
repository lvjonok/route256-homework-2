package db

import (
	"bytes"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c client) CreateImage(ctx context.Context, raw []byte, href string) (*models.ID, error) {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction, err: <%v>", err)
	}

	const queryExisting = `SELECT id, image from images WHERE href=$1 ORDER BY created_at DESC LIMIT 1;`

	var newID models.ID
	var originalBytes []byte
	err = tx.QueryRow(ctx, queryExisting, href).Scan(&newID, &originalBytes)
	if err != nil && err != pgx.ErrNoRows {
		if rerr := tx.Rollback(ctx); rerr != nil {
			return nil, fmt.Errorf("failed to query existing images, err: <%v>, failed to rollback: <%v>", err, rerr)
		}
		return nil, fmt.Errorf("failed to query existing images, err: <%v>", err)
	}

	if err == pgx.ErrNoRows || !bytes.Equal(originalBytes, raw) {
		// either we did not find image inside, or it was updated
		const query = `insert into images(image, href) values($1, $2) returning id;`
		err := c.pool.QueryRow(ctx, query, raw, href).Scan(&newID)
		if err != nil {
			return nil, fmt.Errorf("failed to add new image, err: <%v>", err)
		}
	}

	err = tx.Commit(ctx)

	return &newID, err
}
