package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name; not null" json:"name"`
	Phone    string `gorm:"column:phone;not null;unique" json:"phone"`
	Password string `gorm:"column:password;not null" json:"password"`
}
