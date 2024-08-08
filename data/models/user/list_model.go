package user

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

func (ListModel) FillResultsSQL(p *Paginate, conds map[string]any) (string, []any, error) {
	b := sq.Select(
		"id",
		"username",
		"created_at",
	).
		From("users")

	if conds != nil {
		b = b.Where(conds)

	}

	sql, args, err := b.
		OrderBy(p.Order.ToQuery()).
		Offset(uint64(p.GetOffset(p.Page, p.Limit))).
		Limit(uint64(p.Limit)).ToSql()
	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build users select query")
	}

	return sql, args, nil
}

func (ListModel) FillCountSQL(conds map[string]any) (string, []any, error) {
	b := sq.Select("COUNT(*)").
		From("users")
	if conds != nil {
		b = b.Where(conds)
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build users count query")
	}

	return sql, args, nil
}
