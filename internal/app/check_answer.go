package app

import (
	"context"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/db"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ChechAnswer compares last user's pending submission correct answer with given
func (s *Service) CheckAnswer(ctx context.Context, req *pb.CheckAnswerRequest) (*pb.CheckAnswerResponse, error) {
	lastUserSub, err := s.DB.GetLastUserSubmission(ctx, models.ID(req.ChatId))
	if err == db.ErrNotFound {
		return nil, status.Error(codes.NotFound, "cannot check answer without queried problem")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get last user submission, err: <%v>", err)
	}

	problem, err := s.DB.GetProblem(ctx, lastUserSub.ProblemID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get problem by id, err: <%v>", err)
	}

	var res models.Result
	if req.Answer == problem.Answer {
		res = models.Correct
	} else {
		res = models.Wrong
	}

	if err := s.DB.UpdateSubmission(ctx, models.Submission{
		SubmissionID: lastUserSub.SubmissionID,
		ChatID:       lastUserSub.ChatID,
		ProblemID:    lastUserSub.ProblemID,
		Result:       res,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update submission, err: <%v>", err)
	}

	var pbRes pb.SubmitResult
	if req.Answer == problem.Answer {
		pbRes = pb.SubmitResult_CORRECT
	} else {
		pbRes = pb.SubmitResult_WRONG
	}

	return &pb.CheckAnswerResponse{
		ProblemId: int64(problem.ProblemID),
		Answer:    problem.Answer,
		Result:    pbRes,
	}, nil
}
