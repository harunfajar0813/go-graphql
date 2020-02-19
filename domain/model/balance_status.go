package model

type BalanceStatus struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"column:name; not null" json:"name"`
}

func (BalanceStatus) TableName() string { return "balances_status" }
