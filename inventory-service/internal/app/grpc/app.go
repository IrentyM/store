package grpcapp

import (
	"log"
	"net"

	grpchandler "inventory-service/internal/adapter/grpc"
	"inventory-service/internal/usecase"
	categoryproto "inventory-service/proto/category"
	productproto "inventory-service/proto/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	categoryUseCase usecase.CategoryUseCase
	productUseCase  usecase.ProductUseCase
	port            string
}

func NewApp(categoryUseCase usecase.CategoryUseCase, productUseCase usecase.ProductUseCase, port string) *App {
	return &App{
		categoryUseCase: categoryUseCase,
		productUseCase:  productUseCase,
		port:            port,
	}
}

func (a *App) Run() error {
	grpcServer := grpc.NewServer()

	// Register CategoryService
	categoryHandler := grpchandler.NewCategoryServer(a.categoryUseCase)
	categoryproto.RegisterCategoryServiceServer(grpcServer, categoryHandler)

	// Register ProductService
	productHandler := grpchandler.NewProductServer(a.productUseCase)
	productproto.RegisterProductServiceServer(grpcServer, productHandler)

	// Enable reflection for debugging
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", a.port)
	if err != nil {
		return err
	}

	log.Printf("gRPC server is running on port %s", a.port)
	return grpcServer.Serve(listener)
}
