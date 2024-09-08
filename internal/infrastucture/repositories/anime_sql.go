package repositories

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

func (r *Anime) CreateSQL(cm *dto.AnimeCreateRequestDTO) (string, []any, error) {
	return squirrel.Insert(AnimeTN).
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

func (r *Anime) GetSQL(conds any) (string, []any, error) {
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
		From(AnimeTN).
		Where(conds).
		Limit(1).
		ToSql()
}

func (r *Anime) GetByIDSQL(id uint64) (string, []any, error) {
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
		From(AnimeTN).
		Where(squirrel.Eq{
			"id": id,
		}).
		Limit(1).
		ToSql()
}

func (r *Anime) FindSQL(req *dto.PaginationRequestDTO) (string, []any, error) {
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
		From(AnimeTN).
		OrderBy(req.Order.ToQuery()).
		Offset((req.Page - 1) * req.Limit).
		Limit(req.Limit).
		ToSql()
}

func (r *Anime) FindWithPaginatorSQL(dto *dto.AnimeListRequestDTO, conds any) (string, []any, error) {
	b := squirrel.Select(
		"id",
		"title",
		"poster",
		"episodes",
		"episodes_released",
	).From(AnimeTN)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.
		OrderBy(dto.Order.ToQuery()).
		Offset((dto.Page - 1) * dto.Limit).
		Limit(dto.Limit).
		ToSql()
}

func (r *Anime) FindFromCollectionWithPaginatorSQL(dto *dto.AnimeListFromCollectionRequestDTO) (string, []any, error) {
	sub, args, err := squirrel.Select(
		"id",
	).
		From("collections").
		Where(
			"id = ? AND (is_public = TRUE OR user_id = ?)",
			dto.CollectionID,
			dto.UserID,
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
		OrderBy(dto.Order.ToQuery()).
		Offset((dto.Page - 1) * dto.Limit).
		Limit(dto.Limit).
		ToSql()
}

func (r *Anime) CountInCollectionSQL(dto *dto.AnimeListFromCollectionRequestDTO) (string, []any, error) {
	return squirrel.Select("COUNT(*)").
		From("animes_collections").
		Join("collections ON collections.id = animes_collections.collection_id").
		Where(
			"collection_id = ? AND (is_public = TRUE OR user_id = ?)",
			dto.CollectionID,
			dto.UserID,
		).
		ToSql()
}

func (r *Anime) CountSQL(conds any) (string, []any, error) {
	b := squirrel.Select("COUNT(*)").
		From(AnimeTN)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}
