package autoload

import (
	"graphi/domain/model"
	"time"
)

var user_roles = []model.UserRole{
	model.UserRole{
		Name:      "event organizer",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
	model.UserRole{
		Name:      "client",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
}

var users = []model.User{
	model.User{
		Name:        "FILKOM UB",
		Description: "Fakultas Ilmu Komputer Universitas Brawijaya",
		Email:       "filkom@gmail.com",
		Phone:       "0813",
		Password:    "1234567890",
		UserRoleID:  1,
	},
	model.User{
		Name:        "Harun",
		Email:       "harun@gmail.com",
		Phone:       "0813",
		Password:    "1234567890",
		UserRoleID:  2,
	},
}

var events = []model.Event{
	model.Event{
		Name:        "Music Ardhito",
		Description: "Music event",
		Address:     "malang",
		StartEvent:  "2020-10-10",
		Price:       100000,
		Stock:       10,
		UserID:      1,
	},
}