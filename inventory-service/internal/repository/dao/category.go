package dao

import (
	"inventory-service/internal/domain"
)

type Category struct {
	ID          int32  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

func ToCategory(category domain.Category) Category {
	return Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Dscription,
	}
}

func FromCategory(category Category) domain.Category {
	return domain.Category{
		ID:         category.ID,
		Name:       category.Name,
		Dscription: category.Description,
	}
}
