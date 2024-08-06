package userActionModels

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

func (DetailModel) TableName() string {
	return "user_actions"
}

func (dm *DetailModel) Get(ctx context.Context, conds map[string]any) error {
	sql, args, err := dm.getSQL(conds)
	if err != nil {
		return eris.Wrap(err, "failed to build select query")
	}

	err = pgxscan.Get(ctx, database.Instance(), dm, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to get user action")
	}

	return nil
}

func (DetailModel) getSQL(conds map[string]any) (string, []any, error) {
	return sq.Select(
		"id",
		"user_id",
		"title",
		"description",
		"created_at",
		"updated_at",
	).
		From("user_actions").
		Where(conds).
		ToSql()
}
