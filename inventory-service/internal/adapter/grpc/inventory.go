package grpc

import (
	"context"
	"inventory-service/internal/domain"
	"inventory-service/internal/usecase"
	inventory "inventory-service/proto"
)

type InventoryServer struct {
	inventory.UnimplementedInventoryServiceServer
	productUseCase usecase.ProductUseCase
}

func NewInventoryServer(productUseCase usecase.ProductUseCase) *InventoryServer {
	return &InventoryServer{productUseCase: productUseCase}
}

func (s *InventoryServer) CreateProduct(ctx context.Context, req *inventory.CreateProductRequest) (*inventory.ProductResponse, error) {
	product := domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryId,
	}

	if err := s.productUseCase.CreateProduct(ctx, product); err != nil {
		return nil, err
	}

	return &inventory.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CategoryId:  product.CategoryID,
		CreatedAt:   product.CreatedAt.String(),
		UpdatedAt:   product.UpdatedAt.String(),
	}, nil
}

func (s *InventoryServer) GetProductByID(ctx context.Context, req *inventory.GetProductRequest) (*inventory.ProductResponse, error) {
	product, err := s.productUseCase.GetProductByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, nil
	}

	return &inventory.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
		CategoryId:  int32(product.CategoryID),
		CreatedAt:   product.CreatedAt.String(),
		UpdatedAt:   product.UpdatedAt.String(),
	}, nil
}

func (s *InventoryServer) UpdateProduct(ctx context.Context, req *inventory.UpdateProductRequest) (*inventory.ProductResponse, error) {
	product := domain.Product{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryId,
	}

	if err := s.productUseCase.UpdateProduct(ctx, int(req.Id), product); err != nil {
		return nil, err
	}

	return &inventory.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
		CategoryId:  int32(product.CategoryID),
		CreatedAt:   product.CreatedAt.String(),
		UpdatedAt:   product.UpdatedAt.String(),
	}, nil
}

func (s *InventoryServer) DeleteProduct(ctx context.Context, req *inventory.DeleteProductRequest) (*inventory.Empty, error) {
	if err := s.productUseCase.DeleteProduct(ctx, int(req.Id)); err != nil {
		return nil, err
	}
	return &inventory.Empty{}, nil
}

func (s *InventoryServer) ListProducts(ctx context.Context, req *inventory.ListProductsRequest) (*inventory.ListProductsResponse, error) {
	products, err := s.productUseCase.ListProducts(ctx, nil, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var productResponses []*inventory.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, &inventory.ProductResponse{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       int32(product.Stock),
			CategoryId:  int32(product.CategoryID),
			CreatedAt:   product.CreatedAt.String(),
			UpdatedAt:   product.UpdatedAt.String(),
		})
	}

	return &inventory.ListProductsResponse{
		Products: productResponses,
		Total:    int32(len(productResponses)),
	}, nil
}
