package main

import (
	"github.com/briantkho/bling/internal/database"
	"github.com/briantkho/bling/internal/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	address := os.Getenv("ADDRESS")

	db, err := database.InitDB(dbURL)

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal("Error closing database connection: ", err)
		}
		log.Println("Database connection closed")
	}()

	database.MigrateDB(db)

	s := server.NewServer(db)
	s.SetupRoutes()

	if err := s.Run(address); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
