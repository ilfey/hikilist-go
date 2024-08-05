package userModels

import (
	"context"

	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/orm"
)

type ListModel struct {
	Results []*ListItemModel `json:"results"`

	Count *int64 `json:"count,omitempty"`
}

func (lm *ListModel) Paginate(ctx context.Context, p *Paginate) error {
	p.Normalize()

	results, err := orm.Select(&ListItemModel{}).
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
