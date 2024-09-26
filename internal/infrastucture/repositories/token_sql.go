package repositories

import (
	"github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-database/agg"
)

func (r *Token) CreateSQL(cm *agg.TokenCreate) (string, []any, error) {
	return squirrel.Insert(TokenTN).
		Columns(
			"token",
		).
		Values(
			cm.Token,
		).
		Suffix("RETURNING id").
		ToSql()
}

func (r *Token) GetSQL(conds any) (string, []any, error) {
	b := squirrel.Select(
		"id",
		"token",
		"created_at",
	).
		From(TokenTN).
		Limit(1)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}

func (r *Token) HasSQL(token string) (string, []any, error) {
	return squirrel.Select(
		"id",
	).
		From(TokenTN).
		Where("token = ?", token).
		Limit(1).
		ToSql()
}

func (r *Token) DeleteSQL(conds any) (string, []any, error) {
	if conds == nil {
		panic("conds is nil")
	}

	return squirrel.Delete(TokenTN).
		Where(conds).
		Suffix("RETURNING id").
		ToSql()
}
