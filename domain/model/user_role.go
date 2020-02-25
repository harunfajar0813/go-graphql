package model

import "time"

type UserRole struct {
	ID        int       `grom:"primary_key" json:"id"`
	Name      string    `grom:"column:name; not null" json:"name"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	DeletedAt time.Time `gorm:"default:null" json:"created_at"`
}
