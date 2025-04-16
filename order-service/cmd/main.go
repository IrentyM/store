package main

import (
	"log"

	grpchandler "order-service/internal/adapter/grpc"
	"order-service/internal/repository"
	"order-service/internal/server"
	"order-service/internal/usecase"
	"order-service/pkg/db"
)

func main() {
	// Load configuration
	config := server.GetConfig()

	// Connect to the database
	database, err := db.PostgresConnection(config.DBhost, config.DBport, config.DBuser, config.DBpassword, config.DBname)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize repositories and use cases
	orderRepo := repository.NewOrderRepository(database)
	orderItemRepo := repository.NewOrderItemRepository(database)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, orderItemRepo)

	// Start the gRPC server
	grpcServer := grpchandler.NewServer(*orderUseCase)
	if err := grpcServer.Run(config.Port); err != nil {
		log.Fatalf("Failed to start the gRPC server: %v", err)
	}
}
