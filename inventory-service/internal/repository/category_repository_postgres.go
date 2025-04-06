package repository

import (
	"context"
	"database/sql"
	"fmt"
	"inventory-service/internal/domain"
)

type categoryRepo struct {
	db    *sql.DB
	table string
}

const (
	category_table = "categories"
)

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepo{
		db:    db,
		table: category_table,
	}
}

func (r *categoryRepo) Create(ctx context.Context, category domain.Category) error {
	query := `
        INSERT INTO ` + r.table + ` (id, name, description)
        VALUES ($1, $2, $3)
    `
	_, err := r.db.ExecContext(ctx, query, category.ID, category.Name, category.Dscription)
	return err
}

func (r *categoryRepo) GetByID(ctx context.Context, id int) (*domain.Category, error) {
	query := `
        SELECT id, name, description
        FROM ` + r.table + `
        WHERE id = $1
    `
	row := r.db.QueryRowContext(ctx, query, id)

	var category domain.Category
	err := row.Scan(&category.ID, &category.Name, &category.Dscription)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No category found
		}
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepo) Update(ctx context.Context, id int, category domain.Category) error {
	query := `
        UPDATE ` + r.table + `
        SET name = $1, description = $2
        WHERE id = $3
    `
	_, err := r.db.ExecContext(ctx, query, category.Name, category.Dscription, id)
	return err
}

func (r *categoryRepo) Delete(ctx context.Context, id int) error {
	query := `
        DELETE FROM ` + r.table + `
        WHERE id = $1
    `
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *categoryRepo) List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*domain.Category, error) {
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
		err := rows.Scan(&category.ID, &category.Name, &category.Dscription)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}
