package model

type Event struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	StartEvent  string `json:"start_event"`
	Price       int    `json:"price"`
	UserID      int    `json:"user_id"`
}

func (Event) TableName() string { return "events" }
