package dao

import (
	"time"

	"github.com/IrentyM/store/statistics-service/internal/domain/model"
)

type OrderEvent struct {
	ID        int64     `db:"id"`
	EventID   int64     `db:"event_id"`
	Operation string    `db:"operation"`
	OrderID   int64     `db:"order_id"`
	UserID    int64     `db:"user_id"`
	Total     float64   `db:"total"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type OrderItem struct {
	ProductID int64   `db:"product_id"`
	Quantity  int     `db:"quantity"`
	Price     float64 `db:"price"`
}

type InventoryEvent struct {
	ID        int64     `db:"id"`
	EventID   int64     `db:"event_id"`
	Operation string    `db:"operation"`
	ProductID int64     `db:"product_id"`
	Stock     int32     `db:"stock"`
	Price     float64   `db:"price"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserStatistics struct {
	UserID              int64       `db:"user_id"`
	TotalOrders         int         `db:"total_orders"`
	CompletedOrders     int         `db:"completed_orders"`
	TotalItemsPurchased int         `db:"total_items_purchased"`
	AverageOrderValue   float64     `db:"average_order_value"`
	MostPurchasedItem   string      `db:"most_purchased_item"`
	HourlyDistribution  map[int]int `db:"hourly_distribution"`
}

func FromModelOrderEvent(event *model.OrderEvent) OrderEvent {
	return OrderEvent{
		ID:        event.ID,
		EventID:   event.EventID,
		Operation: event.Operation,
		OrderID:   event.OrderID,
		UserID:    event.UserID,
		Total:     event.Total,
		Status:    event.Status,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}
}

func ToModelOrderEvent(event OrderEvent, items []model.OrderItem) *model.OrderEvent {
	return &model.OrderEvent{
		ID:        event.ID,
		EventID:   event.EventID,
		Operation: event.Operation,
		OrderID:   event.OrderID,
		UserID:    event.UserID,
		Items:     items,
		Total:     event.Total,
		Status:    event.Status,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}
}

func FromModelOrderItem(item model.OrderItem) OrderItem {
	return OrderItem{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		Price:     item.Price,
	}
}

func ToModelOrderItem(item OrderItem) model.OrderItem {
	return model.OrderItem{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		Price:     item.Price,
	}
}

func FromModelInventoryEvent(event *model.InventoryEvent) InventoryEvent {
	return InventoryEvent{
		ID:        event.ID,
		EventID:   event.EventID,
		Operation: event.Operation,
		ProductID: event.ProductID,
		Stock:     event.Stock,
		Price:     event.Price,
		UpdatedAt: event.UpdatedAt,
	}
}

func ToModelInventoryEvent(event InventoryEvent) *model.InventoryEvent {
	return &model.InventoryEvent{
		ID:        event.ID,
		EventID:   event.EventID,
		Operation: event.Operation,
		ProductID: event.ProductID,
		Stock:     event.Stock,
		Price:     event.Price,
		UpdatedAt: event.UpdatedAt,
	}
}

func FromModelUserStatistics(stats *model.UserStatistics) UserStatistics {
	return UserStatistics{
		UserID:              stats.UserID,
		TotalOrders:         stats.TotalOrders,
		CompletedOrders:     stats.CompletedOrders,
		TotalItemsPurchased: stats.TotalItemsPurchased,
		AverageOrderValue:   stats.AverageOrderValue,
		MostPurchasedItem:   stats.MostPurchasedItem,
		HourlyDistribution:  stats.HourlyDistribution,
	}
}

func ToModelUserStatistics(stats UserStatistics) *model.UserStatistics {
	return &model.UserStatistics{
		UserID:              stats.UserID,
		TotalOrders:         stats.TotalOrders,
		CompletedOrders:     stats.CompletedOrders,
		TotalItemsPurchased: stats.TotalItemsPurchased,
		AverageOrderValue:   stats.AverageOrderValue,
		MostPurchasedItem:   stats.MostPurchasedItem,
		HourlyDistribution:  stats.HourlyDistribution,
	}
}
