package model

type BalanceStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (BalanceStatus) TableName() string { return "balances_status" }
