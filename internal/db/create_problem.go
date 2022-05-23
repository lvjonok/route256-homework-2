package db

import (
	"context"
	"fmt"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) CreateProblem(ctx context.Context, problem models.Problem) (*models.ID, error) {

	var newID models.ID

	const query = `INSERT INTO problems(problem_id, category_id, image, parts, answer) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	err := c.pool.QueryRow(ctx, query,
		problem.ProblemID,
		problem.CategoryID,
		problem.ProblemImage,
		problem.Parts,
		problem.Answer,
	).Scan(&newID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert new problem, err: <%v>", err)
	}

	return &newID, err
}
