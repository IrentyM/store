package cache

import (
	"context"
	"encoding/json"
	"strconv"

	"inventory-service/internal/domain"
	"inventory-service/internal/dto"
)

type ProductCache struct {
	cache Cache
}

func NewProductCache(cache Cache) *ProductCache {
	return &ProductCache{cache: cache}
}

func (p *ProductCache) GetProduct(ctx context.Context, id int32) (*domain.Product, error) {
	ID := strconv.Itoa(int(id))
	key := "product:" + ID
	val, err := p.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var product domain.Product
	if err := json.Unmarshal([]byte(val), &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductCache) SetProduct(ctx context.Context, product *domain.Product) error {
	ID := strconv.Itoa(int(product.ID))
	key := "product:" + ID
	val, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return p.cache.Set(ctx, key, val, defaultExpiration)
}

func (p *ProductCache) GetProducts(ctx context.Context) ([]dto.ProductResponse, error) {
	val, err := p.cache.Get(ctx, "products:list")
	if err != nil {
		return nil, err
	}

	var products []dto.ProductResponse
	if err := json.Unmarshal([]byte(val), &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductCache) SetProducts(ctx context.Context, products []dto.ProductResponse) error {
	val, err := json.Marshal(products)
	if err != nil {
		return err
	}

	return p.cache.Set(ctx, "products:list", val, defaultExpiration)
}

func (p *ProductCache) InvalidateProduct(ctx context.Context, id int32) error {
	ID := strconv.Itoa(int(id))
	return p.cache.Delete(ctx, "product:"+ID)
}

func (p *ProductCache) InvalidateProductsList(ctx context.Context) error {
	return p.cache.Delete(ctx, "products:list")
}
