package usecase

import (
	"context"
	"order-service/internal/domain"
)



type orderUseCase struct {
	orderRepo     OrderRepository
	orderItemRepo OrderItemRepository
}

func NewOrderUseCase(orderRepo OrderRepository, orderItemRepo OrderItemRepository) *orderUseCase {
	return &orderUseCase{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
	}
}

func (uc *orderUseCase) CreateOrder(ctx context.Context, order domain.Order, items []domain.OrderItem) (int, error) {
	// Validate the order
	if err := order.Validate(); err != nil {
		return 0, err
	}

	// Validate each order item
	for _, item := range items {
		if err := item.Validate(); err != nil {
			return 0, err
		}
	}

	// Create the order
	orderID, err := uc.orderRepo.Create(ctx, order)
	if err != nil {
		return 0, err
	}

	// Create the order items
	for _, item := range items {
		item.OrderID = orderID
		if err := uc.orderItemRepo.Create(ctx, item); err != nil {
			return 0, err
		}
	}

	return orderID, nil
}

func (uc *orderUseCase) GetOrderByID(ctx context.Context, id int) (*domain.Order, []domain.OrderItem, error) {
	// Get the order
	order, err := uc.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	if order == nil {
		return nil, nil, nil // Order not found
	}

	// Get the order items
	items, err := uc.orderItemRepo.GetByOrderID(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	return order, items, nil
}

func (uc *orderUseCase) UpdateOrder(ctx context.Context, id int, order domain.Order) error {
	// Validate the order
	if err := order.Validate(); err != nil {
		return err
	}

	// Update the order
	return uc.orderRepo.Update(ctx, id, order)
}

func (uc *orderUseCase) DeleteOrder(ctx context.Context, id int) error {
	// Delete the order items
	if err := uc.orderItemRepo.DeleteByOrderID(ctx, id); err != nil {
		return err
	}

	// Delete the order
	return uc.orderRepo.Delete(ctx, id)
}

func (uc *orderUseCase) ListOrders(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]domain.Order, error) {
	// List the orders
	return uc.orderRepo.List(ctx, filter, limit, offset)
}
