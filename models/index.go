package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"name"`
	Phone    string `gorm:"phone"`
	Password string `gorm:"password"`
}
