package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func TestGetStat(t *testing.T) {
	client, ctx := Prepare(t)

	// prepare category
	catID, err := client.CreateCategory(ctx, models.Category{CategoryID: 1, TaskNumber: 10, Title: "some title"})
	require.NoError(t, err)
	catID2, err := client.CreateCategory(ctx, models.Category{CategoryID: 1, TaskNumber: 1, Title: "some title"})
	require.NoError(t, err)

	// prepare problem
	probID, err := client.CreateProblem(ctx, models.Problem{ProblemID: 12345, CategoryID: *catID, Parts: []string{"some cool description"}, Answer: "10"})
	require.NoError(t, err)
	probID2, err := client.CreateProblem(ctx, models.Problem{ProblemID: 1234, CategoryID: *catID2, Parts: []string{"some cool description"}, Answer: "20"})
	require.NoError(t, err)

	// prepare submission
	subID, err := client.CreateSubmission(ctx, models.Submission{ChatID: 1, ProblemID: *probID, Result: models.Correct})
	require.NoError(t, err)
	subID2, err := client.CreateSubmission(ctx, models.Submission{ChatID: 1, ProblemID: *probID, Result: models.Wrong})
	require.NoError(t, err)
	subID3, err := client.CreateSubmission(ctx, models.Submission{ChatID: 1, ProblemID: *probID, Result: models.Aborted})
	require.NoError(t, err)
	subID4, err := client.CreateSubmission(ctx, models.Submission{ChatID: 1, ProblemID: *probID2, Result: models.Correct})
	require.NoError(t, err)
	subID5, err := client.CreateSubmission(ctx, models.Submission{ChatID: 1, ProblemID: *probID2, Result: models.Correct})
	require.NoError(t, err)
	subID6, err := client.CreateSubmission(ctx, models.Submission{ChatID: 1, ProblemID: *probID2, Result: models.Pending})
	require.NoError(t, err)

	err = client.UpdateSubmission(ctx, models.Submission{SubmissionID: *subID, ChatID: 1, ProblemID: *probID, Result: models.Correct})
	require.NoError(t, err)
	err = client.UpdateSubmission(ctx, models.Submission{SubmissionID: *subID2, ChatID: 1, ProblemID: *probID, Result: models.Wrong})
	require.NoError(t, err)
	err = client.UpdateSubmission(ctx, models.Submission{SubmissionID: *subID3, ChatID: 1, ProblemID: *probID, Result: models.Aborted})
	require.NoError(t, err)
	err = client.UpdateSubmission(ctx, models.Submission{SubmissionID: *subID4, ChatID: 1, ProblemID: *probID2, Result: models.Correct})
	require.NoError(t, err)
	err = client.UpdateSubmission(ctx, models.Submission{SubmissionID: *subID5, ChatID: 1, ProblemID: *probID2, Result: models.Correct})
	require.NoError(t, err)
	err = client.UpdateSubmission(ctx, models.Submission{SubmissionID: *subID6, ChatID: 1, ProblemID: *probID2, Result: models.Pending})
	require.NoError(t, err)

	res, err := client.GetStat(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 2, len(res.Stat))
	require.Equal(t, 1, res.Stat[0].TaskNumber, "wrong task number")
	require.Equal(t, 2, res.Stat[0].Correct, "wrong number of correct")
	require.Equal(t, 2, res.Stat[0].All, "wrong total")
	require.Equal(t, 10, res.Stat[1].TaskNumber, "wrong task number2")
	require.Equal(t, 1, res.Stat[1].Correct, "wrong number of correct2")
	require.Equal(t, 2, res.Stat[1].All, "wrong total2")

}
