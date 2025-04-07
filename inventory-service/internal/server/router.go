package server

import (
	handler "inventory-service/internal/delivery/http"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
	"inventory-service/pkg/db"
)

func (s *server) registerRoutes() error {
	// Initialize database connection
	database, err := db.PostgresConnection(s.cfg.DBhost, s.cfg.DBport, s.cfg.DBuser, s.cfg.DBpassword, s.cfg.DBname)
	if err != nil {
		return err
	}

	// Initialize repositories
	categoryRepo := repository.NewCategoryRepository(database)
	productRepo := repository.NewProductRepository(database)

	// Initialize use cases
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	productUseCase := usecase.NewProductUseCase(productRepo)

	// Register health check route
	s.router.GET("/health", handler.GetHealth)

	// Register category routes
	categoryHandler := handler.NewCategoryHandler(*categoryUseCase)
	s.router.POST("/categories", categoryHandler.CreateCategory)
	s.router.GET("/categories/:id", categoryHandler.GetCategoryByID)
	s.router.PUT("/categories/:id", categoryHandler.UpdateCategory)
	s.router.DELETE("/categories/:id", categoryHandler.DeleteCategory)
	s.router.GET("/categories", categoryHandler.ListCategories)

	// Register product routes
	productHandler := handler.NewProductHandler(*productUseCase)
	s.router.POST("/products", productHandler.CreateProduct)
	s.router.GET("/products/:id", productHandler.GetProductByID)
	s.router.PUT("/products/:id", productHandler.UpdateProduct)
	s.router.DELETE("/products/:id", productHandler.DeleteProduct)
	s.router.GET("/products", productHandler.ListProducts)

	return nil
}
