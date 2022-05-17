package app_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/app"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
	homework_2 "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
)

func TestGetStat(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetStatMock.Return(&models.Statistics{
		Stat: []models.TaskStat{
			{
				TaskNumber: 1,
				Correct:    1,
				All:        2,
			},
			{
				TaskNumber: 2,
				Correct:    10,
				All:        20,
			},
		},
	}, nil)
	srv := app.New(mockDB)

	ctx := context.Background()
	resp, err := srv.GetStat(ctx, &homework_2.GetStatRequest{ChatId: 100})
	require.NoError(t, err)

	exp := []*homework_2.TaskStat{
		{TaskNumber: 1, Correct: 1, All: 2},
		{TaskNumber: 2, Correct: 10, All: 20},
	}

	require.True(t, reflect.DeepEqual(exp, resp.Stat))
}
