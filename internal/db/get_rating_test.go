package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func TestGetRating(t *testing.T) {
	client, ctx := Prepare(t)

	// prepare category
	catID, err := client.CreateCategory(ctx, models.Category{CategoryID: 1, TaskNumber: 10, Title: "some title"})
	require.NoError(t, err)

	// prepare problem
	probID, err := client.CreateProblem(ctx, models.Problem{ProblemID: 12345, CategoryID: *catID, Parts: []string{"some cool description"}, Answer: "10"})
	require.NoError(t, err)

	// prepare submission
	subID, err := client.CreateSubmission(ctx, models.Submission{ChatID: 1, ProblemID: *probID})
	require.NoError(t, err)
	subID2, err := client.CreateSubmission(ctx, models.Submission{ChatID: 2, ProblemID: *probID})
	require.NoError(t, err)
	subID3, err := client.CreateSubmission(ctx, models.Submission{ChatID: 3, ProblemID: *probID})
	require.NoError(t, err)
	subID4, err := client.CreateSubmission(ctx, models.Submission{ChatID: 1, ProblemID: *probID})
	require.NoError(t, err)
	subID5, err := client.CreateSubmission(ctx, models.Submission{ChatID: 2, ProblemID: *probID})
	require.NoError(t, err)
	subID6, err := client.CreateSubmission(ctx, models.Submission{ChatID: 2, ProblemID: *probID})
	require.NoError(t, err)

	err = client.UpdateSubmission(ctx, models.Submission{SubmissionID: *subID, ChatID: 1, ProblemID: *probID, Result: models.Correct})
	require.NoError(t, err)
	err = client.UpdateSubmission(ctx, models.Submission{SubmissionID: *subID2, ChatID: 2, ProblemID: *probID, Result: models.Correct})
	require.NoError(t, err)
	err = client.UpdateSubmission(ctx, models.Submission{SubmissionID: *subID3, ChatID: 3, ProblemID: *probID, Result: models.Correct})
	require.NoError(t, err)
	err = client.UpdateSubmission(ctx, models.Submission{SubmissionID: *subID4, ChatID: 1, ProblemID: *probID, Result: models.Correct})
	require.NoError(t, err)
	err = client.UpdateSubmission(ctx, models.Submission{SubmissionID: *subID5, ChatID: 2, ProblemID: *probID, Result: models.Correct})
	require.NoError(t, err)
	err = client.UpdateSubmission(ctx, models.Submission{SubmissionID: *subID6, ChatID: 2, ProblemID: *probID, Result: models.Correct})
	require.NoError(t, err)

	res, err := client.GetRating(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 2, res.Position)
	require.Equal(t, 3, res.All)

}
