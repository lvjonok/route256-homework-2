package app

import (
	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
)

type Service struct {
	DB DB // interface of available db functions
	pb.UnimplementedMathHelperServer
}

func New(dbClient DB) *Service {
	return &Service{DB: dbClient}
}
