package models

import "gorm.io/gorm"

// логин и пароль пользователя
type User struct {
	gorm.Model
	Login string `gorm:"primarykey"`
	Role  string
}
