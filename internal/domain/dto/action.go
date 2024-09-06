package dto

import (
	"fmt"
	"github.com/ilfey/hikilist-go/internal/domain/types"
	"time"
)

/* ===== Create ===== */

type ActionCreateRequestDTO struct {
	ID uint64 `json:"-"`

	UserID uint64 `json:"-"`

	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"-"`
}

func NewRegisterUserAction(userId uint64) *ActionCreateRequestDTO {
	return &ActionCreateRequestDTO{
		UserID:      userId,
		Title:       "Регистрация аккаунта",
		Description: "Это начало вашего пути на сайте Hikilist.",
	}
}

func NewCreateCollectionAction(userId uint64, collectionTitle string) *ActionCreateRequestDTO {
	return &ActionCreateRequestDTO{
		UserID:      userId,
		Title:       "Создание коллекции",
		Description: fmt.Sprintf("Вы создали коллекцию \"%s\".", collectionTitle),
	}
}

func NewUpdateUsernameAction(userId uint64, oldUsername, newUsername string) *ActionCreateRequestDTO {
	return &ActionCreateRequestDTO{
		UserID: userId,
		Title:  "Обновление никнейма",
		Description: fmt.Sprintf(
			"%s останется в прошлом. Продолжим путь как %s.",
			oldUsername,
			newUsername,
		),
	}
}

/* ===== List ===== */

type ActionListRequestDTO struct {
	UserID uint64 `json:"user_id"`

	Page  uint64      `json:"page"`
	Limit uint64      `json:"limit"`
	Order types.Order `json:"order"`
}
