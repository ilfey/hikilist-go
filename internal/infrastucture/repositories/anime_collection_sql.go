package repositories

import (
	"fmt"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"strings"

	"github.com/Masterminds/squirrel"
)

func (r *AnimeCollection) GetCollectionIdSQL(userId, collectionId uint64) (string, []any, error) {
	return squirrel.Select(
		"id",
	).
		From(CollectionTN).
		Where(squirrel.Eq{
			"id":      collectionId,
			"user_id": userId,
		}).
		ToSql()
}

func (r *AnimeCollection) AddAnimesSQL(aam *dto.CollectionAddAnimeRequestDTO) (string, []any, error) {
	if len(aam.Animes) == 0 {
		panic("animes is empty")
	}

	b := squirrel.Insert(AnimeCollectionTN).
		Columns(
			"collection_id",
			"anime_id",
		)

	for _, animeId := range aam.Animes {
		b = b.Values(aam.CollectionID, animeId)
	}

	return b.ToSql()
}

func (r *AnimeCollection) RemoveAnimesSQL(ram *dto.CollectionRemoveAnimeRequestDTO) (string, []any, error) {
	if len(ram.Animes) == 0 {
		panic("animes is empty")
	}

	var expression string

	for _, animeId := range ram.Animes {
		expression += fmt.Sprintf("%d,", animeId)
	}

	return squirrel.Delete(AnimeCollectionTN).
		Where(fmt.Sprintf(
			"collection_id = %d AND anime_id IN (%s)",
			ram.CollectionID,
			strings.TrimSuffix(expression, ","),
		)).ToSql()
}
