package main

import (
	"inventory-service/internal/repository"
	"inventory-service/pkg/utils"
)

func main() {
	db := utils.InitPostgres()

	// Автоматическая миграция
	db.AutoMigrate(&repository.Product{})

	productRepo := repository.NewProductRepository(db)

	// Тест: создать продукт
	productRepo.Create(&repository.Product{
		Name:     "Laptop",
		Category: "Electronics",
		Price:    1299.99,
		Stock:    10,
	})
}
