package animeModels

import (
	"context"
	"fmt"

	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/orm"
)

type ListModel struct {
	Results []*ListItemModel `json:"results"`

	Count *int64 `json:"count,omitempty"`
}

func (lm *ListModel) Paginate(ctx context.Context, p *Paginate, conds any) error {
	p.Normalize()

	results, err := orm.Select(&ListItemModel{}).
		Where(conds).
		Limit(p.Limit).
		Offset(p.GetOffset(p.Page, p.Limit)).
		Order(p.Order.ToQuery()).
		Query(ctx, database.Instance())
	if err != nil {
		return err
	}

	lm.Results = results

	// TODO: count

	return nil
}

func (lm *ListModel) PaginateFromCollection(ctx context.Context, p *Paginate, userId, collectionId any) error {
	p.Normalize()

	sql := fmt.Sprintf(`
SELECT 
	a.id,
	a.title,
	a.poster,
	a.episodes,
	a.episodes_released
FROM animes_collections AS ac
JOIN animes AS a ON ac.anime_id = a.id
WHERE ac.collection_id = (
	SELECT 
		c.id 
	FROM collections AS c
	WHERE c.id = %d AND (c.is_public = TRUE OR c.user_id = %d)
	LIMIT 1
)
`, collectionId, userId)

	if p.Order.Field() != "" {
		sql += fmt.Sprintf(" ORDER BY %s", p.Order.ToQuery())
	}

	offset := p.GetOffset(p.Page, p.Limit)
	if offset > 0 {
		sql += fmt.Sprintf(" OFFSET %d", offset)
	}

	sql += fmt.Sprintf(" LIMIT %d;", p.Limit)

	results, err := orm.Select(&ListItemModel{}).
		QuerySQL(ctx, database.Instance(), sql)
	if err != nil {
		return err
	}

	lm.Results = results

	// TODO: count

	return nil
}
