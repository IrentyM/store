package repository

import (
	"context"
	"database/sql"
	"fmt"
	"inventory-service/internal/domain"
	"inventory-service/internal/repository/dao"
)

type productRepo struct {
	db    *sql.DB
	table string
}

const (
	product_table = "product"
)

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepo{
		db:    db,
		table: product_table,
	}
}

func (r *productRepo) Create(ctx context.Context, product domain.Product) error {
	object := dao.ToProduct(product)
	query := `
        INSERT INTO ` + r.table + ` (id, name, description, price, category_id, stock)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := r.db.ExecContext(ctx, query, object.ID, object.Name, object.Description, object.Price, object.CategoryID, object.Stock)
	return err
}

func (r *productRepo) GetByID(ctx context.Context, id int) (*domain.Product, error) {
	query := `
        SELECT id, name, description, price, category_id, stock, created_at, updated_at
        FROM ` + r.table + `
        WHERE id = $1
    `
	row := r.db.QueryRowContext(ctx, query, id)

	var product dao.Product
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CategoryID,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No product found
		}
		return nil, err
	}

	domainProduct := dao.FromProduct(product)
	return &domainProduct, nil
}

func (r *productRepo) Update(ctx context.Context, id int, product domain.Product) error {
	object := dao.ToProduct(product)
	query := `
        UPDATE ` + r.table + `
        SET name = $1, description = $2, price = $3, category_id = $4, stock = $5, updated_at = $6
        WHERE id = $7
    `
	_, err := r.db.ExecContext(ctx, query, object.Name, object.Description, object.Price, object.CategoryID, object.Stock, object.UpdatedAt, id)
	return err
}

func (r *productRepo) Delete(ctx context.Context, id int) error {
	query := `
        DELETE FROM ` + r.table + `
        WHERE id = $1
    `
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *productRepo) List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Product, error) {
	query := `
        SELECT id, name, description, price, category_id, stock, created_at, updated_at
        FROM ` + r.table + `
        WHERE 1=1
    `
	args := []interface{}{}
	argIndex := 1

	// Dynamically build the WHERE clause based on the filter
	for key, value := range filter {
		query += ` AND ` + key + ` = $` + fmt.Sprint(argIndex)
		args = append(args, value)
		argIndex++
	}

	query += ` LIMIT $` + fmt.Sprint(argIndex) + ` OFFSET $` + fmt.Sprint(argIndex+1)
	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var product dao.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CategoryID,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		domainProduct := dao.FromProduct(product)
		products = append(products, &domainProduct)
	}

	return products, nil
}
