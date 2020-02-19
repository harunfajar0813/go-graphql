package model

import "time"

type Client struct {
	ID           int           `gorm:"primary_key" json:"id"`
	FirstName    string        `gorm:"column:first_name" json:"first_name"`
	LastName     string        `gorm:"column:last_name" json:"last_name"`
	Email        string        `gorm:"column:email" json:"email"`
	Phone        string        `gorm:"column:phone" json:"phone"`
	Password     string        `gorm:"column:password" json:"password"`
	Balance      string        `json:"balance"`
	Transactions []Transaction `gorm:"PRELOAD:false" json:"transactions"`
	CreatedAt    time.Time     `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt    time.Time     `gorm:"default:current_timestamp()" json:"created_at"`
	DeletedAt    time.Time     `gorm:"default:null" json:"created_at"`
}
