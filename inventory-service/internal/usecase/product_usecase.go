package usecase

import (
	"context"
	"inventory-service/internal/domain"
	"inventory-service/internal/repository"
)

type ProductUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, product *domain.Product) error {
	// Бизнес-логика перед созданием
	if product.Stock < 0 {
		return domain.ErrInvalidStock
	}
	return uc.repo.Create(ctx, product)
}
