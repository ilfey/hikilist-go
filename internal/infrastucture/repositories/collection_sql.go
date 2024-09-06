package repositories

import (
	"github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

func (r *Collection) CreateSQL(cm *dto.CollectionCreateRequestDTO) (string, []any, error) {
	return squirrel.Insert(CollectionTN).
		Columns(
			"title",
			"user_id",
			"description",
			"is_public",
		).
		Values(
			cm.Title,
			cm.UserID,
			cm.Description,
			cm.IsPublic,
		).
		Suffix("RETURNING id").
		ToSql()
}

func (r *Collection) GetSQL(conds any) (string, []any, error) {
	if conds == nil {
		panic("conds is nil")
	}

	return squirrel.Select(
		"id",
		"title",
		"user_id",
		"description",
		"is_public",
		"created_at",
		"updated_at",
	).
		From(CollectionTN).
		Where(conds).
		Limit(1).
		ToSql()
}

func (r *Collection) FindSQL(dto *dto.CollectionListRequestDTO, conds any) (string, []any, error) {
	b := squirrel.Select(
		"id",
		"user_id",
		"title",
		"description",
	).
		From(CollectionTN)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.
		OrderBy("id DESC").
		Offset((dto.Page - 1) * dto.Limit).
		Limit(dto.Limit).
		ToSql()
}

func (r *Collection) CountSQL(conds any) (string, []any, error) {
	b := squirrel.Select("COUNT(*)").
		From(CollectionTN)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}

func (r *Collection) UpdateSQL(um *dto.CollectionUpdateRequestDTO) (string, []any, error) {
	return squirrel.Update(CollectionTN).
		SetMap(updateModelToMap(um)).
		Where(squirrel.Eq{
			"id":      um.CollectionID,
			"user_id": um.UserID,
		}).
		ToSql()
}
