package grpchandler

import (
	"context"
	"inventory-service/internal/domain"
	"inventory-service/internal/usecase"
	productproto "inventory-service/proto/product"
)

type ProductServer interface {
	CreateProduct(ctx context.Context, req *productproto.CreateProductRequest) (*productproto.ProductResponse, error)
	GetProductByID(ctx context.Context, req *productproto.GetProductRequest) (*productproto.ProductResponse, error)
	UpdateProduct(ctx context.Context, req *productproto.UpdateProductRequest) (*productproto.ProductResponse, error)
	DeleteProduct(ctx context.Context, req *productproto.DeleteProductRequest) (*productproto.Empty, error)
	ListProducts(ctx context.Context, req *productproto.ListProductsRequest) (*productproto.ListProductsResponse, error)
}

type productServer struct {
	productproto.UnimplementedProductServiceServer
	productUseCase usecase.ProductUseCase
}

func NewProductServer(productUseCase usecase.ProductUseCase) *productServer {
	return &productServer{productUseCase: productUseCase}
}

func (s *productServer) CreateProduct(ctx context.Context, req *productproto.CreateProductRequest) (*productproto.ProductResponse, error) {
	product := domain.Product{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryId,
	}

	if err := s.productUseCase.CreateProduct(ctx, product); err != nil {
		return nil, err
	}

	return &productproto.ProductResponse{
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

func (s *productServer) GetProductByID(ctx context.Context, req *productproto.GetProductRequest) (*productproto.ProductResponse, error) {
	product, err := s.productUseCase.GetProductByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, nil
	}

	return &productproto.ProductResponse{
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

func (s *productServer) UpdateProduct(ctx context.Context, req *productproto.UpdateProductRequest) (*productproto.ProductResponse, error) {
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

	return &productproto.ProductResponse{
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

func (s *productServer) DeleteProduct(ctx context.Context, req *productproto.DeleteProductRequest) (*productproto.Empty, error) {
	if err := s.productUseCase.DeleteProduct(ctx, int(req.Id)); err != nil {
		return nil, err
	}
	return &productproto.Empty{}, nil
}

func (s *productServer) ListProducts(ctx context.Context, req *productproto.ListProductsRequest) (*productproto.ListProductsResponse, error) {
	products, err := s.productUseCase.ListProducts(ctx, nil, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var productResponses []*productproto.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, &productproto.ProductResponse{
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

	return &productproto.ListProductsResponse{
		Products: productResponses,
		Total:    int32(len(productResponses)),
	}, nil
}
