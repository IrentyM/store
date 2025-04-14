package grpchandler

import (
	"context"
	"inventory-service/internal/domain"
	"inventory-service/internal/usecase"
	categoryproto "inventory-service/proto/category"
)

type CategoryServer struct {
	categoryproto.UnimplementedCategoryServiceServer
	categoryUseCase usecase.CategoryUseCase
}

func NewCategoryServer(categoryUseCase usecase.CategoryUseCase) *CategoryServer {
	return &CategoryServer{categoryUseCase: categoryUseCase}
}

func (s *CategoryServer) CreateCategory(ctx context.Context, req *categoryproto.CreateCategoryRequest) (*categoryproto.CategoryResponse, error) {
	category := domain.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.categoryUseCase.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return &categoryproto.CategoryResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (s *CategoryServer) GetCategoryByID(ctx context.Context, req *categoryproto.GetCategoryRequest) (*categoryproto.CategoryResponse, error) {
	category, err := s.categoryUseCase.GetCategoryByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &categoryproto.CategoryResponse{
		Id:          int32(category.ID),
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (s *CategoryServer) UpdateCategory(ctx context.Context, req *categoryproto.UpdateCategoryRequest) (*categoryproto.CategoryResponse, error) {
	category := domain.Category{
		ID:          int(req.Id),
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.categoryUseCase.UpdateCategory(ctx, category.ID, category)
	if err != nil {
		return nil, err
	}

	return &categoryproto.CategoryResponse{
		Id:          int32(category.ID),
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (s *CategoryServer) DeleteCategory(ctx context.Context, req *categoryproto.DeleteCategoryRequest) (*categoryproto.Empty, error) {
	err := s.categoryUseCase.DeleteCategory(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &categoryproto.Empty{}, nil
}

func (s *CategoryServer) ListCategories(ctx context.Context, req *categoryproto.ListCategoriesRequest) (*categoryproto.ListCategoriesResponse, error) {
	categories, err := s.categoryUseCase.ListCategories(ctx, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var categoryResponses []*categoryproto.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, &categoryproto.CategoryResponse{
			Id:          int32(category.ID),
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &categoryproto.ListCategoriesResponse{
		Categories: categoryResponses,
		Total:      int32(len(categoryResponses)),
	}, nil
}
