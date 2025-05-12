package model

import (
	"time"
)

type OrderEvent struct {
	ID        int64
	EventID   int64
	Operation string
	OrderID   int64
	UserID    int64
	Items     []OrderItem
	Total     float64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItem struct {
	ProductID int64
	Quantity  int
	Price     float64
}

type InventoryEvent struct {
	ID        int64
	EventID   int64
	Operation string
	ProductID int64
	Stock     int32
	Price     float64
	UpdatedAt time.Time
}

type UserStatistics struct {
	UserID              int64
	TotalOrders         int
	CompletedOrders     int
	TotalItemsPurchased int
	AverageOrderValue   float64
	MostPurchasedItem   string
	HourlyDistribution  map[int]int
}
