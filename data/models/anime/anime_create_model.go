package animeModels

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/validator"
)

// Модель создания аниме
type AnimeCreateModel struct {
	Title            string  `json:"title"`
	Description      *string `json:"description"`
	Poster           *string `json:"poster"`
	Episodes         *uint   `json:"episodes"`
	EpisodesReleased uint    `json:"episodes_released"`

	MalID   *uint `json:"mal_id"`
	ShikiID *uint `json:"shiki_id"`

	Related *[]uint `json:"related"`
}

// Собрать модель `AnimeCreateModel` из `http.Request`
func AnimeCreateModelFromRequest(request *http.Request) *AnimeCreateModel {
	model := new(AnimeCreateModel)

	json.NewDecoder(request.Body).Decode(model)

	return model
}

// Валидация модели
func (model AnimeCreateModel) Validate() ([]byte, bool) {
	e := validator.ValidateStruct(
		model,
		map[string][]validator.Option{
			"Title": {
				validator.LenGreaterThan(3),
				validator.LenLessThan(256),
			},
			// "Description": {
			// 	validator.LenLessThan(4096),
			// },
			// "Poster": {
			// 	validator.LenLessThan(256),
			// },
		},
	)

	if e.Success() {
		return nil, true
	}

	return e.JSON(), false
}
