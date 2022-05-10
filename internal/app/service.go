package app

import (
	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
)

type service struct {
	DB DB // interface of available db functions
	pb.UnimplementedMathHelperServer
}

func New(dbClient DB) *service {
	return &service{DB: dbClient}
}
