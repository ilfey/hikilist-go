package tokenModels

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/rotisserie/eris"
)

type CreateModel struct {
	ID uint `json:"-"`

	Token string `json:"-"`
}

func (CreateModel) TableName() string {
	return "tokens"
}

func (cm *CreateModel) Insert(ctx context.Context) error {
	sql, args, err := cm.insertSQL()
	if err != nil {
		return eris.Wrap(err, "failed to build insert query")
	}

	err = database.Instance().QueryRow(ctx, sql, args...).Scan(&cm.ID)
	if err != nil {
		return eris.Wrap(err, "failed to insert token")
	}

	return nil
}

func (cm *CreateModel) insertSQL() (string, []any, error) {
	return sq.Insert("tokens").
		Columns(
			"token",
		).
		Values(
			cm.Token,
		).
		Suffix("RETURNING id").
		ToSql()
}
