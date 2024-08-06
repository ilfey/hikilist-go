package userModels

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/rotisserie/eris"
)

type CreateModel struct {
	ID uint `json:"-"`

	Username string `json:"username"`
	Password string `json:"password"`

	CreatedAt time.Time `json:"-"`
}

func (CreateModel) TableName() string {
	return "users"
}

func (cm *CreateModel) Insert(ctx context.Context) error {
	sql, args, err := cm.insertSQL()
	if err != nil {
		return eris.Wrap(err, "failed to build insert query")
	}

	err = database.Instance().QueryRow(ctx, sql, args...).Scan(&cm.ID)
	if err != nil {
		return eris.Wrap(err, "failed to insert user")
	}

	return nil
}

func (cm *CreateModel) insertSQL() (string, []any, error) {
	return sq.Insert("users").
		Columns(
			"username",
			"password",
			"created_at",
		).
		Values(
			cm.Username,
			cm.Password,
			time.Now(),
		).
		Suffix("RETURNING id").
		ToSql()
}
