package app

import (
	"context"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

type DB interface {
	CreateProblem(context.Context, models.Problem) error
	CreateCategory(context.Context, models.Category) error
	CreateSubmission(context.Context, models.Submission) (*models.ID, error)

	GetProblem(context.Context, models.ID) (*models.Problem, error)
	GetProblemByTaskNumber(context.Context, int) (*models.Problem, error)
	GetLastUserSubmission(context.Context, models.ID) (*models.Submission, error)

	UpdateSubmission(context.Context, models.Submission) error
	UpdateAbortedSubmissions(ctx context.Context, chatID models.ID) error

	// TODO: get stat and rating
}
