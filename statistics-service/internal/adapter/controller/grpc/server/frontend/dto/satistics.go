package dto

import (
	"time"

	"github.com/IrentyM/store/statistics-service/internal/domain/model"
)

type OrderEventDTO struct {
	EventID   int64          `json:"event_id"`
	Operation string         `json:"operation"`
	OrderID   int64          `json:"order_id"`
	UserID    int64          `json:"user_id"`
	Items     []OrderItemDTO `json:"items"`
	Total     float64        `json:"total"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
}

type OrderItemDTO struct {
	ProductID int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

func ToDomainOrderEvent(dto OrderEventDTO) model.OrderEvent {
	return model.OrderEvent{
		EventID:   dto.EventID,
		Operation: dto.Operation,
		OrderID:   dto.OrderID,
		UserID:    dto.UserID,
		Items:     ToDomainOrderItems(dto.Items),
		Total:     dto.Total,
		Status:    dto.Status,
		CreatedAt: dto.CreatedAt,
	}
}

type InventoryEventDTO struct {
	EventID   int64     `json:"event_id"`
	Operation string    `json:"operation"`
	ProductID int64     `json:"product_id"`
	Stock     int32     `json:"stock"`
	Price     float64   `json:"price"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDomainInventoryEvent(dto InventoryEventDTO) model.InventoryEvent {
	return model.InventoryEvent{
		EventID:   dto.EventID,
		Operation: dto.Operation,
		ProductID: dto.ProductID,
		Stock:     dto.Stock,
		Price:     dto.Price,
		UpdatedAt: dto.UpdatedAt,
	}
}

func ToDomainOrderItems(dtoItems []OrderItemDTO) []model.OrderItem {
	var items []model.OrderItem
	for _, dtoItem := range dtoItems {
		items = append(items, model.OrderItem{
			ProductID: dtoItem.ProductID,
			Quantity:  dtoItem.Quantity,
			Price:     dtoItem.Price,
		})
	}
	return items
}

type UserStatisticsDTO struct {
	UserID              int64   `json:"user_id"`
	TotalOrders         int     `json:"total_orders"`
	CompletedOrders     int     `json:"completed_orders"`
	TotalItemsPurchased int     `json:"total_items_purchased"`
	AverageOrderValue   float64 `json:"average_order_value"`
	MostPurchasedItem   string  `json:"most_purchased_item"`
}

func ToDomainUserStatistics(dto UserStatisticsDTO) model.UserStatistics {
	return model.UserStatistics{
		UserID:              dto.UserID,
		TotalOrders:         dto.TotalOrders,
		CompletedOrders:     dto.CompletedOrders,
		TotalItemsPurchased: dto.TotalItemsPurchased,
		AverageOrderValue:   dto.AverageOrderValue,
		MostPurchasedItem:   dto.MostPurchasedItem,
	}
}
