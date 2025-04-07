package server

import (
	handler "order-service/internal/delivery/http"
	"order-service/internal/repository"
	"order-service/internal/usecase"
	"order-service/pkg/db"
)

func (s *server) registerRoutes() error {
	// Initialize database connection
	database, err := db.PostgresConnection(s.cfg.DBhost, s.cfg.DBport, s.cfg.DBuser, s.cfg.DBpassword, s.cfg.DBname)
	if err != nil {
		return err
	}

	// Initialize repositories
	orderRepo := repository.NewOrderRepository(database)
	orderItemRepo := repository.NewOrderItemRepository(database)

	// Initialize use cases
	orderUseCase := usecase.NewOrderUseCase(orderRepo, orderItemRepo)

	// Register health check route
	s.router.GET("/health", handler.GetHealth)

	// Register order routes
	orderHandler := handler.NewOrderHandler(orderUseCase)
	s.router.POST("/orders", orderHandler.CreateOrder)
	s.router.GET("/orders/:id", orderHandler.GetOrderByID)
	s.router.PUT("/orders/:id", orderHandler.UpdateOrder)
	s.router.DELETE("/orders/:id", orderHandler.DeleteOrder)
	s.router.GET("/orders", orderHandler.ListOrders)

	return nil
}
