package db

import (
	"context"

	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c client) CreateCategory(ctx context.Context, cat m.Category) error {
	const query = `INSERT INTO categories(id, task_number, title)
VALUES ($1, $2, $3)
ON CONFLICT(id) DO UPDATE set task_number=$2, title=$3;`

	_, err := c.pool.Exec(ctx, query,
		cat.CategoryID,
		cat.TaskNumber,
		cat.Title,
	)

	return err
}
