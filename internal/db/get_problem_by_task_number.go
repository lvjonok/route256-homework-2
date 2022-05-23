package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) GetProblemByTaskNumber(ctx context.Context, taskNumber int) (*m.Problem, error) {
	const query = `SELECT p.id, p.problem_id, p.category_id, p.image, p.parts, p.answer
		FROM (SELECT DISTINCT ON (p.problem_id) *
					FROM problems p
					ORDER BY p.problem_id, p.updated_at DESC) AS p
						JOIN categories c ON p.category_id = c.id
		WHERE c.task_number = $1
		ORDER BY RANDOM()
		LIMIT 1;`

	p := m.Problem{}

	err := c.pool.QueryRow(ctx, query, taskNumber).
		Scan(&p.ID, &p.ProblemID, &p.CategoryID, &p.ProblemImage, &p.Parts, &p.Answer)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get problem from database, err: %v", err)
	}
	return &p, nil
}
