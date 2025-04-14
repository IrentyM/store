package main

import (
	"log"

	grpchandler "inventory-service/internal/adapter/grpc"
	"inventory-service/internal/repository"
	"inventory-service/internal/server"
	"inventory-service/internal/usecase"
	"inventory-service/pkg/db"
)

func main() {
	config := server.GetConfig()

	database, err := db.PostgresConnection(config.DBhost, config.DBport, config.DBuser, config.DBpassword, config.DBname)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	productRepo := repository.NewProductRepository(database)
	categoryRepo := repository.NewCategoryRepository(database)

	productUseCase := usecase.NewProductUseCase(productRepo)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)

	grpcServer := grpchandler.NewServer(*productUseCase, *categoryUseCase)
	if err := grpcServer.Run(config.Port); err != nil {
		log.Fatalf("Failed to start the gRPC server: %v", err)
	}
}
