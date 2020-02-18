package model

type Balance struct {
	ID              int `json:"id"`
	Amount          int `json:"amount"`
	BalanceStatusID int `json:"balance_status_id"`
	UserID          int `json:"user_id"`
}

func (Balance) TableName() string { return "balances" }
