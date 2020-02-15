package model

import "time"

type Event struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	StartEvent  time.Time `json:"start_event"`
	Price       uint      `json:"price"`
	UserID      uint      `json:"user_id"`
}

func (Event) TableName() string { return "events" }
