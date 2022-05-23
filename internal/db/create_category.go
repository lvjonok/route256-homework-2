package db

import (
	"context"
	"fmt"

	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) CreateCategory(ctx context.Context, cat m.Category) (*m.ID, error) {
	const query = `INSERT INTO categories(category_id, task_number, title) 
			VALUES ($1, $2, $3) RETURNING id;`

	var newID m.ID

	err := c.pool.QueryRow(ctx, query,
		cat.CategoryID,
		cat.TaskNumber,
		cat.Title,
	).Scan(&newID)
	if err != nil {
		return nil, fmt.Errorf("failed to add new category, err: <%v>", err)
	}

	return &newID, err
}
