package anime

import (
	"context"
	"fmt"

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

	err = pgxscan.Get(ctx, database.Instance(), &lm.Count, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to execute count query")
	}

	return nil
}

func (ListModel) FillResultsSQL(p *Paginate, conds map[string]any) (string, []any, error) {
	b := sq.Select(
		"id",
		"title",
		"poster",
		"episodes",
		"episodes_released",
	).From("animes")

	if conds != nil {
		b = b.Where(conds)
	}

	sql, args, err := b.
		OrderBy(p.Order.ToQuery()).
		Offset(uint64(p.GetOffset(p.Page, p.Limit))).
		Limit(uint64(p.Limit)).
		ToSql()
	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build animes select query")
	}

	return sql, args, nil
}
func (ListModel) FillCountSQL(conds map[string]any) (string, []any, error) {
	b := sq.Select("COUNT(*)").
		From("animes")

	if conds != nil {
		b = b.Where(conds)
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build animes count query")
	}

	return sql, args, nil
}

func (lm *ListModel) FillFromCollection(ctx context.Context, p *Paginate, userId, collectionId uint) error {
	err := p.Validate()
	if err != nil {
		return eris.Wrap(err, "failed to validate pagination")
	}

	p.Normalize()

	sql, args, err := lm.FillFromCollectionResultsSQL(p, userId, collectionId)
	if err != nil {
		return err
	}

	err = pgxscan.Select(ctx, database.Instance(), &lm.Results, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to execute select query")
	}

	sql, args, err = lm.FillFromCollectionCountSQL(userId, collectionId)
	if err != nil {
		return err
	}

	err = pgxscan.Get(ctx, database.Instance(), &lm.Count, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to execute count query")
	}

	return nil
}

func (ListModel) FillFromCollectionResultsSQL(p *Paginate, userId, collectionId uint) (string, []any, error) {
	sub, args, err := sq.Select(
		"id",
	).
		From("collections").
		Where(
			"id = ? AND (is_public = TRUE OR user_id = ?)",
			collectionId,
			userId,
		).
		ToSql()

	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build collections select subquery")
	}

	sql, args, err := sq.Select(
		"id",
		"title",
		"poster",
		"episodes",
		"episodes_released",
	).
		From("animes_collections").
		Join("animes ON animes.id = animes_collections.anime_id").
		Where(sq.Expr(fmt.Sprintf("collection_id = (%s)", sub), args...)).
		OrderBy(p.Order.ToQuery()).
		Offset(uint64(p.GetOffset(p.Page, p.Limit))).
		Limit(uint64(p.Limit)).
		ToSql()

	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build anime select query")
	}

	return sql, args, nil
}

func (ListModel) FillFromCollectionCountSQL(userId, collectionId uint) (string, []any, error) {
	sql, args, err := sq.Select("COUNT(*)").
		From("animes_collections").
		Join("collections ON collections.id = animes_collections.collection_id").
		Where(
			"collection_id = ? AND (is_public = TRUE OR user_id = ?)",
			collectionId,
			userId,
		).
		ToSql()
	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build collections count query")
	}

	return sql, args, nil
}
