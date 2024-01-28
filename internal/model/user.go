package model

import (
	"gorm.io/gorm"
)

// User represents the user model.
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
}
