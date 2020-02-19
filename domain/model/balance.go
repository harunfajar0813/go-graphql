package model

type Balance struct {
	ID              int `gorm:"primary_key" json:"id"`
	Amount          int `gorm:"column:amount; not null" json:"amount"`
	UserID          int `gorm:"column:user_id; null" json:"user_id"`
}

func (Balance) TableName() string { return "balances" }
