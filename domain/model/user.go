package model

type User struct {
	ID        int     `gorm:"primary_key" json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Password  string  `json:"password"`
	Events    []Event `gorm:"PRELOAD:false" json:"events"`
}

func (User) TableName() string { return "users" }
