package db

import (
	"context"
	"fmt"

	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c client) GetProblem(ctx context.Context, problemID m.ID) (*m.Problem, error) {
	const query = `select id, category_id, task_number, image, parts, answer from problems where id=$1`

	p := m.Problem{}

	err := c.pool.QueryRow(ctx, query, problemID).
		Scan(&p.ProblemID, &p.CategoryID, &p.TaskNumber, &p.ProblemImage, &p.Parts, &p.Answer)
	if err != nil {
		return nil, fmt.Errorf("failed to get problem from database, err: %v", err)
	}

	return &p, nil
}