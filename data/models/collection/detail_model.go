package collectionModels

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/rotisserie/eris"
)

type DetailModel struct {
	ID uint `json:"id"`

	UserID uint `json:"user_id"`

	// User *userModels.ListItemModel `json:"user"`

	Title string `json:"title"`

	Description *string `json:"description"`

	IsPublic bool `json:"is_public"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DetailModel) TableName() string {
	return "collections"
}

func (dm *DetailModel) Get(ctx context.Context, conds any) error {
	sql, args, err := dm.getSQL(conds)
	if err != nil {
		return eris.Wrap(err, "failed to build select query")
	}

	err = pgxscan.Get(ctx, database.Instance(), dm, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to get collection")
	}

	return nil
}

func (DetailModel) getSQL(conds any) (string, []any, error) {
	return sq.Select(
		"id",
		"title",
		"user_id",
		"description",
		"is_public",
		"created_at",
		"updated_at",
	).
		From("collections").
		Where(conds).
		ToSql()
}
