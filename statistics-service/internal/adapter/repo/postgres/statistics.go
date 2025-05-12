package statisticsrepo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/IrentyM/store/statistics-service/internal/adapter/repo/postgres/dao"
	"github.com/IrentyM/store/statistics-service/internal/domain/model"
)

type statsRepository struct {
	db *sql.DB
}

func NewStatsRepository(db *sql.DB) *statsRepository {
	return &statsRepository{db: db}
}

func (r *statsRepository) SaveOrderEvent(ctx context.Context, event *model.OrderEvent) error {
	query := `
        INSERT INTO order_events (event_id, operation, order_id, user_id, total, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `
	_, err := r.db.ExecContext(ctx, query, event.EventID, event.Operation, event.OrderID, event.UserID, event.Total, event.Status, event.CreatedAt, event.UpdatedAt)
	return err
}

func (r *statsRepository) SaveInventoryEvent(ctx context.Context, event *model.InventoryEvent) error {
	query := `
        INSERT INTO inventory_events (event_id, operation, product_id, stock, price, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := r.db.ExecContext(ctx, query, event.EventID, event.Operation, event.ProductID, event.Stock, event.Price, event.UpdatedAt)
	return err
}

func (r *statsRepository) GetUserOrderStats(ctx context.Context, userID int64) ([]model.OrderEvent, error) {
	query := `
        SELECT id, event_id, operation, order_id, user_id, total, status, created_at, updated_at
        FROM order_events
        WHERE user_id = $1
    `
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.OrderEvent
	for rows.Next() {
		var event dao.OrderEvent
		if err := rows.Scan(&event.ID, &event.EventID, &event.Operation, &event.OrderID, &event.UserID, &event.Total, &event.Status, &event.CreatedAt, &event.UpdatedAt); err != nil {
			return nil, err
		}
		events = append(events, *dao.ToModelOrderEvent(event, nil))
	}

	return events, nil
}

func (r *statsRepository) GetUserOrderStatistics(ctx context.Context, userID int64) (*model.UserStatistics, error) {
	query := `
        SELECT items.product_id, SUM(items.quantity) AS total_quantity, COUNT(*) AS total_orders, SUM(events.total) AS total_spent
        FROM order_events AS events
        JOIN order_items AS items ON events.order_id = items.order_id
        WHERE events.user_id = $1 AND events.status = 'completed'
        GROUP BY items.product_id
        ORDER BY total_quantity DESC
        LIMIT 1
    `
	row := r.db.QueryRowContext(ctx, query, userID)

	var result struct {
		ProductID   string
		TotalQty    int
		TotalOrders int
		TotalSpent  float64
	}

	stats := &model.UserStatistics{
		UserID: userID,
	}

	if err := row.Scan(&result.ProductID, &result.TotalQty, &result.TotalOrders, &result.TotalSpent); err != nil {
		if err == sql.ErrNoRows {
			return stats, nil // No data found, return empty stats
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	stats.MostPurchasedItem = result.ProductID
	stats.TotalItemsPurchased = result.TotalQty
	if result.TotalOrders > 0 {
		stats.AverageOrderValue = result.TotalSpent / float64(result.TotalOrders)
	}

	// Get total orders
	query = `
        SELECT COUNT(*)
        FROM order_events
        WHERE user_id = $1
    `
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&stats.TotalOrders)
	if err != nil {
		return nil, fmt.Errorf("failed to get total orders: %w", err)
	}

	return stats, nil
}
