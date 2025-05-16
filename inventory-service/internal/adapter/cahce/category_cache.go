package cache

import (
	"context"
	"encoding/json"
	"strconv"

	"inventory-service/internal/domain"
	"inventory-service/internal/dto"
)

type CategoryCache struct {
	cache Cache
}

func NewCategoryCache(cache Cache) *CategoryCache {
	return &CategoryCache{cache: cache}
}

func (c *CategoryCache) GetCategory(ctx context.Context, id int32) (*domain.Category, error) {
	ID := strconv.Itoa(int(id))
	key := "category:" + ID
	val, err := c.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var category domain.Category
	if err := json.Unmarshal([]byte(val), &category); err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *CategoryCache) SetCategory(ctx context.Context, category *domain.Category) error {
	ID := strconv.Itoa(int(category.ID))
	key := "category:" + ID
	val, err := json.Marshal(category)
	if err != nil {
		return err
	}

	return c.cache.Set(ctx, key, val, defaultExpiration)
}

func (c *CategoryCache) GetCategories(ctx context.Context) ([]dto.CategoryResponse, error) {
	val, err := c.cache.Get(ctx, "categories:list")
	if err != nil {
		return nil, err
	}

	var categories []dto.CategoryResponse
	if err := json.Unmarshal([]byte(val), &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *CategoryCache) SetCategories(ctx context.Context, categories []dto.CategoryResponse) error {
	val, err := json.Marshal(categories)
	if err != nil {
		return err
	}

	return c.cache.Set(ctx, "categories:list", val, defaultExpiration)
}

func (c *CategoryCache) InvalidateCategory(ctx context.Context, id int32) error {
	ID := strconv.Itoa(int(id))
	return c.cache.Delete(ctx, "category:"+ID)
}

func (c *CategoryCache) InvalidateCategoriesList(ctx context.Context) error {
	return c.cache.Delete(ctx, "categories:list")
}
