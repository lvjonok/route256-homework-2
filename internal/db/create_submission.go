package db

import (
	"context"
	"fmt"

	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) CreateSubmission(ctx context.Context, sub m.Submission) (*m.ID, error) {
	const query = `insert into submissions(chat_id, problem_id) 
	VALUES ($1, $2) returning id;`

	var newId m.ID

	err := c.pool.QueryRow(ctx, query,
		sub.ChatID,
		sub.ProblemID,
	).Scan(&newId)
	if err != nil {
		return nil, fmt.Errorf("failed to create submission for user %v, problem %v, err: %v", sub.ChatID, sub.ProblemID, err)
	}

	return &newId, err
}
