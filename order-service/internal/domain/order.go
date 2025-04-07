package domain

import (
	"fmt"
	"time"
)

type OrderStatus string
type PaymentStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusRefunded   OrderStatus = "refunded"

	PaymentStatusPending  PaymentStatus = "pending"
	PaymentStatusPaid     PaymentStatus = "paid"
	PaymentStatusFailed   PaymentStatus = "failed"
	PaymentStatusRefunded PaymentStatus = "refunded"
)

type Order struct {
	ID            int           `json:"id"`
	UserID        int           `json:"user_id"`
	Status        OrderStatus   `json:"status"`
	PaymentStatus PaymentStatus `json:"payment_status"`
	TotalAmount   float64       `json:"total_amount"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

func (o *Order) Validate() error {
	switch {
	case o.UserID <= 0:
		return fmt.Errorf("invalid user ID")
	case o.TotalAmount < 0:
		return fmt.Errorf("total amount cannot be negative")
	case o.Status == "":
		return fmt.Errorf("order status cannot be empty")
	case o.PaymentStatus == "":
		return fmt.Errorf("payment status cannot be empty")
	default:
		return nil
	}
}
