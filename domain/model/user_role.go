package model

type UserRole struct {
	ID   int    `grom:"primary_key" json:"id"`
	Name string `grom:"column:name; not null" json:"name"`
}
