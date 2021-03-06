package model

import "time"

type Invoice struct {
	ID        int       `grom:"primary_key" json:"id"`
	EventID   int       `grom:"column:event_id; not null" json:"event_id"`
	UserID    int       `gorm:"column:user_id" json:"user_id"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	DeletedAt time.Time `gorm:"default:null" json:"created_at"`
}
