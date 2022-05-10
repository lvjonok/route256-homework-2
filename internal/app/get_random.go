package app

import (
	"context"
	"math/rand"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetRandom(ctx context.Context, req *pb.GetRandomRequest) (*pb.GetRandomResponse, error) {
	randTaskNumber := rand.Intn(10) + 1

	p, err := s.DB.GetProblemByTaskNumber(ctx, randTaskNumber)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get random, err: %v", err)
	}

	pbProblem := pb.Problem{
		ProblemId:   int64(p.ProblemID),
		Image:       p.ProblemImage,
		Description: p.Parts,
	}

	if err := s.DB.UpdateAbortedSubmissions(ctx, models.ID(req.ChatId)); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to abort pending subs of user: %v, err: %v", req.ChatId, err)
	}

	if _, err := s.DB.CreateSubmission(ctx, models.Submission{
		ChatID:    models.ID(req.ChatId),
		ProblemID: p.ProblemID,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create submission, err: %v", err)
	}

	return &pb.GetRandomResponse{TaskNumber: int64(randTaskNumber), Problem: &pbProblem}, nil
}
