package db

import (
	"context"
	"fmt"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) CreateImage(ctx context.Context, raw []byte, href string) (*models.ID, error) {
	var newID models.ID
	const query = `insert into images(image, href) values($1, $2) returning id;`
	err := c.pool.QueryRow(ctx, query, raw, href).Scan(&newID)
	if err != nil {
		return nil, fmt.Errorf("failed to add new image, err: <%v>", err)
	}

	return &newID, err
}
