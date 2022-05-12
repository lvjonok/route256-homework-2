package app

import (
	"context"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetProblem(ctx context.Context, req *pb.GetProblemRequest) (*pb.GetProblemResponse, error) {
	if req.TaskNumber < 1 || req.TaskNumber > 11 {
		return nil, status.Error(codes.OutOfRange, "wrong task number, only [1, 11] are allowed")
	}

	problem, err := s.DB.GetProblemByTaskNumber(ctx, int(req.TaskNumber))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get problem by task number, err: %v", err)
	}

	if err := s.DB.UpdateAbortedSubmissions(ctx, models.ID(req.ChatId)); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to abort pending subs of user: %v, err: %v", req.ChatId, err)
	}

	if _, err := s.DB.CreateSubmission(ctx, models.Submission{
		ChatID:    models.ID(req.ChatId),
		ProblemID: problem.ProblemID,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create submission, err: %v", err)
	}

	return &pb.GetProblemResponse{Problem: &pb.Problem{
		ProblemId:   int64(problem.ProblemID),
		Image:       problem.ProblemImage,
		Description: problem.Parts,
	}}, nil
}
