package grpchandler

import (
	"log"
	"net"

	"order-service/internal/usecase"
	orderproto "order-service/proto/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer(orderUseCase usecase.OrderUseCase) *Server {
	grpcServer := grpc.NewServer()

	// Register OrderService
	orderHandler := NewOrderServer(orderUseCase)
	orderproto.RegisterOrderServiceServer(grpcServer, orderHandler)

	// Enable reflection for debugging
	reflection.Register(grpcServer)

	return &Server{
		grpcServer: grpcServer,
	}
}

func (s *Server) Run(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	log.Printf("gRPC server is running on port %s", port)
	return s.grpcServer.Serve(listener)
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
