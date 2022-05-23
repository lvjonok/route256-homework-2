package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) GetImage(ctx context.Context, id models.ID) (*models.Image, error) {
	const query = `select id, image, href from images where id=$1;`

	var image models.Image
	err := c.pool.QueryRow(ctx, query, id).Scan(&image.ID, &image.Content, &image.Href)

	if err != nil && err == pgx.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get image, err: <%v>", err)
	}

	return &image, nil
}

func (c *Client) GetImageByHref(ctx context.Context, href string) (*models.Image, error) {
	const queryExisting = `SELECT id, image, href from images WHERE href=$1 ORDER BY created_at DESC LIMIT 1;`

	var image models.Image
	err := c.pool.QueryRow(ctx, queryExisting, href).Scan(&image.ID, &image.Content, &image.Href)

	if err != nil && err == pgx.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get image by href, err: <%v>", err)
	}

	return &image, nil
}
