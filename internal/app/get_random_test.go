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

func TestGetRandom(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetProblemByTaskNumberMock.Return(&models.Problem{ProblemID: 12345, Parts: []string{"some part"}}, nil)
	mockDB.UpdateAbortedSubmissionsMock.Return(nil)
	mockDB.CreateSubmissionMock.Return(nil, nil)
	srv := app.New(mockDB)

	ctx := context.Background()
	resp, err := srv.GetRandom(ctx, &homework_2.GetRandomRequest{ChatId: 100})
	require.NoError(t, err)

	require.Equal(t, resp.Problem.ProblemId, int64(12345))
	require.Equal(t, resp.Problem.Description, []string{"some part"})

}
