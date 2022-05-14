package db

import (
	"context"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c client) GetImage(ctx context.Context, id models.ID) ([]byte, error) {
	const query = `select image from images where id=$1;`

	var image []byte
	err := c.pool.QueryRow(ctx, query, id).Scan(&image)

	return image, err
}
