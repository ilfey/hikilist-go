package userActionModels

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/rotisserie/eris"
)

type CreateModel struct {
	ID uint `json:"-"`

	UserID uint `json:"-"`

	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"-"`
}

func (CreateModel) TableName() string {
	return "user_actions"
}

func (cm *CreateModel) Insert(ctx context.Context) error {
	sql, args, err := cm.insertSQL()
	if err != nil {
		return eris.Wrap(err, "failed to build insert query")
	}

	err = database.Instance().QueryRow(ctx, sql, args...).Scan(&cm.ID)
	if err != nil {
		return eris.Wrap(err, "failed to insert user action")
	}

	return nil
}

func (cm *CreateModel) insertSQL() (string, []any, error) {
	return sq.Insert("user_actions").
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
			cm.CreatedAt,
		).
		Suffix("RETURNING id").
		ToSql()
}
