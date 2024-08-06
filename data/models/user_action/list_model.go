package userActionModels

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

func (lm *ListModel) Fill(ctx context.Context, p *Paginate, conds map[string]any) error {
	err := p.Validate()
	if err != nil {
		return eris.Wrap(err, "failed to validate pagination")
	}
	
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

func (ListModel) fillResultsSQL(p *Paginate, conds map[string]any) (string, []any, error) {
	return sq.Select(
		"id",
		"title",
		"description",
		"created_at",
	).
		From("user_actions").
		Where(conds).
		Offset(uint64(p.GetOffset(p.Page, p.Limit))).
		Limit(uint64(p.Limit)).
		ToSql()
}

func (ListModel) fillCountSQL(conds map[string]any) (string, []any, error) {
	return sq.Select("COUNT(*)").
		From("user_actions").
		Where(conds).
		ToSql()
}
