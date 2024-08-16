package repositories

import (
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/pkg/models/action"
)

func (r *ActionImpl) CreateSQL(cm *action.CreateModel) (string, []any, error) {
	return squirrel.Insert("user_actions").
		Columns(
			"user_id",
			"title",
			"description",
			"created_at",
		).
		Values(
			cm.UserID,
			cm.Title,
			cm.Description,
			time.Now(),
		).
		Suffix("RETURNING id").
		ToSql()
}

func (r *ActionImpl) FindWithPaginatorSQL(p *paginate.Paginator, conds any) (string, []any, error) {
	b := squirrel.Select(
		"id",
		"title",
		"description",
		"created_at",
	).
		From("user_actions").
		OrderBy(p.Order.ToQuery()).
		Offset(uint64(p.GetOffset(p.Page, p.Limit))).
		Limit(uint64(p.Limit))

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}

func (r *ActionImpl) CountSQL(conds any) (string, []any, error) {
	b := squirrel.Select("COUNT(*)").
		From("user_actions")

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}
