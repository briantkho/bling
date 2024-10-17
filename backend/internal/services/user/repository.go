package user

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUserByID(userID uuid.UUID) (User, error) {
	return User{
		UserID:              userID,
		FirstName:           "",
		LastName:            "",
		Username:            "",
		Email:               "",
		Password:            "",
		CreatedAt:           time.Time{},
		CurrentLevelPoints:  0,
		TotalLifetimePoints: 0,
		LevelID:             0,
	}, nil
}

func (r *Repository) CreateUser(user *User) uuid.UUID {
	return user.UserID
}
