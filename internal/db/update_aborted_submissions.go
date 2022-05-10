package db

import (
	"context"
	"fmt"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c client) UpdateAbortedSubmissions(ctx context.Context, chatID models.ID) error {
	const query = `UPDATE submissions
		SET result='aborted'
		WHERE chat_id=$1 and result='pending';`

	_, err := c.pool.Exec(ctx, query, chatID)
	if err != nil {
		return fmt.Errorf("failed to get abort submissions user %v submission, err: %v", chatID, err)
	}
	return nil
}
