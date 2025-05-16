package usecase

import (
	"context"
	"inventory-service/internal/domain"
	"inventory-service/internal/dto"
)

type CategoryRepository interface {
	Create(ctx context.Context, category domain.Category) error
	GetByID(ctx context.Context, id int) (*domain.Category, error)
	Update(ctx context.Context, id int32, category domain.Category) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Category, error)
}

type ProductRepository interface {
	Create(ctx context.Context, product domain.Product) error
	GetByID(ctx context.Context, id int) (*domain.Product, error)
	Update(ctx context.Context, id int, product domain.Product) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Product, error)
}

type CategoryCache interface {
	GetCategory(ctx context.Context, id int32) (*domain.Category, error)
	SetCategory(ctx context.Context, category *domain.Category) error
	GetCategories(ctx context.Context) ([]dto.CategoryResponse, error)
	SetCategories(ctx context.Context, categories []dto.CategoryResponse) error
	InvalidateCategory(ctx context.Context, id int32) error
	InvalidateCategoriesList(ctx context.Context) error
}
type ProductCache interface {
	GetProduct(ctx context.Context, id int32) (*domain.Product, error)
	SetProduct(ctx context.Context, product *domain.Product) error
	GetProducts(ctx context.Context) ([]dto.ProductResponse, error)
	SetProducts(ctx context.Context, products []dto.ProductResponse) error
	InvalidateProduct(ctx context.Context, id int32) error
	InvalidateProductsList(ctx context.Context) error
}
