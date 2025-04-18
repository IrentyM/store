package main

import (
	"log"

	"order-service/internal/app"
	"order-service/internal/server"
)

func main() {
	// Load configuration
	config := server.GetConfig()

	// Connect to the database
	// database, err := db.PostgresConnection(config.DBhost, config.DBport, config.DBuser, config.DBpassword, config.DBname)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the database: %v", err)
	// }

	// // Initialize repositories and use cases
	// orderRepo := repository.NewOrderRepository(database)
	// orderItemRepo := repository.NewOrderItemRepository(database)
	// orderUseCase := usecase.NewOrderUseCase(orderRepo, orderItemRepo)

	// Start the gRPC server
	application, err := app.New(config)
	if err != nil {
		log.Printf("failed to setup application: %v", err)
		return
	}

	if err := application.Run(config.Port); err != nil {
		log.Printf("failed to run application: %v", err)
		return
	}
}
