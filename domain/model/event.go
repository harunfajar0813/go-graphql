package model

import "time"

type Event struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Address     string    `gorm:"column:address" json:"address"`
	StartEvent  string    `gorm:"column:start_event" json:"start_event"`
	Price       int       `gorm:"column:price" json:"price"`
	Stock       int       `gorm:"column:stock" json:"stock"`
	UserID      int       `gorm:"column:user_id" json:"user_id"`
	CreatedAt   time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	DeletedAt   time.Time `gorm:"default:null" json:"created_at"`
}