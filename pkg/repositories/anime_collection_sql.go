package repositories

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/internal/slicex"
	animecollection "github.com/ilfey/hikilist-go/pkg/models/anime_collection"
	"github.com/ilfey/hikilist-go/pkg/models/collection"
)

func (r *AnimeCollectionImpl) GetCollectionIdSQL(userId, collectionId uint) (string, []any, error) {
	return squirrel.Select(
		"id",
	).
		From(collection.TableName).
		Where(squirrel.Eq{
			"id":      collectionId,
			"user_id": userId,
		}).
		ToSql()
}

func (r *AnimeCollectionImpl) AddAnimesSQL(aam *animecollection.AddAnimesModel) (string, []any, error) {
	if len(aam.Animes) == 0 {
		panic("animes is empty")
	}

	b := squirrel.Insert(animecollection.TableName).
		Columns(
			"collection_id",
			"anime_id",
		)

	for _, animeId := range aam.Animes {
		b = b.Values(aam.CollectionID, animeId)
	}

	return b.ToSql()
}

func (r *AnimeCollectionImpl) RemoveAnimesSQL(ram *animecollection.RemoveAnimesModel) (string, []any, error) {
	if len(ram.Animes) == 0 {
		panic("animes is empty")
	}

	expression := slicex.Reduce(ram.Animes, func(r string, item uint) string {
		return r + fmt.Sprintf("%d,", item)
	}, "")

	return squirrel.Delete(animecollection.TableName).
		Where(fmt.Sprintf(
			"collection_id = %d AND anime_id IN (%s)",
			ram.CollectionID,
			strings.TrimSuffix(expression, ","),
		)).ToSql()
}
