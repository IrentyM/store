package usecase

import (
	"context"
	"fmt"
	"inventory-service/internal/domain"
)

type ProductUseCase struct {
	repo  ProductRepository
	cache ProductCache
}

func NewProductUseCase(repo ProductRepository, productCache ProductCache) *ProductUseCase {
	return &ProductUseCase{
		repo:  repo,
		cache: productCache,
	}
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, product domain.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}

	// Create in database first
	err := uc.repo.Create(ctx, product)
	if err != nil {
		return err
	}

	// Update cache in background
	go func() {
		// Get the full product with generated ID
		fullProduct, err := uc.repo.GetByID(ctx, int(product.ID))
		if err != nil {
			fmt.Printf("Failed to get created product for caching: %v\n", err)
			return
		}

		// Set in cache
		if err := uc.cache.SetProduct(context.Background(), fullProduct); err != nil {
			fmt.Printf("Failed to cache created product: %v\n", err)
		}

		// Invalidate products list cache
		if err := uc.cache.InvalidateProductsList(context.Background()); err != nil {
			fmt.Printf("Failed to invalidate products list cache: %v\n", err)
		}
	}()

	return nil
}

func (uc *ProductUseCase) GetProductByID(ctx context.Context, id int) (*domain.Product, error) {
	// Try cache first
	cachedProduct, err := uc.cache.GetProduct(ctx, int32(id))
	if err == nil {
		return cachedProduct, nil
	}

	// Fallback to repository
	product, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update cache in background
	go func() {
		if err := uc.cache.SetProduct(context.Background(), product); err != nil {
			fmt.Printf("Failed to cache product: %v\n", err)
		}
	}()

	return product, nil
}

func (uc *ProductUseCase) UpdateProduct(ctx context.Context, id int, product domain.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}

	err := uc.repo.Update(ctx, id, product)
	if err != nil {
		return err
	}

	// Invalidate cache in background
	go func() {
		// Invalidate both the specific product and the list
		if err := uc.cache.InvalidateProduct(context.Background(), int32(id)); err != nil {
			fmt.Printf("Failed to invalidate product cache: %v\n", err)
		}
		if err := uc.cache.InvalidateProductsList(context.Background()); err != nil {
			fmt.Printf("Failed to invalidate products list cache: %v\n", err)
		}
	}()

	return nil
}

func (uc *ProductUseCase) DeleteProduct(ctx context.Context, id int) error {
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	// Invalidate cache in background
	go func() {
		if err := uc.cache.InvalidateProduct(context.Background(), int32(id)); err != nil {
			fmt.Printf("Failed to invalidate product cache: %v\n", err)
		}
		if err := uc.cache.InvalidateProductsList(context.Background()); err != nil {
			fmt.Printf("Failed to invalidate products list cache: %v\n", err)
		}
	}()

	return nil
}

func (uc *ProductUseCase) ListProducts(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Product, error) {
	return uc.repo.List(ctx, filter, limit, offset)
}
