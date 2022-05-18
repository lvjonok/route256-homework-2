package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func TestCreateCategory(t *testing.T) {
	client, ctx := Prepare(t)

	// test adding first category
	resID, err := client.CreateCategory(ctx, models.Category{CategoryID: 1, TaskNumber: 1, Title: "some 1st task"})
	require.NoError(t, err)

	// test inserting the same category
	resID2, err := client.CreateCategory(ctx, models.Category{CategoryID: 1, TaskNumber: 1, Title: "some 1st task"})
	require.NoError(t, err)

	// they should be the same, because we didn't get anything new to the database
	require.Equal(t, resID, resID2)

	resID3, err := client.CreateCategory(ctx, models.Category{CategoryID: 1, TaskNumber: 1, Title: "new cool title"})
	require.NoError(t, err)

	// they should not be equal, because we updated title, and created new entry
	require.NotEqual(t, resID, resID3)
}
