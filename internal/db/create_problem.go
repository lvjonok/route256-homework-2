package db

import (
	"context"

	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c client) CreateProblem(ctx context.Context, problem m.Problem) error {

	const query = `insert into problems(id, category_id, task_number, image, parts, answer) 
	VALUES ($1, $2, $3, $4, $5, $6);`

	_, err := c.pool.Exec(ctx, query,
		problem.ProblemID,
		problem.CategoryID,
		problem.TaskNumber,
		problem.ProblemImage,
		problem.Parts,
		problem.Answer,
	)

	return err
}
