package usecase

import (
	"context"
	"order-service/internal/domain"
)

type OrderUseCase struct {
	orderRepo     OrderRepository
	orderItemRepo OrderItemRepository
}

func NewOrderUseCase(orderRepo OrderRepository, orderItemRepo OrderItemRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
	}
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, order domain.Order, items []domain.OrderItem) (int32, error) {
	if err := order.Validate(); err != nil {
		return 0, err
	}

	for _, item := range items {
		if err := item.Validate(); err != nil {
			return 0, err
		}
	}

	orderID, err := uc.orderRepo.Create(ctx, order)
	if err != nil {
		return 0, err
	}

	for _, item := range items {
		item.OrderID = orderID
		if err := uc.orderItemRepo.Create(ctx, item); err != nil {
			return 0, err
		}
	}

	return orderID, nil
}

func (uc *OrderUseCase) GetOrderByID(ctx context.Context, id int) (*domain.Order, []domain.OrderItem, error) {
	order, err := uc.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	if order == nil {
		return nil, nil, nil
	}

	items, err := uc.orderItemRepo.GetByOrderID(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	return order, items, nil
}

func (uc *OrderUseCase) UpdateOrder(ctx context.Context, id int, order domain.Order) error {
	if err := order.Validate(); err != nil {
		return err
	}

	return uc.orderRepo.Update(ctx, id, order)
}

func (uc *OrderUseCase) DeleteOrder(ctx context.Context, id int) error {
	if err := uc.orderItemRepo.DeleteByOrderID(ctx, id); err != nil {
		return err
	}

	return uc.orderRepo.Delete(ctx, id)
}

func (uc *OrderUseCase) ListOrders(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]domain.Order, error) {
	return uc.orderRepo.List(ctx, filter, limit, offset)
}
