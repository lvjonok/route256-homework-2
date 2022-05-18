package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func TestLastUsersSubmission(t *testing.T) {
	client, ctx := Prepare(t)

	// prepare category
	catID, err := client.CreateCategory(ctx, models.Category{CategoryID: 1, TaskNumber: 10, Title: "some title"})
	require.NoError(t, err)

	// prepare problem
	probID, err := client.CreateProblem(ctx, models.Problem{ProblemID: 12345, CategoryID: *catID, Parts: []string{"some cool description"}, Answer: "10"})
	require.NoError(t, err)

	// prepare submission
	subID, err := client.CreateSubmission(ctx, models.Submission{ChatID: 12345, ProblemID: *probID})
	require.NoError(t, err)
	subID2, err := client.CreateSubmission(ctx, models.Submission{ChatID: 12345, ProblemID: *probID})
	require.NoError(t, err)
	subID3, err := client.CreateSubmission(ctx, models.Submission{ChatID: 12345, ProblemID: *probID})
	require.NoError(t, err)

	sub, err := client.GetLastUserSubmission(ctx, 12345)
	require.NoError(t, err)

	require.NotEqual(t, *subID, sub.SubmissionID)
	require.NotEqual(t, *subID2, sub.SubmissionID)
	require.Equal(t, *subID3, sub.SubmissionID)

}
