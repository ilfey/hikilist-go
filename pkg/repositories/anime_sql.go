package repositories

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/pkg/models/anime"
)

func (r *AnimeImpl) CreateSQL(cm *anime.CreateModel) (string, []any, error) {
	return squirrel.Insert(anime.TableName).
		Columns(
			"title",
			"description",
			"poster",
			"episodes",
			"episodes_released",
			"mal_id",
			"shiki_id",
		).
		Values(
			cm.Title,
			cm.Description,
			cm.Poster,
			cm.Episodes,
			cm.EpisodesReleased,
			cm.MalID,
			cm.ShikiID,
		).
		Suffix("RETURNING id").
		ToSql()
}

func (r *AnimeImpl) GetSQL(conds any) (string, []any, error) {
	if conds == nil {
		panic("conds is nil")
	}

	return squirrel.Select(
		"id",
		"title",
		"description",
		"poster",
		"episodes",
		"episodes_released",
		"mal_id",
		"shiki_id",
		"created_at",
		"updated_at",
	).
		From(anime.TableName).
		Where(conds).
		Limit(1).
		ToSql()
}

func (r *AnimeImpl) FindWithPaginatorSQL(p *paginate.Paginator, conds any) (string, []any, error) {
	b := squirrel.Select(
		"id",
		"title",
		"poster",
		"episodes",
		"episodes_released",
	).From(anime.TableName)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.
		OrderBy(p.Order.ToQuery()).
		Offset(uint64(p.GetOffset(p.Page, p.Limit))).
		Limit(uint64(p.Limit)).
		ToSql()
}

func (r *AnimeImpl) FindFromCollectionWithPaginatorSQL(p *paginate.Paginator, userId, collectionId uint) (string, []any, error) {
	sub, args, err := squirrel.Select(
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

	return squirrel.Select(
		"id",
		"title",
		"poster",
		"episodes",
		"episodes_released",
	).
		From("animes_collections").
		Join("animes ON animes.id = animes_collections.anime_id").
		Where(squirrel.Expr(fmt.Sprintf("collection_id = (%s)", sub), args...)).
		OrderBy(p.Order.ToQuery()).
		Offset(uint64(p.GetOffset(p.Page, p.Limit))).
		Limit(uint64(p.Limit)).
		ToSql()
}

func (r *AnimeImpl) CountInCollectionSQL(userId, collectionId uint) (string, []any, error) {
	return squirrel.Select("COUNT(*)").
		From("animes_collections").
		Join("collections ON collections.id = animes_collections.collection_id").
		Where(
			"collection_id = ? AND (is_public = TRUE OR user_id = ?)",
			collectionId,
			userId,
		).
		ToSql()
}

func (r *AnimeImpl) CountSQL(conds any) (string, []any, error) {
	b := squirrel.Select("COUNT(*)").
		From(anime.TableName)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}
