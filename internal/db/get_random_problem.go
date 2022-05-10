package db

import (
	"context"
	"fmt"

	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c client) GetRandomProblem(ctx context.Context) (*m.Problem, error) {
	const query = `SELECT id, category_id, task_number, image, parts, answer
		FROM problems
		ORDER BY RANDOM()
		LIMIT 1;`

	var out m.Problem

	if err := c.pool.QueryRow(ctx, query).Scan(
		&out.ProblemID,
		&out.CategoryID,
		&out.TaskNumber,
		&out.ProblemImage,
		&out.Parts,
		&out.Answer,
	); err != nil {
		return nil, fmt.Errorf("failed to query random problem, err: %v", err)
	}

	return &out, nil
}
