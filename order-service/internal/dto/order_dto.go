package dto

import (
	"order-service/internal/domain"
	"time"
)

type CreateOrderRequest struct {
	UserID        int                  `json:"user_id"`
	Status        domain.OrderStatus   `json:"status"`
	PaymentStatus domain.PaymentStatus `json:"payment_status"`
	TotalAmount   float64              `json:"total_amount"`
	Items         []OrderItemRequest   `json:"items"`
}

type OrderItemRequest struct {
	ProductID       int     `json:"product_id"`
	Quantity        int     `json:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase"`
}

func (r *CreateOrderRequest) ToDomain() (domain.Order, []domain.OrderItem) {
	order := domain.Order{
		UserID:        r.UserID,
		Status:        r.Status,
		PaymentStatus: r.PaymentStatus,
		TotalAmount:   r.TotalAmount,
	}

	var items []domain.OrderItem
	for _, item := range r.Items {
		items = append(items, domain.OrderItem{
			ProductID:       item.ProductID,
			Quantity:        item.Quantity,
			PriceAtPurchase: item.PriceAtPurchase,
		})
	}

	return order, items
}

type OrderResponse struct {
	ID            int                  `json:"id"`
	UserID        int                  `json:"user_id"`
	Status        domain.OrderStatus   `json:"status"`
	PaymentStatus domain.PaymentStatus `json:"payment_status"`
	TotalAmount   float64              `json:"total_amount"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
	Items         []OrderItemResponse  `json:"items"`
}

type OrderItemResponse struct {
	ID              int     `json:"id"`
	ProductID       int     `json:"product_id"`
	Quantity        int     `json:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase"`
}

func NewOrderResponse(order domain.Order, items []domain.OrderItem) OrderResponse {
	var itemResponses []OrderItemResponse
	for _, item := range items {
		itemResponses = append(itemResponses, OrderItemResponse{
			ID:              item.ID,
			ProductID:       item.ProductID,
			Quantity:        item.Quantity,
			PriceAtPurchase: item.PriceAtPurchase,
		})
	}

	return OrderResponse{
		ID:            order.ID,
		UserID:        order.UserID,
		Status:        order.Status,
		PaymentStatus: order.PaymentStatus,
		TotalAmount:   order.TotalAmount,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
		Items:         itemResponses,
	}
}
