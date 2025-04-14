package grpchandler

import (
	"inventory-service/internal/usecase"
	categoryproto "inventory-service/proto/category"
	productproto "inventory-service/proto/product"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer(productUsecase usecase.ProductUseCase, categoryUsecase usecase.CategoryUseCase) *Server {
	grpcServer := grpc.NewServer()

	// Register ProductService
	productHandler := NewProductServer(productUsecase)
	productproto.RegisterProductServiceServer(grpcServer, productHandler)

	// Register CategoryService
	categoryHandler := NewCategoryServer(categoryUsecase)
	categoryproto.RegisterCategoryServiceServer(grpcServer, categoryHandler)

	// Enable reflection
	reflection.Register(grpcServer)

	return &Server{
		grpcServer: grpcServer,
	}
}

func (s *Server) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	log.Printf("gRPC server listening on port %s", port)
	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
