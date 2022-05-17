package app_test

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/app"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
	homework_2 "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
)

func TestCheckAnswerWrong(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	srv := app.New(mockDB)

	mockDB.GetLastUserSubmissionMock.Return(&models.Submission{}, nil)
	mockDB.GetProblemMock.Return(&models.Problem{Answer: "hello"}, nil)
	mockDB.UpdateSubmissionMock.Return(nil)

	ctx := context.Background()
	resp, err := srv.CheckAnswer(ctx, &homework_2.CheckAnswerRequest{ChatId: 1, Answer: "some user answer"})

	require.NoError(t, err)
	require.Equal(t, "hello", resp.Answer)
	require.Equal(t, resp.Result, homework_2.SubmitResult_WRONG)
}

func TestCheckAnswerCorrect(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	srv := app.New(mockDB)

	mockDB.GetLastUserSubmissionMock.Return(&models.Submission{}, nil)
	mockDB.GetProblemMock.Return(&models.Problem{Answer: "correct"}, nil)
	mockDB.UpdateSubmissionMock.Return(nil)

	ctx := context.Background()
	resp, err := srv.CheckAnswer(ctx, &homework_2.CheckAnswerRequest{ChatId: 1, Answer: "correct"})

	require.NoError(t, err)
	require.Equal(t, "correct", resp.Answer)
	require.Equal(t, resp.Result, homework_2.SubmitResult_CORRECT)
}
