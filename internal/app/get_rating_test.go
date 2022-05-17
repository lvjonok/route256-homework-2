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

func TestGetRating(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetRatingMock.Return(&models.Rating{Position: 10, All: 20}, nil)
	srv := app.New(mockDB)

	ctx := context.Background()
	resp, err := srv.GetRating(ctx, &homework_2.GetRatingRequest{ChatId: 100})
	require.NoError(t, err)

	require.Equal(t, int64(10), resp.Position)
	require.Equal(t, int64(20), resp.All)
}
