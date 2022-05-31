package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `json:"username" gorm:"column:username;size:255"`
	Password string `json:"password" gorm:"column:password;size:255"`
}
