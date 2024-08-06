package tokenModels

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/rotisserie/eris"
)

type DetailModel struct {
	ID    uint   `json:"-"`
	Token string `json:"-"`
}

func (DetailModel) TableName() string {
	return "tokens"
}

func (dm *DetailModel) Get(ctx context.Context, conds map[string]any) error {
	sql, args, err := dm.getSQL(conds)
	if err != nil {
		return eris.Wrap(err, "failed to build select query")
	}

	err = pgxscan.Get(ctx, database.Instance(), dm, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to get token")
	}

	return nil
}

func (DetailModel) getSQL(conds map[string]any) (string, []any, error) {
	return sq.Select(
		"id",
		"token",
	).
		From("tokens").
		Where(conds).
		ToSql()
}

func (dm *DetailModel) Delete(ctx context.Context) error {
	sql, args, err := dm.deleteSQL()
	if err != nil {
		return eris.Wrap(err, "failed to build delete query")
	}

	_, err = database.Instance().Exec(ctx, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to delete token")
	}

	return nil
}

func (dm *DetailModel) deleteSQL() (string, []any, error) {
	return sq.Delete("tokens").
		Where(sq.Eq{"id": dm.ID}).
		ToSql()
}
