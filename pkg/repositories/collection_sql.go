package repositories

import (
	"github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/pkg/models/collection"
)

func (r *CollectionImpl) CreateSQL(cm *collection.CreateModel) (string, []any, error) {
	return squirrel.Insert(collection.TableName).
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

func (r *CollectionImpl) GetSQL(conds any) (string, []any, error) {
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
		From(collection.TableName).
		Where(conds).
		Limit(1).
		ToSql()
}

func (r *CollectionImpl) FindSQL(p *paginate.Paginator, conds any) (string, []any, error) {
	b := squirrel.Select(
		"id",
		"user_id",
		"title",
		"description",
	).
		From(collection.TableName)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.
		OrderBy("id DESC").
		Offset(uint64(p.GetOffset(p.Page, p.Limit))).
		Limit(uint64(p.Limit)).
		ToSql()
}

func (r *CollectionImpl) CountSQL(conds any) (string, []any, error) {
	b := squirrel.Select("COUNT(*)").
		From(collection.TableName)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}

func (r *CollectionImpl) UpdateSQL(um *collection.UpdateModel) (string, []any, error) {
	return squirrel.Update(collection.TableName).
		SetMap(updateModelToMap(um)).
		Where(squirrel.Eq{
			"id":      um.ID,
			"user_id": um.UserID,
		}).
		ToSql()
}
