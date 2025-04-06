package server

import (
	"log"

	handler "inventory-service/internal/delivery/http"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
	"inventory-service/pkg/db"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	// Initialize database connection
	database, err := db.NewPostgresConnection()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close()

	// Initialize repositories
	categoryRepo := repository.NewCategoryRepository(database)
	productRepo := repository.NewProductRepository(database)

	// Initialize use cases
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	productUseCase := usecase.NewProductUseCase(productRepo)

	// Initialize Gin router
	router := gin.Default()

	// Register handlers
	handler.NewCategoryHandler(router, *categoryUseCase)
	handler.NewProductHandler(router, *productUseCase)

	// Start the server
	log.Println("Starting server on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
