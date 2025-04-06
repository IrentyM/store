package repository

import (
	"context"
	"inventory-service/internal/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product domain.Product) error
	GetByID(ctx context.Context, id int) (*domain.Product, error)
	Update(ctx context.Context, id int, product domain.Product) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Product, error)
}
