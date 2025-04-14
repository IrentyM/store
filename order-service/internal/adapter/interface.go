package handler

import (
	"context"
	"order-service/internal/domain"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, order domain.Order, items []domain.OrderItem) (int32, error)
	GetOrderByID(ctx context.Context, id int) (*domain.Order, []domain.OrderItem, error)
	UpdateOrder(ctx context.Context, id int, order domain.Order) error
	DeleteOrder(ctx context.Context, id int) error
	ListOrders(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]domain.Order, error)
}
