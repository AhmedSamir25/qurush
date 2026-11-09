package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:100:not null" json:"name"`
	Email    string `gorm:"uniqueIndex not null" json:"email"`
	Password string `gorm:"size:255:not null" json:"password"`
}
