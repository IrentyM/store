package repository

import (
	"context"
	"database/sql"
	"order-service/internal/domain"
	"order-service/internal/repository/dao"
)

type orderItemRepository struct {
	db    *sql.DB
	table string
}

const (
	orderItemTable = "orders.order_items"
)

func NewOrderItemRepository(db *sql.DB) *orderItemRepository {
	return &orderItemRepository{
		db:    db,
		table: orderItemTable,
	}
}

func (r *orderItemRepository) Create(ctx context.Context, orderItem domain.OrderItem) error {
	object := dao.ToOrderItem(orderItem)
	query := `
        INSERT INTO ` + r.table + ` (order_id, product_id, quantity, price_at_purchase)
        VALUES ($1, $2, $3, $4)
    `
	_, err := r.db.ExecContext(ctx, query, object.OrderID, object.ProductID, object.Quantity, object.PriceAtPurchase)
	return err
}

func (r *orderItemRepository) GetByOrderID(ctx context.Context, orderID int) ([]domain.OrderItem, error) {
	query := `
        SELECT id, order_id, product_id, quantity, price_at_purchase
        FROM ` + r.table + `
        WHERE order_id = $1
    `
	rows, err := r.db.QueryContext(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderItems []domain.OrderItem
	for rows.Next() {
		var item dao.OrderItem
		if err := rows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Quantity, &item.PriceAtPurchase); err != nil {
			return nil, err
		}
		orderItems = append(orderItems, dao.FromOrderItem(item))
	}

	return orderItems, nil
}

func (r *orderItemRepository) DeleteByOrderID(ctx context.Context, orderID int) error {
	query := `
        DELETE FROM ` + r.table + `
        WHERE order_id = $1
    `
	_, err := r.db.ExecContext(ctx, query, orderID)
	return err
}
