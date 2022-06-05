package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"memo-RPC/userserver/conf"
)

type User struct {
	gorm.Model

	Username string `json:"username" gorm:"column:username;size:255"`
	Password string `json:"password" gorm:"column:password;size:255"`
}

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), conf.Cfg.App.PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err
}
