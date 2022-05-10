package db

import (
	"context"

	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c client) CreateProblem(ctx context.Context, problem m.Problem) error {
	const query = `insert into problems(id, category_id, image, parts, answer) 
	VALUES ($1, $2, $3, $4, $5) ON CONFLICT(id) DO UPDATE set image=$3, parts=$4, answer=$5`

	_, err := c.pool.Exec(ctx, query,
		problem.ProblemID,
		problem.CategoryID,
		problem.ProblemImage,
		problem.Parts,
		problem.Answer,
	)

	return err
}
