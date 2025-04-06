package usecase

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

type productUseCase struct {
	repo ProductRepository
}

func NewProductUseCase(repo ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}

func (uc *productUseCase) CreateProduct(ctx context.Context, product domain.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}
	return uc.repo.Create(ctx, product)
}

func (uc *productUseCase) GetProductByID(ctx context.Context, id int) (*domain.Product, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *productUseCase) UpdateProduct(ctx context.Context, id int, product domain.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}
	return uc.repo.Update(ctx, id, product)
}

func (uc *productUseCase) DeleteProduct(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *productUseCase) ListProducts(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Product, error) {
	return uc.repo.List(ctx, filter, limit, offset)
}
