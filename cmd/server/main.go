package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/app"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/config"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/db"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/dbconnector"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/mw"
	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runRest(cfg *config.Config) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterMathHelperHandlerFromEndpoint(ctx, mux, cfg.Server.Host+":"+cfg.Server.Port, opts)
	if err != nil {
		log.Fatalf("failed to register endpoint, err: <%v>", err)
	}
	if err := http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.RestPort, mux); err != nil {
		log.Fatalf("failed to listen and serve rest, err: <%v>", err)
	}
}

func main() {
	cfg, err := config.New("config.yaml")
	if err != nil {
		panic(err)
	}
	go runRest(cfg)

	ctx := context.Background()

	adp, err := dbconnector.New(ctx, cfg.Database.Url)
	if err != nil {
		log.Fatalf("failed to connect to database, err: <%v>", err)
	}

	newServer := app.New(db.New(adp))

	// Parse problems worker
	go func(delaySec int64) {
		for {
			log.Printf("Worker start")
			err := newServer.ParseProblems(ctx)
			if err != nil {
				log.Printf("worker encountered err: <%v>", err)
			}
			log.Printf("worker finished")
			time.Sleep(time.Second * time.Duration(delaySec))
		}
	}(cfg.Parser.DelaySec)

	lis, err := net.Listen("tcp", cfg.Server.Host+":"+cfg.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(mw.LogInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMathHelperServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc, err: <%v>", err)
	}

	for {
		time.Sleep(time.Second)
	}
}
