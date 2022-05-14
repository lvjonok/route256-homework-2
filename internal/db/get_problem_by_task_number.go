package db

import (
	"context"
	"fmt"

	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c client) GetProblemByTaskNumber(ctx context.Context, taskNumber int) (*m.Problem, error) {
	const query = `SELECT p.id, p.problem_id, p.category_id, p.image, p.parts, p.answer
		FROM (SELECT DISTINCT *
					FROM problems p
					ORDER BY p.updated_at DESC) AS p
						JOIN categories c ON p.category_id = c.id
		WHERE c.task_number = $1
		ORDER BY RANDOM()
		LIMIT 1;`

	p := m.Problem{}

	err := c.pool.QueryRow(ctx, query, taskNumber).
		Scan(&p.ID, &p.ProblemID, &p.CategoryID, &p.ProblemImage, &p.Parts, &p.Answer)
	if err != nil {
		return nil, fmt.Errorf("failed to get problem from database, err: %v", err)
	}
	return &p, nil
}
