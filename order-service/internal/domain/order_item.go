package domain

import "fmt"

type OrderItem struct {
	ID              int32   `json:"id"`
	OrderID         int32   `json:"order_id"`
	ProductID       int32   `json:"product_id"`
	Quantity        int32   `json:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase"`
}

func (oi *OrderItem) Validate() error {
	switch {
	case oi.OrderID <= 0:
		return fmt.Errorf("invalid order ID")
	case oi.ProductID <= 0:
		return fmt.Errorf("invalid product ID")
	case oi.Quantity <= 0:
		return fmt.Errorf("quantity must be greater than zero")
	case oi.PriceAtPurchase < 0:
		return fmt.Errorf("price at purchase cannot be negative")
	default:
		return nil
	}
}
