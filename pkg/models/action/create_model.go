package action

import (
	"fmt"
	"time"
)

type CreateModel struct {
	ID uint `json:"-"`

	UserID uint `json:"-"`

	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"-"`
}

func NewRegisterUserAction(userId uint) *CreateModel {
	return &CreateModel{
		UserID:      userId,
		Title:       "Регистрация аккаунта",
		Description: "Это начало вашего пути на сайте Hikilist.",
	}
}

func NewCreateCollectionAction(userId uint, collectionTitle string) *CreateModel {
	return &CreateModel{
		UserID:      userId,
		Title:       "Создание коллекции",
		Description: fmt.Sprintf("Вы создали коллекцию \"%s\".", collectionTitle),
	}
}

func NewUpdateUsernameAction(userId uint, oldUsername, newUsername string) *CreateModel {
	return &CreateModel{
		UserID: userId,
		Title:  "Обновление никнейма",
		Description: fmt.Sprintf(
			"%s останется в прошлом. Продолжим путь как %s.",
			oldUsername,
			newUsername,
		),
	}
}
