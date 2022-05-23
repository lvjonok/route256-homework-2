package db

import (
	"context"
	"fmt"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) UpdateSubmission(ctx context.Context, sub models.Submission) error {
	const query = `UPDATE submissions set result=$2 where id=$1;`

	_, err := c.pool.Exec(ctx, query, sub.SubmissionID, sub.Result)
	if err != nil {
		return fmt.Errorf("failed to update submission, err: <%v>", err)
	}
	return nil
}
