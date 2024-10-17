package user

import "database/sql"

func InitUserDependencies(db *sql.DB) *Handler {
	userRepo := NewRepository(db)
	userService := NewService(userRepo)
	userHandler := NewHandler(userService)

	return userHandler
}
