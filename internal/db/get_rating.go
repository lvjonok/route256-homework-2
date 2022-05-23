package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) GetRating(ctx context.Context, chatID m.ID) (*m.Rating, error) {
	const query = `SELECT row_number, al.amount
		FROM (SELECT chat_id,
								COUNT(s.chat_id) FILTER ( WHERE s.result = 'correct' ),
								ROW_NUMBER() OVER (ORDER BY COUNT(s.chat_id) FILTER ( WHERE s.result = 'correct' ) DESC )
					FROM submissions s
					GROUP BY chat_id) AS rating,
				(SELECT COUNT(DISTINCT chat_id) AS amount FROM submissions) AS al
		WHERE chat_id = $1;`

	var position int
	var all int

	err := c.pool.QueryRow(ctx, query, chatID).Scan(&position, &all)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get rating from database, err: <%v>", err)
	}

	return &m.Rating{Position: position, All: all}, nil
}
