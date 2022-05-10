package app

import (
	"context"

	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s service) GetRating(ctx context.Context, req *pb.GetRatingRequest) (*pb.GetRatingResponse, error) {

	return nil, status.Error(codes.Unimplemented, "")
}
