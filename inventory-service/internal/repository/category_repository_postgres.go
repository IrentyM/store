package repository

import (
	"context"
	"database/sql"
	"fmt"
	"inventory-service/internal/domain"
)

type categoryRepository struct {
	db    *sql.DB
	table string
}

const (
	category_table = "inventory.categories"
)

func NewCategoryRepository(db *sql.DB) *categoryRepository {
	return &categoryRepository{
		db:    db,
		table: category_table,
	}
}

func (r *categoryRepository) Create(ctx context.Context, category domain.Category) error {
	query := `
        INSERT INTO ` + r.table + ` (id, name, description)
        VALUES ($1, $2, $3)
    `
	_, err := r.db.ExecContext(ctx, query, category.ID, category.Name, category.Description)
	return err
}

func (r *categoryRepository) GetByID(ctx context.Context, id int) (*domain.Category, error) {
	query := `
        SELECT id, name, description
        FROM ` + r.table + `
        WHERE id = $1
    `
	row := r.db.QueryRowContext(ctx, query, id)

	var category domain.Category
	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No category found
		}
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) Update(ctx context.Context, id int32, category domain.Category) error {
	query := `
        UPDATE ` + r.table + `
        SET name = $1, description = $2
        WHERE id = $3
    `
	_, err := r.db.ExecContext(ctx, query, category.Name, category.Description, id)
	return err
}

func (r *categoryRepository) Delete(ctx context.Context, id int) error {
	query := `
        DELETE FROM ` + r.table + `
        WHERE id = $1
    `
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *categoryRepository) List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Category, error) {
	query := `
        SELECT id, name, description
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

	var categories []*domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}
