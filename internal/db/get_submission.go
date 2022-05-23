package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) GetSubmission(ctx context.Context, id models.ID) (*models.Submission, error) {
	const query = `SELECT id, chat_id, problem_id, result from submissions where id=$1;`
	var sub models.Submission

	err := c.pool.QueryRow(ctx, query, id).Scan(&sub.SubmissionID, &sub.ChatID, &sub.ProblemID, &sub.Result)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get submission by id, err: <%v>", err)
	}
	return &sub, nil
}
