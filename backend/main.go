package main

import (
	"database/sql"
	"github.com/briantkho/bling/config"
	"github.com/briantkho/bling/internal/database"
	"github.com/briantkho/bling/internal/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := database.InitDB()

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing database: %v", err)
		}
	}(db)

	s := server.NewServer(db)
	s.SetupRoutes()

	if err := s.Run(cfg.Server.Address); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
