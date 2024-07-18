package animeModels

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/data/entities"
	"github.com/ilfey/hikilist-go/internal/validator"
)

// Модель создания аниме
type CreateModel struct {
	Title            string  `json:"title"`
	Description      *string `json:"description"`
	Poster           *string `json:"poster"`
	Episodes         *uint   `json:"episodes"`
	EpisodesReleased uint    `json:"episodes_released"`

	MalID   *uint `json:"mal_id"`
	ShikiID *uint `json:"shiki_id"`

	Related *[]uint `json:"related"`
}

// Собрать модель `CreateModel` из `http.Request`
func CreateModelFromRequest(request *http.Request) *CreateModel {
	model := new(CreateModel)

	json.NewDecoder(request.Body).Decode(model)

	return model
}

// Валидация модели
func (model CreateModel) Validate() validator.ValidateError {
	return validator.Validate(
		model,
		map[string][]validator.Option{
			"Title": {
				validator.LenGreaterThat(3),
				validator.LenLessThat(256),
			},
			// "Description": {
			// 	validator.LenLessThan(4096),
			// },
			// "Poster": {
			// 	validator.LenLessThan(256),
			// },
		},
	)
}

func (m *CreateModel) ToEntity() *entities.Anime {
	var related []*entities.Anime // Массив связанных аниме (пустые сущности с указанными id)

	if m.Related != nil { // Если в модели указан массив связанных аниме
		related = make([]*entities.Anime, len(*m.Related))

		for i, item := range *m.Related { // Заполняем массив моделями
			related[i] = &entities.Anime{}
			related[i].ID = item
		}
	}

	return &entities.Anime{
		Title:            m.Title,
		Description:      m.Description,
		Poster:           m.Poster,
		Episodes:         m.Episodes,
		EpisodesReleased: m.EpisodesReleased,

		MalID:   m.MalID,
		ShikiID: m.ShikiID,

		Related: related,
	}
}
