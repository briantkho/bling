package main

import (
	"github.com/briantkho/bling/config"
	"github.com/briantkho/bling/internal/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	router := server.NewServer()

	if err := router.Run(cfg.Server.Address); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
