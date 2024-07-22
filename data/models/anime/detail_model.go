package animeModels

import (
	"time"

	"github.com/ilfey/hikilist-go/data/entities"
)

// Модель аниме
type DetailModel struct {
	ID uint `json:"id"`

	Title            string  `json:"title"`
	Description      *string `json:"description"`
	Poster           *string `json:"poster"`
	Episodes         *uint   `json:"episodes"`
	EpisodesReleased uint    `json:"episodes_released"`

	MalID   *uint `json:"mal_id"`
	ShikiID *uint `json:"shiki_id"`

	Related []*ListItemModel `json:"related"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Собрать модель `DetailModel` из `entities.Anime`
func DetailModelFromEntity(entity *entities.Anime) *DetailModel {
	relatedList := make([]*ListItemModel, len(entity.Related))

	for i, entity := range entity.Related {
		relatedList[i] = &ListItemModel{
			ID: entity.ID,

			Title:            entity.Title,
			Poster:           entity.Poster,
			Episodes:         entity.Episodes,
			EpisodesReleased: entity.EpisodesReleased,
		}
	}

	return &DetailModel{
		ID: entity.ID,

		Title:            entity.Title,
		Description:      entity.Description,
		Poster:           entity.Poster,
		Episodes:         entity.Episodes,
		EpisodesReleased: entity.EpisodesReleased,

		MalID:   entity.MalID,
		ShikiID: entity.ShikiID,

		Related: relatedList,

		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
