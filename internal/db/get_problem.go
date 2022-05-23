package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) GetProblem(ctx context.Context, id m.ID) (*m.Problem, error) {
	const query = `select id, problem_id, category_id, image, parts, answer from problems where id=$1`

	p := m.Problem{}

	err := c.pool.QueryRow(ctx, query, id).
		Scan(&p.ID, &p.ProblemID, &p.CategoryID, &p.ProblemImage, &p.Parts, &p.Answer)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get problem from database, err: <%v>", err)
	}

	return &p, nil
}

func (c *Client) GetProblemByProblemID(ctx context.Context, problemID m.ID) (*m.Problem, error) {
	const queryExisting = `SELECT id, problem_id, category_id, image, parts, answer
		FROM problems
		WHERE problem_id=$1
		ORDER BY updated_at DESC
		LIMIT 1;`
	p := m.Problem{}

	err := c.pool.QueryRow(ctx, queryExisting, problemID).
		Scan(&p.ID, &p.ProblemID, &p.CategoryID, &p.ProblemImage, &p.Parts, &p.Answer)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get problem from database, err: <%v>", err)
	}

	return &p, nil
}
