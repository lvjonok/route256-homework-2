package app

import (
	"context"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetRating(ctx context.Context, req *pb.GetRatingRequest) (*pb.GetRatingResponse, error) {
	rat, err := s.DB.GetRating(ctx, models.ID(req.ChatId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to query rating for user: %v, err: <%v>", req.ChatId, err)
	}

	return &pb.GetRatingResponse{
		Position: int64(rat.Position),
		All:      int64(rat.All),
	}, nil
}
