package dao

import (
	"order-service/internal/domain"
)

type OrderItem struct {
	ID              int     `db:"id"`
	OrderID         int     `db:"order_id"`
	ProductID       int     `db:"product_id"`
	Quantity        int     `db:"quantity"`
	PriceAtPurchase float64 `db:"price_at_purchase"`
}

func ToOrderItem(orderItem domain.OrderItem) OrderItem {
	return OrderItem{
		ID:              orderItem.ID,
		OrderID:         orderItem.OrderID,
		ProductID:       orderItem.ProductID,
		Quantity:        orderItem.Quantity,
		PriceAtPurchase: orderItem.PriceAtPurchase,
	}
}

func FromOrderItem(orderItem OrderItem) domain.OrderItem {
	return domain.OrderItem{
		ID:              orderItem.ID,
		OrderID:         orderItem.OrderID,
		ProductID:       orderItem.ProductID,
		Quantity:        orderItem.Quantity,
		PriceAtPurchase: orderItem.PriceAtPurchase,
	}
}
