package collection

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
	err := p.Validate()
	if err != nil {
		return eris.Wrap(err, "failed to validate pagination")
	}

	p.Normalize()

	sql, args, err := lm.FillResultsSQL(p, conds)
	if err != nil {
		return err
	}

	err = pgxscan.Select(ctx, database.Instance(), &lm.Results, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to execute select query")
	}

	sql, args, err = lm.FillCountSQL(conds)
	if err != nil {
		return err
	}

	err = database.Instance().QueryRow(ctx, sql, args...).Scan(&lm.Count)
	if err != nil {
		return eris.Wrap(err, "failed to execute count query")
	}

	return nil
}

func (ListModel) FillResultsSQL(p *Paginate, conds any) (string, []any, error) {
	b := sq.Select(
		"id",
		"user_id",
		"title",
		"created_at",
		"updated_at",
	).
		From("collections")

	if conds != nil {
		b = b.Where(conds)
	}

	sql, args, err := b.
		OrderBy("id DESC").
		Offset(uint64(p.GetOffset(p.Page, p.Limit))).
		Limit(uint64(p.Limit)).
		ToSql()
	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build collections select query")
	}

	return sql, args, nil
}

func (ListModel) FillCountSQL(conds any) (string, []any, error) {
	b := sq.Select("COUNT(*)").
		From("collections")

	if conds != nil {
		b = b.Where(conds)
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build collections count query")
	}

	return sql, args, err
}
