package repositories

import (
	"github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/pkg/models/token"
)

func (r *TokenImpl) CreateSQL(cm *token.CreateModel) (string, []any, error) {
	return squirrel.Insert(token.TableName).
		Columns(
			"token",
		).
		Values(
			cm.Token,
		).
		Suffix("RETURNING id").
		ToSql()
}

func (r *TokenImpl) GetSQL(conds any) (string, []any, error) {
	b := squirrel.Select(
		"id",
		"token",
		"created_at",
	).
		From(token.TableName).
		Limit(1)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}

func (r *TokenImpl) DeleteSQL(conds any) (string, []any, error) {
	if conds == nil {
		panic("conds is nil")
	}

	return squirrel.Delete(token.TableName).
		Where(conds).
		Suffix("RETURNING id").
		ToSql()
}
