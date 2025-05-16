package usecase

import (
	"context"
	"fmt"

	"inventory-service/internal/domain"
)

type CategoryUseCase struct {
	repo  CategoryRepository
	cache CategoryCache
}

func NewCategoryUseCase(repo CategoryRepository, categoryCache CategoryCache) *CategoryUseCase {
	return &CategoryUseCase{
		repo:  repo,
		cache: categoryCache,
	}
}

func (uc *CategoryUseCase) CreateCategory(ctx context.Context, category domain.Category) error {
	if err := category.Validate(); err != nil {
		return err
	}

	err := uc.repo.Create(ctx, category)
	if err != nil {
		return err
	}

	// Update cache in background
	go func() {
		// Get the full category with generated ID
		fullCategory, err := uc.repo.GetByID(ctx, int(category.ID))
		if err != nil {
			fmt.Printf("Failed to get created category for caching: %v\n", err)
			return
		}

		// Set in cache
		if err := uc.cache.SetCategory(context.Background(), fullCategory); err != nil {
			fmt.Printf("Failed to cache created category: %v\n", err)
		}

		// Invalidate categories list cache
		if err := uc.cache.InvalidateCategoriesList(context.Background()); err != nil {
			fmt.Printf("Failed to invalidate categories list cache: %v\n", err)
		}
	}()

	return nil
}

func (uc *CategoryUseCase) GetCategoryByID(ctx context.Context, id int) (*domain.Category, error) {
	// Try cache first
	cachedCategory, err := uc.cache.GetCategory(ctx, int32(id))
	if err == nil {
		return cachedCategory, nil
	}

	// Fallback to repository
	category, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update cache in background
	go func() {
		if err := uc.cache.SetCategory(context.Background(), category); err != nil {
			fmt.Printf("Failed to cache category: %v\n", err)
		}
	}()

	return category, nil
}

func (uc *CategoryUseCase) UpdateCategory(ctx context.Context, id int32, category domain.Category) error {
	if err := category.Validate(); err != nil {
		return err
	}

	err := uc.repo.Update(ctx, id, category)
	if err != nil {
		return err
	}

	// Invalidate cache in background
	go func() {
		// Invalidate both the specific category and the list
		if err := uc.cache.InvalidateCategory(context.Background(), id); err != nil {
			fmt.Printf("Failed to invalidate category cache: %v\n", err)
		}
		if err := uc.cache.InvalidateCategoriesList(context.Background()); err != nil {
			fmt.Printf("Failed to invalidate categories list cache: %v\n", err)
		}
	}()

	return nil
}

func (uc *CategoryUseCase) DeleteCategory(ctx context.Context, id int) error {
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	// Invalidate cache in background
	go func() {
		if err := uc.cache.InvalidateCategory(context.Background(), int32(id)); err != nil {
			fmt.Printf("Failed to invalidate category cache: %v\n", err)
		}
		if err := uc.cache.InvalidateCategoriesList(context.Background()); err != nil {
			fmt.Printf("Failed to invalidate categories list cache: %v\n", err)
		}
	}()

	return nil
}

func (uc *CategoryUseCase) ListCategories(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Category, error) {
	return uc.repo.List(ctx, filter, limit, offset)
}
