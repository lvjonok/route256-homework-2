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

func TestGetImage(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetImageMock.Return(&models.Image{Content: []byte{0, 1, 2, 3}}, nil)
	srv := app.New(mockDB)

	ctx := context.Background()
	resp, err := srv.GetImage(ctx, &homework_2.GetImageRequest{ImageId: 1})
	require.NoError(t, err)

	require.True(t, reflect.DeepEqual(resp.Image, []byte{0, 1, 2, 3}))
}
