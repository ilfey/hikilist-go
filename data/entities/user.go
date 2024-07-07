package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Tokens   []*Token
	Username string `gorm:"unique;not null"`
	Password string
}

func (User) TableName() string {
	return "users"
}
