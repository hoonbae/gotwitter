package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"-"`

	IsAdmin bool `json:"is_admin"`
}
