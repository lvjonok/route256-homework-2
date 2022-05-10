package main

import (
	"context"
	"log"
	"net"
	"time"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/app"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/db"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/dbconnector"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/mw"
	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	adp, err := dbconnector.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	newServer := app.New(db.New(adp))
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// var opts []grpc.ServerOption
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(mw.LogInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMathHelperServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
