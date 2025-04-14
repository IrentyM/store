package usecase

import (
	"context"
	"inventory-service/internal/domain"
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
