package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserID    uuid.UUID `json:"userID"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	TeamID    uuid.UUID `json:"teamID"`
}
