package app

import (
	"context"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) GetStat(ctx context.Context, req *pb.GetStatRequest) (*pb.GetStatResponse, error) {
	stat, err := s.DB.GetStat(ctx, models.ID(req.ChatId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to query statistics for user: %v, err: <%v>", req.ChatId, err)
	}

	resp := pb.GetStatResponse{}
	for _, data := range stat.Stat {
		resp.Stat = append(resp.Stat, &pb.TaskStat{
			Correct:    int64(data.Correct),
			All:        int64(data.All),
			TaskNumber: int64(data.TaskNumber),
		})
	}

	return &resp, nil
}
