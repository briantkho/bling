package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

func InitDB() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		return nil, errors.New("DATABASE_URL environment variable not set")
	}

	db, err := sql.Open("mysql", dbURL)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database")

	return db, nil
}
