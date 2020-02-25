package model

import "time"

type User struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string    `gorm:"column:name; not null" json:"first_name"`
	Description  string    `gorm:"column:description; null" json:"description"`
	Email        string    `gorm:"column:email; not null" json:"email"`
	Phone        string    `gorm:"column:phone; not null" json:"phone"`
	Password     string    `gorm:"column:password; not null" json:"password"`
	UserRoleID   int       `gorm:"column:user_role_id" json:"user_role_id"`
	Events       []Event   `gorm:"PRELOAD:false" json:"events"`
	TopUpHistory []Balance `json:"balance"`
	BalanceNow   int       `gorm:"-" json:"balance"`
	CreatedAt    time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	DeletedAt    time.Time `gorm:"default:null" json:"created_at"`
}
