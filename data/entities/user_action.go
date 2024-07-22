package entities

import "gorm.io/gorm"

type UserAction struct {
	gorm.Model

	// ID пользователя
	UserID uint
	User   *User `gorm:"foreignKey:UserID;not null;constraint:onDelete:CASCADE"`

	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (UserAction) TableName() string {
	return "user_actions"
}
