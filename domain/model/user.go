package model

import "time"

type User struct {
	ID           int           `gorm:"primary_key" json:"id"`
	Name         string        `gorm:"column:name; not null" json:"first_name"`
	Description  string        `gorm:"column:description; null" json:"description"`
	Email        string        `gorm:"column:email; not null" json:"email"`
	Phone        string        `gorm:"column:phone; not null" json:"phone"`
	Password     string        `gorm:"column:password; not null" json:"password"`
	Balance      string        `gorm:"-" json:"balance"`
	Events       []Event       `gorm:"PRELOAD:false" json:"events"`
	CreatedAt    time.Time     `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt    time.Time     `gorm:"default:current_timestamp()" json:"created_at"`
	DeletedAt    time.Time     `gorm:"default:null" json:"created_at"`
}
