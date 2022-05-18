package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func TestUpdateSubmission(t *testing.T) {
	client, ctx := Prepare(t)

	// prepare category
	catID, err := client.CreateCategory(ctx, models.Category{CategoryID: 1, TaskNumber: 10, Title: "some title"})
	require.NoError(t, err)

	// prepare problem
	probID, err := client.CreateProblem(ctx, models.Problem{ProblemID: 12345, CategoryID: *catID, Parts: []string{"some cool description"}, Answer: "10"})
	require.NoError(t, err)

	// test new creation of submission
	subID, err := client.CreateSubmission(ctx, models.Submission{ChatID: 12345, ProblemID: *probID})
	require.NoError(t, err)

	sub, err := client.GetSubmission(ctx, *subID)
	require.NoError(t, err)

	sub.Result = models.Wrong
	err = client.UpdateSubmission(ctx, *sub)
	require.NoError(t, err)

	sub, err = client.GetSubmission(ctx, *subID)
	require.NoError(t, err)
	// we just recently updated it to wrong
	require.Equal(t, models.Wrong, sub.Result)

	sub.Result = models.Correct
	err = client.UpdateSubmission(ctx, *sub)
	require.NoError(t, err)

	sub, err = client.GetSubmission(ctx, *subID)
	require.NoError(t, err)
	require.Equal(t, models.Correct, sub.Result)

}
