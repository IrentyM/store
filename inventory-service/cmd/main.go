package main

import (
	"log"

	cache "inventory-service/internal/adapter/cahce"
	grpchandler "inventory-service/internal/adapter/grpc"
	"inventory-service/internal/repository"
	"inventory-service/internal/server"
	"inventory-service/internal/usecase"
	"inventory-service/pkg/db"
)

func main() {
	config := server.GetConfig()

	// Initialize database connection
	database, err := db.PostgresConnection(config.DBhost, config.DBport, config.DBuser, config.DBpassword, config.DBname)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize Redis cache
	redisCache, err := cache.NewRedisCache(config.RedisAddr, config.RedisPass, config.RedisDB)
	if err != nil {
		log.Fatalf("Failed to initialize Redis cache: %v", err)
	}

	// Initialize repositories
	productRepo := repository.NewProductRepository(database)
	categoryRepo := repository.NewCategoryRepository(database)

	// Initialize caches
	productCache := cache.NewProductCache(redisCache)
	categoryCache := cache.NewCategoryCache(redisCache)

	// Initialize use cases with both repositories and caches
	productUseCase := usecase.NewProductUseCase(productRepo, productCache)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo, categoryCache)

	// Initialize and run gRPC server
	grpcServer := grpchandler.NewServer(*productUseCase, *categoryUseCase)
	if err := grpcServer.Run(config.Port); err != nil {
		log.Fatalf("Failed to start the gRPC server: %v", err)
	}
}
