package animeModels

import (
	"github.com/ilfey/hikilist-go/data/entities"
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
	"github.com/ilfey/hikilist-go/internal/utils/resx"
)

// Модель аниме
type AnimeDetailModel struct {
	baseModels.DetailModel

	ID uint `json:"id"`

	Title            string  `json:"title"`
	Description      *string `json:"description"`
	Poster           *string `json:"poster"`
	Episodes         *uint   `json:"episodes"`
	EpisodesReleased uint    `json:"episodes_released"`

	MalID   *uint `json:"mal_id"`
	ShikiID *uint `json:"shiki_id"`

	Related []*AnimeListItemModel `json:"related"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Собрать модель `AnimeDetailModel` из `entities.Anime`
func AnimeDetailModelFromEntity(entity *entities.Anime) *AnimeDetailModel {
	relatedList := make([]*AnimeListItemModel, len(entity.Related))

	for i, entity := range entity.Related {
		relatedList[i] = &AnimeListItemModel{
			ID: entity.ID,

			Title:            entity.Title,
			Poster:           entity.Poster,
			Episodes:         entity.Episodes,
			EpisodesReleased: entity.EpisodesReleased,
		}
	}

	return &AnimeDetailModel{
		ID: entity.ID,

		Title:            entity.Title,
		Description:      entity.Description,
		Poster:           entity.Poster,
		Episodes:         entity.Episodes,
		EpisodesReleased: entity.EpisodesReleased,

		MalID:   entity.MalID,
		ShikiID: entity.ShikiID,

		Related: relatedList,

		CreatedAt: entity.CreatedAt.String(),
		UpdatedAt: entity.UpdatedAt.String(),
	}
}

// Преобразовать в json
func (m *AnimeDetailModel) JSON() []byte {
	return m.DetailModel.JSON(m)
}

// Преобразовать в ответ
func (m *AnimeDetailModel) Response() *resx.Response {
	return m.DetailModel.Response(m)
}
