package levelpointstransaction

import (
	"github.com/google/uuid"
	"time"
)

type LevelPointTransaction struct {
	LevelPointTransactionID int       `json:"levelPointTransactionID"`
	Points                  int       `json:"points"`
	ActionType              int       `json:"actionType"`
	DateEarned              time.Time `json:"dateEarned"`
	UserID                  uuid.UUID `json:"userId"`
}
