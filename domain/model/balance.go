package model

type Balance struct {
	ID              int    `gorm:"primary_key" json:"id"`
	Amount          string `gorm:"column:name; not null" json:"name"`
	UserID          int    `gorm:"column:user_id" json:"user_id"`
	BalanceStatusID int    `gorm:"column:balance_status_id" json:"balance_status_id"`
}

func (Balance) TableName() string { return "balances" }
