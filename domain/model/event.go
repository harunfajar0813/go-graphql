package model


type Event struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Address     string `gorm:"column:address" json:"address"`
	StartEvent  string `gorm:"column:start_at" json:"start_event"`
	Price       int    `gorm:"column:price" json:"price"`
	UserID      int    `gorm:"column:user_id" json:"user_id"`
}

func (Event) TableName() string { return "events" }
