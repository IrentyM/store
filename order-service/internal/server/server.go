package server

import (
	"log"

	handler "order-service/internal/delivery/http"
	"order-service/internal/repository"
	"order-service/internal/usecase"
	"order-service/pkg/db"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Start() error
}

type server struct {
	router *gin.Engine
	cfg    *Config
}

func NewServer(cfg *Config) Server {
	r := gin.Default()

	return &server{
		router: r,
		cfg:    cfg,
	}
}

func (s *server) Start() error {
	// Initialize database connection
	database, err := db.PostgresConnection(s.cfg.DBhost, s.cfg.DBport, s.cfg.DBuser, s.cfg.DBpassword, s.cfg.DBname)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return err
	}
	defer database.Close()

	// Initialize repositories
	orderRepo := repository.NewOrderRepository(database)
	orderItemRepo := repository.NewOrderItemRepository(database)

	// Initialize use cases
	orderUseCase := usecase.NewOrderUseCase(orderRepo, orderItemRepo)

	// Register routes
	s.router.GET("/health", handler.GetHealth)

	orderHandler := handler.NewOrderHandler(orderUseCase)
	s.router.POST("/orders", orderHandler.CreateOrder)
	s.router.GET("/orders/:id", orderHandler.GetOrderByID)
	s.router.PUT("/orders/:id", orderHandler.UpdateOrder)
	s.router.DELETE("/orders/:id", orderHandler.DeleteOrder)
	s.router.GET("/orders", orderHandler.ListOrders)

	// Start the server
	log.Printf("Starting server on port %s...", s.cfg.Port)
	if err := s.router.Run(s.cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return err
	}

	return nil
}
