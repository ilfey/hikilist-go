package token

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/rotisserie/eris"
)

type DetailModel struct {
	ID uint `json:"-"`

	Token string `json:"-"`

	CreatedAt time.Time `json:"created_at"`
}

func (dm *DetailModel) Get(ctx context.Context, conds map[string]any) error {
	sql, args, err := dm.GetSQL(conds)
	if err != nil {
		return err
	}

	err = pgxscan.Get(ctx, database.Instance(), dm, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to get token")
	}

	return nil
}

func (DetailModel) GetSQL(conds map[string]any) (string, []any, error) {
	b := sq.Select(
		"id",
		"token",
		"created_at",
	).
		From("tokens")

	if conds != nil {
		b = b.Where(conds)
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build token select query")
	}

	return sql, args, nil
}

func (dm *DetailModel) Delete(ctx context.Context) error {
	sql, args, err := dm.DeleteSQL()
	if err != nil {
		return err
	}

	err = database.Instance().QueryRow(ctx, sql, args...).Scan(&dm.ID)
	if err != nil {
		return eris.Wrap(err, "failed to delete token")
	}

	return nil
}

func (dm *DetailModel) DeleteSQL() (string, []any, error) {
	sql, args, err := sq.Delete("tokens").
		Where(sq.Eq{"id": dm.ID}).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build token delete query")
	}

	return sql, args, nil
}
