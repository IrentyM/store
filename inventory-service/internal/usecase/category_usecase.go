package usecase

import (
	"context"
	"inventory-service/internal/domain"
)

type CategoryUseCase struct {
	repo CategoryRepository
}

func NewCategoryUseCase(repo CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{repo: repo}
}

func (uc *CategoryUseCase) CreateCategory(ctx context.Context, category domain.Category) error {
	if err := category.Validate(); err != nil {
		return err
	}
	return uc.repo.Create(ctx, category)
}

func (uc *CategoryUseCase) GetCategoryByID(ctx context.Context, id int) (*domain.Category, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *CategoryUseCase) UpdateCategory(ctx context.Context, id int32, category domain.Category) error {
	if err := category.Validate(); err != nil {
		return err
	}
	return uc.repo.Update(ctx, id, category)
}

func (uc *CategoryUseCase) DeleteCategory(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *CategoryUseCase) ListCategories(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Category, error) {
	return uc.repo.List(ctx, filter, limit, offset)
}
