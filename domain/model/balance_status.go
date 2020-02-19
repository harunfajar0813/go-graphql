package model

type BalanceStatus struct {
	ID       int       `gorm:"primary_key" json:"id"`
	Name     string    `gorm:"column:name; not null" json:"name"`
	Balances []Balance `gorm:"PRELOAD:false" json:"balances"`
}

func (BalanceStatus) TableName() string { return "balances_status" }
