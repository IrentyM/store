package dto

import (
	"inventory-service/internal/domain"
)

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateCategoryRequest) ToDomain() domain.Category {
	return domain.Category{
		Name:       r.Name,
		Dscription: r.Description,
	}
}

type CategoryResponse struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewCategoryResponse(category domain.Category) CategoryResponse {
	return CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Dscription,
	}
}
