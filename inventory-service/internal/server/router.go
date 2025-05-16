package server

import (
	"context"
	cache "inventory-service/internal/adapter/cahce"
	handler "inventory-service/internal/adapter/http"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
	"inventory-service/pkg/db"
	"log"
	"time"
)

func (s *server) registerRoutes() error {
	database, err := db.PostgresConnection(s.cfg.DBhost, s.cfg.DBport, s.cfg.DBuser, s.cfg.DBpassword, s.cfg.DBname)
	if err != nil {
		return err
	}

	// Initialize repositories
	categoryRepo := repository.NewCategoryRepository(database)
	productRepo := repository.NewProductRepository(database)

	// Initialize caches
	categoryCache := cache.NewCategoryCache(s.cache)
	productCache := cache.NewProductCache(s.cache)

	// Initialize use cases with both repositories and caches
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo, categoryCache)
	productUseCase := usecase.NewProductUseCase(productRepo, productCache)

	s.router.GET("/health", handler.GetHealth)

	// Initialize handlers
	categoryHandler := handler.NewCategoryHandler(*categoryUseCase)
	s.router.POST("/categories", categoryHandler.CreateCategory)
	s.router.GET("/categories/:id", categoryHandler.GetCategoryByID)
	s.router.PUT("/categories/:id", categoryHandler.UpdateCategory)
	s.router.DELETE("/categories/:id", categoryHandler.DeleteCategory)
	s.router.GET("/categories", categoryHandler.ListCategories)

	productHandler := handler.NewProductHandler(*productUseCase)
	s.router.POST("/products", productHandler.CreateProduct)
	s.router.GET("/products/:id", productHandler.GetProductByID)
	s.router.PUT("/products/:id", productHandler.UpdateProduct)
	s.router.DELETE("/products/:id", productHandler.DeleteProduct)
	s.router.GET("/products", productHandler.ListProducts)

	// Start background cache refresh
	go s.startCacheRefresh(categoryUseCase, productUseCase)

	return nil
}

func (s *server) startCacheRefresh(categoryUC *usecase.CategoryUseCase, productUC *usecase.ProductUseCase) {
	ticker := time.NewTicker(12 * time.Hour)
	defer ticker.Stop()

	ctx := context.Background()

	// Initial refresh
	s.refreshCaches(ctx, categoryUC, productUC)

	for {
		select {
		case <-ticker.C:
			s.refreshCaches(ctx, categoryUC, productUC)
		}
	}
}

func (s *server) refreshCaches(ctx context.Context, categoryUC *usecase.CategoryUseCase, productUC *usecase.ProductUseCase) {
	// Refresh product cache
	if _, err := productUC.ListProducts(ctx, nil, 0, 0); err != nil {
		log.Printf("Failed to refresh product cache: %v", err)
	}

	// Refresh category cache
	if _, err := categoryUC.ListCategories(ctx, nil, 0, 0); err != nil {
		log.Printf("Failed to refresh category cache: %v", err)
	}
}
