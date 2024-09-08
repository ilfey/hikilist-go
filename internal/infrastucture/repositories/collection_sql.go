package repositories

import (
	"github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
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
		OrderBy(dto.Order.ToQuery()).
		Offset((dto.Page - 1) * dto.Limit).
		Limit(dto.Limit).
		ToSql()
}

func (r *Collection) FindUserPublicCollectionListSQL(dto *dto.UserCollectionListRequestDTO) (string, []any, error) {
	return squirrel.Select(
		"id",
		"user_id",
		"title",
		"description",
	).
		From(CollectionTN).
		Where(squirrel.Eq{
			"is_public": true,
			"user_id":   dto.UserID,
		}).
		OrderBy("id DESC").
		Offset((dto.Page - 1) * dto.Limit).
		Limit(dto.Limit).
		ToSql()
}

func (r *Collection) FindUserCollectionListSQL(dto *dto.UserCollectionListRequestDTO) (string, []any, error) {
	return squirrel.Select(
		"id",
		"user_id",
		"title",
		"description",
	).
		From(CollectionTN).
		Where(squirrel.Eq{
			"user_id": dto.UserID,
		}).
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

func (r *Collection) CountUserPublicCollectionSQL(req *dto.UserCollectionListRequestDTO) (string, []any, error) {
	return squirrel.Select("COUNT(*)").
		From(CollectionTN).
		Where(squirrel.Eq{
			"is_public": true,
			"user_id":   req.UserID,
		}).
		ToSql()
}

func (r *Collection) CountUserCollectionSQL(req *dto.UserCollectionListRequestDTO) (string, []any, error) {
	return squirrel.Select("COUNT(*)").
		From(CollectionTN).
		Where(squirrel.Eq{
			"user_id": req.UserID,
		}).
		ToSql()
}

func (r *Collection) UpdateSQL(req *agg.CollectionDetail) (string, []any, error) {
	return squirrel.Update(CollectionTN).
		Set("title", req.Title).
		Set("description", req.Description).
		Set("is_public", req.IsPublic).
		Set("updated_at", req.UpdatedAt).
		Where(squirrel.Eq{
			"id":      req.ID,
			"user_id": req.UserID,
		}).
		ToSql()
}

func (r *Collection) DeleteSQL(req *dto.CollectionDeleteRequestDTO) (string, []any, error) {
	return squirrel.Delete(CollectionTN).
		Where(squirrel.Eq{
			"id":      req.CollectionID,
			"user_id": req.UserID,
		}).
		Suffix("RETURNING id").
		ToSql()
}
