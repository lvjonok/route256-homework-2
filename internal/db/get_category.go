package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) GetCategoryByID(ctx context.Context, categoryID models.ID) (*models.Category, error) {
	var cat models.Category

	const queryExisting = `SELECT id, task_number, category_id, title 
		FROM categories 
		WHERE category_id=$1 
		ORDER BY updated_at DESC 
		LIMIT 1;`
	err := c.pool.QueryRow(ctx, queryExisting, categoryID).
		Scan(&cat.ID, &cat.TaskNumber, &cat.CategoryID, &cat.Title)
	if err != nil && err != pgx.ErrNoRows {
		return nil, fmt.Errorf("failed to query existing category, err: <%v>", err)
	}

	if err == pgx.ErrNoRows {
		return nil, ErrNotFound
	}

	return &cat, nil
}
