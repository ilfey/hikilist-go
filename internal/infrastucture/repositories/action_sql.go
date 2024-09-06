package repositories

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"time"

	"github.com/Masterminds/squirrel"
)

func (r *Action) CreateSQL(cm *dto.ActionCreateRequestDTO) (string, []any, error) {
	return squirrel.Insert(ActionTN).
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

func (r *Action) FindWithPaginatorSQL(dto *dto.ActionListRequestDTO, conds any) (string, []any, error) {
	b := squirrel.Select(
		"id",
		"title",
		"description",
		"created_at",
	).
		From(ActionTN).
		OrderBy(dto.Order.ToQuery()).
		Offset((dto.Page - 1) * dto.Limit).
		Limit(dto.Limit)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}

func (r *Action) CountSQL(conds any) (string, []any, error) {
	b := squirrel.Select("COUNT(*)").
		From(ActionTN)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}
