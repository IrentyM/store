package grpchandler

import (
	"inventory-service/internal/usecase"
	inventory "inventory-service/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer(ProductUsecase usecase.ProductUseCase) *Server {
	grpcServer := grpc.NewServer()
	handler := NewInventoryServer(ProductUsecase)
	inventory.RegisterInventoryServiceServer(grpcServer, handler)

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
