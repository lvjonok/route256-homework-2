package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func TestCreateCategory(t *testing.T) {
	client, ctx := Prepare(t)

	// test adding first category
	_, err := client.CreateCategory(ctx, models.Category{CategoryID: 1, TaskNumber: 1, Title: "some 1st task"})
	require.NoError(t, err)
}
