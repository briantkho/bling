package account

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Account struct {
	AccountID      string          `json:"account_id"`
	AccountName    string          `json:"account_name"`
	AccountType    string          `json:"account_type"`
	CurrentBalance decimal.Decimal `json:"current_balance"`
	UserID         uuid.UUID       `json:"user_id"`
}
