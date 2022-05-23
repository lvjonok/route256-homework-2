package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/db"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func TestGetCategory(t *testing.T) {
	client, ctx := Prepare(t)

	id, err := client.CreateCategory(ctx, models.Category{CategoryID: 12345, TaskNumber: 12345, Title: "title"})
	require.NoError(t, err)

	cat, err := client.GetCategoryByID(ctx, 12345)
	require.NoError(t, err)

	require.Equal(t, models.Category{ID: *id, CategoryID: 12345, TaskNumber: 12345, Title: "title"}, *cat)
}

func TestGetCategoryNotFound(t *testing.T) {
	client, ctx := Prepare(t)

	_, err := client.GetCategoryByID(ctx, 123456)
	require.Equal(t, db.ErrNotFound, err)
}
