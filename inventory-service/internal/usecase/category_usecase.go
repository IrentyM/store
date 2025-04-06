package usecase

import (
	"context"
	"inventory-service/internal/domain"
)

type CategoryUseCase interface {
	CreateCategory(ctx context.Context, category domain.Category) error
	GetCategoryByID(ctx context.Context, id int) (*domain.Category, error)
	UpdateCategory(ctx context.Context, id int, category domain.Category) error
	DeleteCategory(ctx context.Context, id int) error
	ListCategories(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Category, error)
}

type categoryUseCase struct {
	repo CategoryRepository
}

func NewCategoryUseCase(repo CategoryRepository) CategoryUseCase {
	return &categoryUseCase{repo: repo}
}

func (uc *categoryUseCase) CreateCategory(ctx context.Context, category domain.Category) error {
	if err := category.Validate(); err != nil {
		return err
	}
	return uc.repo.Create(ctx, category)
}

func (uc *categoryUseCase) GetCategoryByID(ctx context.Context, id int) (*domain.Category, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *categoryUseCase) UpdateCategory(ctx context.Context, id int, category domain.Category) error {
	if err := category.Validate(); err != nil {
		return err
	}
	return uc.repo.Update(ctx, id, category)
}

func (uc *categoryUseCase) DeleteCategory(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *categoryUseCase) ListCategories(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Category, error) {
	return uc.repo.List(ctx, filter, limit, offset)
}
