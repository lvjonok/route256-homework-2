package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) CreateCategory(ctx context.Context, cat m.Category) (*m.ID, error) {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction, err: <%v>", err)
	}

	var newID m.ID
	var oldCat m.Category

	const queryExisting = `SELECT id, task_number, title 
		FROM categories 
		WHERE category_id=$1 
		ORDER BY updated_at DESC 
		LIMIT 1;`
	err = tx.QueryRow(ctx, queryExisting, cat.CategoryID).
		Scan(&newID, &oldCat.TaskNumber, &oldCat.Title)
	if err != nil && err != pgx.ErrNoRows {
		if rerr := tx.Rollback(ctx); rerr != nil {
			return nil, fmt.Errorf("failed to query existing category, err: <%v>, failed to rollback: <%v>", err, rerr)
		}
		return nil, fmt.Errorf("failed to query existing category, err: <%v>", err)
	}

	if err == pgx.ErrNoRows || (oldCat.TaskNumber != cat.TaskNumber || oldCat.Title != cat.Title) {
		const query = `INSERT INTO categories(category_id, task_number, title) 
			VALUES ($1, $2, $3) RETURNING id;`

		err := c.pool.QueryRow(ctx, query,
			cat.CategoryID,
			cat.TaskNumber,
			cat.Title,
		).Scan(&newID)
		if err != nil {
			return nil, fmt.Errorf("failed to add new category, err: <%v>", err)
		}
	}

	err = tx.Commit(ctx)

	return &newID, err
}
