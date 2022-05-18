package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func TestProblemByTaskNumber(t *testing.T) {
	client, ctx := Prepare(t)

	// prepare category
	catID, err := client.CreateCategory(ctx, models.Category{CategoryID: 1, TaskNumber: 10, Title: "some title"})
	require.NoError(t, err)

	// prepare problem
	probID, err := client.CreateProblem(ctx, models.Problem{ProblemID: 12345, CategoryID: *catID, Parts: []string{"some cool description"}, Answer: "10"})
	require.NoError(t, err)

	probID2, err := client.CreateProblem(ctx, models.Problem{ProblemID: 12345, CategoryID: *catID, Parts: []string{"some cool description, but new"}, Answer: "20"})
	require.NoError(t, err)

	res, err := client.GetProblemByTaskNumber(ctx, 10)
	require.NoError(t, err)

	// we should get only updated problem
	require.NotEqual(t, res.ID, *probID)
	require.Equal(t, res.ID, *probID2)
}
