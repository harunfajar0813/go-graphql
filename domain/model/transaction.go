package model

type Transaction struct {
	ID      uint `json:"id"`
	UserID  uint `json:"user_id"`
	EventID uint `json:"event_id"`
}

func (Transaction) TableName() string { return "transactions" }
