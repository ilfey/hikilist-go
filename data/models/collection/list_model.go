package collectionModels

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/rotisserie/eris"
)

type ListModel struct {
	Results []*ListItemModel `json:"results"`

	Count *int64 `json:"count,omitempty"`
}

func (lm *ListModel) Fill(ctx context.Context, p *Paginate, conds any) error {
	p.Normalize()

	sql, args, err := lm.fillResultsSQL(p, conds)
	if err != nil {
		return eris.Wrap(err, "failed to build select query")
	}

	err = pgxscan.Select(ctx, database.Instance(), &lm.Results, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to execute select query")
	}

	sql, args, err = lm.fillCountSQL(conds)
	if err != nil {
		return eris.Wrap(err, "failed to build count query")
	}

	err = database.Instance().QueryRow(ctx, sql, args...).Scan(&lm.Count)
	if err != nil {
		return eris.Wrap(err, "failed to execute count query")
	}

	return nil
}

func (ListModel) fillResultsSQL(p *Paginate, conds any) (string, []any, error) {
	return sq.Select(
		"id",
		"user_id",
		"title",
		"created_at",
		"updated_at",
	).
		From("collections").
		Where(conds).
		OrderBy("id DESC").
		Offset(uint64(p.GetOffset(p.Page, p.Limit))).
		Limit(uint64(p.Limit)).
		ToSql()
}

func (ListModel) fillCountSQL(conds any) (string, []any, error) {
	return sq.Select("COUNT(*)").
		From("collections").
		Where(conds).
		ToSql()
}
