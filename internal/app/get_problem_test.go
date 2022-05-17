package app_test

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/app"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
	homework_2 "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
)

func TestGetProblemOutOfRange(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	srv := app.New(mockDB)

	ctx := context.Background()
	_, err := srv.GetProblem(ctx, &homework_2.GetProblemRequest{TaskNumber: 100})
	if err == nil {
		t.Errorf("should be err, task number > 11")
	}

	_, err = srv.GetProblem(ctx, &homework_2.GetProblemRequest{TaskNumber: 0})
	if err == nil {
		t.Errorf("should be err, task number < 1")
	}
}

func TestGetProblem(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetProblemByTaskNumberMock.Return(
		&models.Problem{
			ProblemID: 12345,
			Parts:     []string{"ab as"},
		},
		nil,
	)
	mockDB.UpdateAbortedSubmissionsMock.Return(nil)
	mockDB.CreateSubmissionMock.Return(nil, nil)

	srv := app.New(mockDB)

	ctx := context.Background()
	resp, err := srv.GetProblem(ctx, &homework_2.GetProblemRequest{TaskNumber: 10})
	require.NoError(t, err)
	target := &homework_2.Problem{ProblemId: 12345, Description: []string{"ab as"}}

	require.Equal(t, resp.Problem.ProblemId, target.ProblemId)
	require.Equal(t, resp.Problem.Description, target.Description)
}
