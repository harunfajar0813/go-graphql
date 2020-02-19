package model

type Transaction struct {
	ID       int `gorm:"primary_key" json:"id"`
	ClientID int `gorm:"column:client_id" json:"user_id"`
	EventID  int `gorm:"column:event_id" json:"event_id"`
}
