package main

import (
	"inventory-service/internal/delivery/http"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
	"inventory-service/pkg/database"
	"inventory-service/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Инициализация логгера
	logger.Init()

	// Подключение к PostgreSQL
	db, err := database.NewPostgresConnection()
	if err != nil {
		logger.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	// Инициализация репозиториев
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	// Инициализация use cases
	productUC := usecase.NewProductUseCase(productRepo)
	categoryUC := usecase.NewCategoryUseCase(categoryRepo)

	// Инициализация HTTP хендлеров
	productHandler := http.NewProductHandler(productUC)
	categoryHandler := http.NewCategoryHandler(categoryUC)

	// Настройка маршрутов Gin
	router := gin.Default()

	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.POST("", productHandler.CreateProduct)
			products.GET("/:id", productHandler.GetProduct)
			products.PATCH("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
			products.GET("", productHandler.ListProducts)
		}

		categories := api.Group("/categories")
		{
			// аналогичные маршруты для категорий
		}
	}

	// Запуск сервера
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Fatal("Failed to start server", err)
	}
}
