package app

import (
	"context"
	"math/rand"

	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetRandom(ctx context.Context, req *pb.GetRandomRequest) (*pb.GetRandomResponse, error) {
	randTaskNumber := rand.Intn(10) + 1

	resp, err := s.GetProblem(ctx, &pb.GetProblemRequest{ChatId: req.ChatId, TaskNumber: int64(randTaskNumber)})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get random problem, err: <%v>", err)
	}

	return &pb.GetRandomResponse{TaskNumber: int64(randTaskNumber), Problem: resp.Problem}, nil
}
