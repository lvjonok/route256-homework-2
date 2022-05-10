package app

import (
	"context"

	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s service) GetProblem(ctx context.Context, req *pb.GetProblemRequest) (*pb.GetProblemResponse, error) {
	if req.TaskNumber < 1 || req.TaskNumber > 11 {
		return nil, status.Error(codes.OutOfRange, "wrong task number, only [1, 11] are allowed")
	}

	problem, err := s.DB.GetProblemByTaskNumber(ctx, int(req.TaskNumber))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get problem by task number, err: %v", err)
	}

	return &pb.GetProblemResponse{Problem: &pb.Problem{
		ProblemId:   int64(problem.ProblemID),
		Image:       problem.ProblemImage,
		Description: problem.Parts,
	}}, nil
}
