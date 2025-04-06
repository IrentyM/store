package usecase

import (
	"context"
	"inventory-service/internal/domain"
)

type ProductUseCase struct {
	repo ProductRepository
}

func NewProductUseCase(repo ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, product domain.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}
	return uc.repo.Create(ctx, product)
}

func (uc *ProductUseCase) GetProductByID(ctx context.Context, id int) (*domain.Product, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *ProductUseCase) UpdateProduct(ctx context.Context, id int, product domain.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}
	return uc.repo.Update(ctx, id, product)
}

func (uc *ProductUseCase) DeleteProduct(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *ProductUseCase) ListProducts(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Product, error) {
	return uc.repo.List(ctx, filter, limit, offset)
}
