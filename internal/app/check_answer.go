package app

import (
	"context"

	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s service) CheckAnswer(ctx context.Context, req *pb.CheckAnswerRequest) (*pb.CheckAnswerResponse, error) {

	return nil, status.Error(codes.Unimplemented, "")
}