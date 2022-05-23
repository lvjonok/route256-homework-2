package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/db"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func TestGetImage(t *testing.T) {
	client, ctx := Prepare(t)

	id, err := client.CreateImage(ctx, []byte{1, 2, 3, 4, 5}, "gitlab.com")
	require.NoError(t, err)

	img, err := client.GetImage(ctx, *id)
	require.NoError(t, err)

	require.Equal(t, models.Image{ID: *id, Content: []byte{1, 2, 3, 4, 5}, Href: "gitlab.com"}, *img)
}

func TestGetImageFail(t *testing.T) {
	client, ctx := Prepare(t)

	_, err := client.GetImage(ctx, models.ID(1))
	require.Equal(t, db.ErrNotFound, err)
}
