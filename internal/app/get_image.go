package app

import (
	"context"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetImage(ctx context.Context, req *pb.GetImageRequest) (*pb.GetImageResponse, error) {
	image, err := s.DB.GetImage(ctx, models.ID(req.ImageId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get image by id, err: %v", err)
	}

	return &pb.GetImageResponse{Image: image.Content}, nil
}
