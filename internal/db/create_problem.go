package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) CreateProblem(ctx context.Context, problem models.Problem) (*models.ID, error) {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction, err: <%v>", err)
	}

	const queryExisting = `SELECT id, problem_id, category_id, image, parts, answer
		FROM problems
		WHERE problem_id=$1
		ORDER BY updated_at DESC
		LIMIT 1;`

	var newID models.ID
	var prevProblem models.Problem
	err = tx.QueryRow(ctx, queryExisting, problem.ProblemID).
		Scan(&newID, &prevProblem.ProblemID, &prevProblem.CategoryID, &prevProblem.ProblemImage, &prevProblem.Parts, &prevProblem.Answer)
	if err != nil && err != pgx.ErrNoRows {
		if rerr := tx.Rollback(ctx); rerr != nil {
			return nil, fmt.Errorf("failed to query existing images, err: <%v>, failed to rollback: <%v>", err, rerr)
		}
		return nil, fmt.Errorf("failed to query existing images, err: <%v>", err)
	}

	equalParts := true
	for idx := 0; idx < len(problem.Parts); idx++ {
		if len(prevProblem.Parts) <= idx || problem.Parts[idx] != prevProblem.Parts[idx] {
			equalParts = false
		}
	}
	if len(prevProblem.Parts) != len(problem.Parts) {
		equalParts = false
	}

	if err == pgx.ErrNoRows ||
		prevProblem.CategoryID != problem.CategoryID ||
		!equalParts ||
		prevProblem.Answer != problem.Answer ||
		prevProblem.ProblemImage != problem.ProblemImage {
		// either we did not find problem inside, or it was updated
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
	}

	err = tx.Commit(ctx)

	return &newID, err
}
