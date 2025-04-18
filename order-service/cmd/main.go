package main

import (
	"log"

	"order-service/internal/app"
	"order-service/internal/server"
)

func main() {
	config := server.GetConfig()

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
