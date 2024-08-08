package userAction

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/data/database"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/rotisserie/eris"
)

type DetailModel struct {
	ID uint

	UserID uint
	User   *userModels.ListItemModel

	Title       string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (dm *DetailModel) Get(ctx context.Context, conds map[string]any) error {
	sql, args, err := dm.GetSQL(conds)
	if err != nil {
		return err
	}

	err = pgxscan.Get(ctx, database.Instance(), dm, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to get user action")
	}

	return nil
}

func (DetailModel) GetSQL(conds map[string]any) (string, []any, error) {
	b := sq.Select(
		"id",
		"user_id",
		"title",
		"description",
		"created_at",
		"updated_at",
	).
		From("user_actions")

	if conds != nil {
		b = b.Where(conds)
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build user action select query")
	}

	return sql, args, nil
}
