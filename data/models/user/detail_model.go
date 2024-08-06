package userModels

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/rotisserie/eris"
)

// var _ baseModels.DetailModel[DetailModel] = &DetailModel{}

// Модель пользователя
type DetailModel struct {
	ID uint `json:"id"`

	Username string `json:"username"`
	Password string `json:"-"`

	LastOnline *time.Time `json:"last_online"`

	CreatedAt time.Time `json:"created_at"`
}

func (m *DetailModel) TableName() string {
	return "users"
}

func (dm *DetailModel) Get(ctx context.Context, conds map[string]any) error {
	sql, args, err := dm.getSQL(conds)
	if err != nil {
		return eris.Wrap(err, "failed to build select query")
	}

	err = pgxscan.Get(ctx, database.Instance(), dm, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to get user")
	}
	return nil
}

func (DetailModel) getSQL(conds map[string]any) (string, []any, error) {
	return sq.Select(
		"id",
		"username",
		"password",
		"last_online",
		"created_at",
	).
		From("users").
		Where(conds).
		ToSql()
}

func (m *DetailModel) Update(ctx context.Context) error {
	sql, args, err := m.updateSQL()
	if err != nil {
		return eris.Wrap(err, "failed to build update query")
	}

	_, err = database.Instance().Exec(ctx, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to update user")
	}

	return nil
}

func (m *DetailModel) updateSQL() (string, []any, error) {
	return sq.Update("users").
		SetMap(map[string]any{
			"username":    m.Username,
			"password":    m.Password,
			"last_online": m.LastOnline,
		}).
		Where(sq.Eq{"id": m.ID}).
		ToSql()
}
