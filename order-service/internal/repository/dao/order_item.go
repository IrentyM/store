package dao

import (
	"order-service/internal/domain"
)

type OrderItem struct {
	ID              int32   `db:"id"`
	OrderID         int32   `db:"order_id"`
	ProductID       int32   `db:"product_id"`
	Quantity        int32   `db:"quantity"`
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
