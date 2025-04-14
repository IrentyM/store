package main

import (
	"log"

	"inventory-service/internal/adapter/grpc"
	"inventory-service/internal/repository"
	"inventory-service/internal/server"
	"inventory-service/internal/usecase"
	"inventory-service/pkg/db"
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
	productRepo := repository.NewProductRepository(database)
	productUseCase := usecase.NewProductUseCase(productRepo)

	// Start the gRPC server
	grpcServer := grpc.NewServer(*productUseCase)
	if err := grpcServer.Run(config.Port); err != nil {
		log.Fatalf("Failed to start the gRPC server: %v", err)
	}
}
