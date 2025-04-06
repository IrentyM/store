package handler

import (
	"context"
	"inventory-service/internal/domain"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, product domain.Product) error
	GetProductByID(ctx context.Context, id int) (*domain.Product, error)
	UpdateProduct(ctx context.Context, id int, product domain.Product) error
	DeleteProduct(ctx context.Context, id int) error
	ListProducts(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Product, error)
}

type CategoryUseCase interface {
	CreateCategory(ctx context.Context, category domain.Category) error
	GetCategoryByID(ctx context.Context, id int) (*domain.Category, error)
	UpdateCategory(ctx context.Context, id int, category domain.Category) error
	DeleteCategory(ctx context.Context, id int) error
	ListCategories(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Category, error)
}
