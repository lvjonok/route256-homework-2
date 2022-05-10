package app

import (
	"context"

	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetRandom(ctx context.Context, req *pb.GetRandomRequest) (*pb.GetRandomResponse, error) {
	p, err := s.DB.GetRandomProblem(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get random, err: %v", err)
	}

	pbProblem := pb.Problem{
		ProblemId:   int64(p.ProblemID),
		Image:       p.ProblemImage,
		Description: p.Parts,
	}

	return &pb.GetRandomResponse{TaskNumber: int64(p.TaskNumber), Problem: &pbProblem}, nil
}
