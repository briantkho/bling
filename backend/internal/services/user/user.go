package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserID              uuid.UUID `json:"userID"`
	FirstName           string    `json:"firstName"`
	LastName            string    `json:"lastName"`
	Email               string    `json:"email"`
	Password            string    `json:"password"`
	CreatedAt           time.Time `json:"createdAt"`
	CurrentLevelPoints  int       `json:"currentLevelPoints"`
	TotalLifetimePoints int       `json:"totalLifetimePoints"`
	LevelID             int       `json:"levelId"`
}
