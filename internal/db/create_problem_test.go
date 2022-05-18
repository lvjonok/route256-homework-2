package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func TestCreateProblem(t *testing.T) {
	client, ctx := Prepare(t)

	// prepare category
	catID, err := client.CreateCategory(ctx, models.Category{CategoryID: 1, TaskNumber: 10, Title: "some title"})
	require.NoError(t, err)

	// test adding first problem
	resID, err := client.CreateProblem(ctx, models.Problem{ProblemID: 12345, CategoryID: *catID, Parts: []string{"some cool description"}, Answer: "10"})
	require.NoError(t, err)

	// test inserting the same problem
	resID2, err := client.CreateProblem(ctx, models.Problem{ProblemID: 12345, CategoryID: *catID, Parts: []string{"some cool description"}, Answer: "10"})
	require.NoError(t, err)

	// they should be the same, because we didn't get anything new to the database
	require.Equal(t, resID, resID2)

	resID3, err := client.CreateProblem(ctx, models.Problem{ProblemID: 12345, CategoryID: *catID, Parts: []string{"some cool description, but changed"}, Answer: "10"})
	require.NoError(t, err)

	// they should not be equal, because we updated parts, and created new entry
	require.NotEqual(t, resID, resID3)
}
