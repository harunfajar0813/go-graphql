package model

import "time"

type Balance struct {
	ID        int       `grom:"primary_key" json:"id"`
	Amount    int       `grom:"column:amount; not null" json:"name"`
	UserID    int       `gorm:"column:user_id" json:"user_id"`
	User      User      `json:"user"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	DeletedAt time.Time `gorm:"default:null" json:"created_at"`
}
