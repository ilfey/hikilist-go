package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	// Имя пользователя
	Username string `gorm:"unique;not null"`
	// Хешированный пароль
	Password string
	// Последняя активность пользователя
	LastOnline *time.Time

	Collections []*Collection
}

func (User) TableName() string {
	return "users"
}
