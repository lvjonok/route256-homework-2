package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func TestGetProblem(t *testing.T) {
	client, ctx := Prepare(t)

	// prepare category
	catID, err := client.CreateCategory(ctx, models.Category{CategoryID: 1, TaskNumber: 10, Title: "some title"})
	require.NoError(t, err)

	// prepare problem
	probID, err := client.CreateProblem(ctx, models.Problem{ProblemID: 12345, CategoryID: *catID, Parts: []string{"some cool description"}, Answer: "10"})
	require.NoError(t, err)

	res, err := client.GetProblem(ctx, *probID)
	require.NoError(t, err)

	require.Equal(t, models.ID(12345), res.ProblemID)
	require.Equal(t, []string{"some cool description"}, res.Parts)
	require.Equal(t, "10", res.Answer)
}
