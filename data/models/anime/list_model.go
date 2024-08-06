package animeModels

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

	err = pgxscan.Get(ctx, database.Instance(), &lm.Count, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to execute count query")
	}

	return nil
}

func (ListModel) fillResultsSQL(p *Paginate, conds map[string]any) (string, []any, error) {
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

	return b.
		OrderBy(p.Order.ToQuery()).
		Offset(uint64(p.GetOffset(p.Page, p.Limit))).
		Limit(uint64(p.Limit)).
		ToSql()
}
func (ListModel) fillCountSQL(conds map[string]any) (string, []any, error) {
	b := sq.Select("COUNT(*)").
		From("animes")

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}

func (lm *ListModel) FillFromCollection(ctx context.Context, p *Paginate, userId, collectionId uint) error {
	err := p.Validate()
	if err != nil {
		return eris.Wrap(err, "failed to validate pagination")
	}

	p.Normalize()

	sql, args, err := lm.fillFromCollectionResultsSQL(userId, collectionId)
	if err != nil {
		return eris.Wrap(err, "failed to build select query")
	}

	err = pgxscan.Select(ctx, database.Instance(), &lm.Results, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to execute select query")
	}

	sql, args, err = lm.fillFromCollectionCountSQL(userId, collectionId)
	if err != nil {
		return eris.Wrap(err, "failed to build count query")
	}

	err = pgxscan.Get(ctx, database.Instance(), &lm.Count, sql, args...)
	if err != nil {
		return eris.Wrap(err, "failed to execute count query")
	}

	return nil
}

func (ListModel) fillFromCollectionResultsSQL(userId, collectionId uint) (string, []any, error) {
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
		return "", nil, err
	}

	return sq.Select(
		"id",
		"title",
		"poster",
		"episodes",
		"episodes_released",
	).
		From("animes_collections").
		Join("animes ON animes.id = animes_collections.anime_id").
		Where(sq.Expr(sub, args)).
		ToSql()
}

func (ListModel) fillFromCollectionCountSQL(userId, collectionId uint) (string, []any, error) {
	return sq.Select("COUNT(*)").
		From("animes_collections").
		Where(
			"id = ? AND (is_public = TRUE OR user_id = ?)",
			collectionId,
			userId,
		).
		ToSql()
}

// func (lm *ListModel) PaginateFromCollection(ctx context.Context, p *Paginate, userId, collectionId any) error {
// 	p.Normalize()

// 	sql := fmt.Sprintf(`
// SELECT
// 	a.id,
// 	a.title,
// 	a.poster,
// 	a.episodes,
// 	a.episodes_released
// FROM animes_collections AS ac
// JOIN animes AS a ON ac.anime_id = a.id
// WHERE ac.collection_id = (
// 	SELECT
// 		c.id
// 	FROM collections AS c
// 	WHERE c.id = %d AND (c.is_public = TRUE OR c.user_id = %d)
// 	LIMIT 1
// )
// `, collectionId, userId)

// 	if p.Order.Field() != "" {
// 		sql += fmt.Sprintf(" ORDER BY %s", p.Order.ToQuery())
// 	}

// 	offset := p.GetOffset(p.Page, p.Limit)
// 	if offset > 0 {
// 		sql += fmt.Sprintf(" OFFSET %d", offset)
// 	}

// 	sql += fmt.Sprintf(" LIMIT %d;", p.Limit)

// 	results, err := orm.Select(&ListItemModel{}).
// 		QuerySQL(ctx, database.Instance(), sql)
// 	if err != nil {
// 		return err
// 	}

// 	lm.Results = results

// 	// TODO: count

// 	return nil
// }
