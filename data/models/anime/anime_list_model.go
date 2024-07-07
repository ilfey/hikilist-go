package animeModels

import (
	"github.com/ilfey/hikilist-go/data/entities"
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
)

// Модель элемента списка аниме
type AnimeListItemModel struct {
	ID uint `json:"id"`

	Title            string  `json:"title"`
	Poster           *string `json:"poster"`
	Episodes         *uint   `json:"episodes"`
	EpisodesReleased uint    `json:"episodes_released"`
}

// Модель списка аниме
type AnimeListModel = baseModels.ListModel[AnimeListItemModel]

// Собрать модель `AnimeListModel` из `entities.Anime`
func AnimeListModelFromEntities(entities []*entities.Anime, count int64) *AnimeListModel {
	results := make([]*AnimeListItemModel, len(entities))

	for i, entity := range entities {
		results[i] = &AnimeListItemModel{
			ID: entity.ID,

			Title:            entity.Title,
			Poster:           entity.Poster,
			Episodes:         entity.Episodes,
			EpisodesReleased: entity.EpisodesReleased,
		}
	}

	return &AnimeListModel{
		Results: results,
		Count:   count,
	}
}