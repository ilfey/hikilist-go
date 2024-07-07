package entities

import "gorm.io/gorm"

type Token struct {
	gorm.Model

	Token  string `gorm:"not null"`
	User   *User
	UserID uint
}

func (Token) TableName() string {
	return "outstanding_tokens"
}
