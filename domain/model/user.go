package model

import "time"

type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Events    []Event   `gorm:"PRELOAD:false" json:"events"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	DeletedAt time.Time `gorm:"default:null" json:"created_at"`
}

func (User) TableName() string { return "users" }
