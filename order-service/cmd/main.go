package main

import (
	"log"

	"order-service/internal/server"
)

func main() {
	config := server.GetConfig()

	srv := server.NewServer(config)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
