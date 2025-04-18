package app

import (
	"fmt"
	"log"
	"net"
	grpchandler "order-service/internal/adapter/grpc"
	"order-service/internal/repository"
	"order-service/internal/server"
	"order-service/internal/usecase"
	"order-service/pkg/db"
	orderproto "order-service/proto/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcServer *grpc.Server
}

func New(config *server.Config) (*App, error) {
	log.Println("connecting to postgresql", "database", config.DBname)
	database, err := db.PostgresConnection(config.DBhost, config.DBport, config.DBuser, config.DBpassword, config.DBname)
	if err != nil {
		return nil, fmt.Errorf("postgresql: %w", err)
	}

	grpcServer := grpc.NewServer()

	orderRepo := repository.NewOrderRepository(database)
	orderItemRepo := repository.NewOrderItemRepository(database)

	orderUseCase := usecase.NewOrderUseCase(orderRepo, orderItemRepo)

	orderHandler := grpchandler.NewOrderServer(*orderUseCase)

	orderproto.RegisterOrderServiceServer(grpcServer, orderHandler)

	reflection.Register(grpcServer)

	return &App{
		grpcServer: grpcServer,
	}, nil
}

func (s *App) Run(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	log.Printf("gRPC server is running on port %s", port)
	return s.grpcServer.Serve(listener)
}

func (s *App) Stop() {
	s.grpcServer.GracefulStop()
}
