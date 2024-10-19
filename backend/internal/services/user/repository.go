package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUserByID(userID uuid.UUID) (*User, error) {
	query := `
		SELECT u.first_name, u.last_name, u.email, u.username, 
			   u.created_at, u.current_level_points, 
			   u.total_lifetime_points, ul.level_number, ul.points_required
		FROM user u
		JOIN user_level ul ON u.level_id = ul.level_id
		WHERE u.user_id = ?
	`

	var user User
	err := r.db.QueryRow(query, userID).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Username,
		&user.CreatedAt,
		&user.CurrentLevelPoints,
		&user.TotalLifetimePoints,
		&user.LevelNumber,
		&user.PointsRequired,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

func (r *Repository) CreateUser(user *User) (uuid.UUID, error) {
	defaultStartingLevelId := 1

	query := `
INSERT INTO user (user_id, first_name, last_name, username, email, password_hash,  level_id)
VALUES (?, ?, ?, ?, ?, ?, ?) `

	_, err := r.db.Exec(query, user.UserID, user.FirstName, user.LastName, user.Username, user.Email, user.Password, defaultStartingLevelId)

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user.UserID, nil
}
