// order-service/internal/adapter/nats/publisher.go
package natsadapter

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"order-service/internal/domain"

	"github.com/nats-io/nats.go"
)

type OrderEventPublisher struct {
	nc *nats.Conn
}

func NewOrderEventPublisher(nc *nats.Conn) *OrderEventPublisher {
	return &OrderEventPublisher{nc: nc}
}

func (p *OrderEventPublisher) PublishOrderCreated(ctx context.Context, order *domain.Order, items []domain.OrderItem) error {
	event := OrderEventDTO{
		EventID:   generateEventID(),
		Operation: "created",
		OrderID:   order.ID,
		UserID:    order.UserID,
		Items:     convertItemsToDTO(items),
		Total:     order.TotalAmount,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	if err := p.nc.Publish("order.created", data); err != nil {
		return err
	}

	log.Printf("Published order.created event for order %d", order.ID)
	return nil
}

func (p *OrderEventPublisher) PublishOrderUpdated(ctx context.Context, order *domain.Order, items []domain.OrderItem) error {
	event := OrderEventDTO{
		EventID:   generateEventID(),
		Operation: "updated",
		OrderID:   order.ID,
		UserID:    order.UserID,
		Items:     convertItemsToDTO(items),
		Total:     order.TotalAmount,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	if err := p.nc.Publish("order.updated", data); err != nil {
		return err
	}

	log.Printf("Published order.updated event for order %d", order.ID)
	return nil
}

func (p *OrderEventPublisher) PublishOrderDeleted(ctx context.Context, orderID int64, userID int64) error {
	event := OrderEventDTO{
		EventID:   generateEventID(),
		Operation: "deleted",
		OrderID:   int32(orderID),
		UserID:    int32(userID),
		Items:     nil,
		Total:     0,
		Status:    "deleted",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	if err := p.nc.Publish("order.deleted", data); err != nil {
		return err
	}

	log.Printf("Published order.deleted event for order %d", orderID)
	return nil
}

// Вспомогательные функции
func generateEventID() int64 {
	// Реализация генерации ID
	return time.Now().UnixNano()
}

func convertItemsToDTO(items []domain.OrderItem) []OrderItemDTO {
	var dtos []OrderItemDTO
	for _, item := range items {
		dtos = append(dtos, OrderItemDTO{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.PriceAtPurchase,
		})
	}
	return dtos
}

// DTO для событий
type OrderEventDTO struct {
	EventID   int64          `json:"event_id"`
	Operation string         `json:"operation"`
	OrderID   int32          `json:"order_id"`
	UserID    int32          `json:"user_id"`
	Items     []OrderItemDTO `json:"items"`
	Total     float64        `json:"total"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type OrderItemDTO struct {
	ProductID int32   `json:"product_id"`
	Quantity  int32   `json:"quantity"`
	Price     float64 `json:"price"`
}
