package main

import (
	"log"

	"order-service/internal/server"
)

func main() {
	// Load configuration
	config := server.GetConfig()

	// Create and start the server
	srv := server.NewServer(config)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
