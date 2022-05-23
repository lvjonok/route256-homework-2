package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) GetLastUserSubmission(ctx context.Context, chatID models.ID) (*models.Submission, error) {
	const query = `SELECT s.id, s.chat_id, s.problem_id, s.result
		FROM submissions s
		WHERE s.chat_id = $1 AND s.result='pending'
		ORDER BY updated_at DESC
		LIMIT 1;`

	var sub models.Submission

	err := c.pool.QueryRow(ctx, query, chatID).
		Scan(
			&sub.SubmissionID,
			&sub.ChatID,
			&sub.ProblemID,
			&sub.Result,
		)
	if err == pgx.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("failed to get last user %v submission, err: <%v>", chatID, err)
	}
	return &sub, nil
}
