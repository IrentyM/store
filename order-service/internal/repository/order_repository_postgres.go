package repository

import (
	"context"
	"database/sql"
	"order-service/internal/domain"
	"order-service/internal/repository/dao"
	"strconv"
)

type orderRepository struct {
	db    *sql.DB
	table string
}

const (
	orderTable = "orders.orders"
)

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{
		db:    db,
		table: orderTable,
	}
}

func (r *orderRepository) Create(ctx context.Context, order domain.Order) (int, error) {
	object := dao.ToOrder(order)
	query := `
        INSERT INTO ` + r.table + ` (user_id, status, payment_status, total_amount)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	var id int
	err := r.db.QueryRowContext(ctx, query, object.UserID, object.Status, object.PaymentStatus, object.TotalAmount).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *orderRepository) GetByID(ctx context.Context, id int) (*domain.Order, error) {
	query := `
        SELECT id, user_id, status, payment_status, total_amount, created_at, updated_at
        FROM ` + r.table + `
        WHERE id = $1
    `
	row := r.db.QueryRowContext(ctx, query, id)

	var order dao.Order
	err := row.Scan(
		&order.ID,
		&order.UserID,
		&order.Status,
		&order.PaymentStatus,
		&order.TotalAmount,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No order found
		}
		return nil, err
	}

	domainOrder := dao.FromOrder(order)
	return &domainOrder, nil
}

func (r *orderRepository) Update(ctx context.Context, id int, order domain.Order) error {
	object := dao.ToOrder(order)
	query := `
        UPDATE ` + r.table + `
        SET user_id = $1, status = $2, payment_status = $3, total_amount = $4, updated_at = CURRENT_TIMESTAMP
        WHERE id = $5
    `
	_, err := r.db.ExecContext(ctx, query, object.UserID, object.Status, object.PaymentStatus, object.TotalAmount, id)
	return err
}

func (r *orderRepository) Delete(ctx context.Context, id int) error {
	query := `
        DELETE FROM ` + r.table + `
        WHERE id = $1
    `
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *orderRepository) List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]domain.Order, error) {
	query := `
        SELECT id, user_id, status, payment_status, total_amount, created_at, updated_at
        FROM ` + r.table + `
        WHERE 1=1
    `
	args := []interface{}{}
	argIndex := 1

	// Dynamically build the WHERE clause based on the filter
	for key, value := range filter {
		query += ` AND ` + key + ` = $` + strconv.Itoa(argIndex)
		args = append(args, value)
		argIndex++
	}

	query += ` LIMIT $` + strconv.Itoa(argIndex) + ` OFFSET $` + strconv.Itoa(argIndex+1)
	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order dao.Order
		err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.Status,
			&order.PaymentStatus,
			&order.TotalAmount,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, dao.FromOrder(order))
	}

	return orders, nil
}
