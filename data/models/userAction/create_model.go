package userAction

import (
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/rotisserie/eris"
)

type CreateModel struct {
	ID uint `json:"-"`

	UserID uint `json:"-"`

	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"-"`
}

func (cm *CreateModel) InsertSQL() (string, []any, error) {
	sql, args, err := sq.Insert("user_actions").
		Columns(
			"user_id",
			"title",
			"description",
			"created_at",
		).
		Values(
			cm.UserID,
			cm.Title,
			cm.Description,
			time.Now(),
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build user action insert query")
	}

	return sql, args, nil
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
