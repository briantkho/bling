package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func InitDB(dbURL string) (*sql.DB, error) {
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

func MigrateDB(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Printf("failed to create database driver: %v", err)
	}

	migration, err := migrate.NewWithDatabaseInstance("file://internal/database/migrations", "mysql", driver)
	if err != nil {
		log.Fatalf("failed to create migration instance: %v", err)
	}

	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("Successfully migrated database")
}
