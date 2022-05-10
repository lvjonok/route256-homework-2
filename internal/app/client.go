package app

import (
	"context"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

type DB interface {
	CreateProblem(context.Context, models.Problem) error
	GetProblem(context.Context, models.ID) (*models.Problem, error)
	GetProblemByTaskNumber(context.Context, int) (*models.Problem, error)
	GetRandomProblem(context.Context) (*models.Problem, error)
	// CreateSubmission(context.Context, models.Submission)
	// GetSubmission(context.Context, models.ID)
}
