package usecase

import (
	"context"
	"fmt"
	"log"
	natsadapter "order-service/internal/adapter/nats"
	"order-service/internal/domain"
)

type OrderUseCase struct {
	orderRepo     OrderRepository
	orderItemRepo OrderItemRepository
	publisher     *natsadapter.OrderEventPublisher
}

func NewOrderUseCase(orderRepo OrderRepository, orderItemRepo OrderItemRepository, publisher *natsadapter.OrderEventPublisher) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
		publisher:     publisher,
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

	if err := uc.publisher.PublishOrderCreated(ctx, &order, items); err != nil {
		log.Printf("Failed to publish order created event: %v", err)
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
	err := uc.orderRepo.Update(ctx, id, order)
	if err != nil {
		return err
	}

	// if err := uc.publisher.PublishOrderCreated(ctx, &order, items); err != nil {
	// 	log.Printf("Failed to publish order created event: %v", err)
	// }
	return nil
}

func (uc *OrderUseCase) UpdateOrderStatus(ctx context.Context, id int32, status string, paystatus string) error {
	if status == "" {
		return fmt.Errorf("order status cannot be empty")
	}
	if paystatus == "" {
		return fmt.Errorf("payment status cannot be empty")
	}

	order, _, err := uc.GetOrderByID(ctx, int(id))
	if err != nil {
		return fmt.Errorf("failed to fetch order: %w", err)
	}
	if order == nil {
		return fmt.Errorf("order with ID %d not found", id)
	}

	order.Status = status
	order.PaymentStatus = paystatus

	if err := uc.orderRepo.Update(ctx, int(id), *order); err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}

	return nil
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
