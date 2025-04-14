package usecase

import (
	"context"
	"order-service/internal/domain"
)

type OrderItemRepository interface {
	Create(ctx context.Context, orderItem domain.OrderItem) error
	GetByOrderID(ctx context.Context, orderID int) ([]domain.OrderItem, error)
	DeleteByOrderID(ctx context.Context, orderID int) error
}

type OrderRepository interface {
	Create(ctx context.Context, order domain.Order) (int32, error)
	GetByID(ctx context.Context, id int) (*domain.Order, error)
	Update(ctx context.Context, id int, order domain.Order) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]domain.Order, error)
}
