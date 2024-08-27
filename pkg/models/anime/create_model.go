package anime

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/sirupsen/logrus"
)

type CreateModel struct {
	ID uint `json:"-"`

	Title            string  `json:"title"`
	Description      *string `json:"description"`
	Poster           *string `json:"poster"`
	Episodes         *uint   `json:"episodes"`
	EpisodesReleased uint    `json:"episodes_released"`

	MalID   *uint `json:"mal_id"`
	ShikiID *uint `json:"shiki_id"`
}

func (cm CreateModel) Validate() error {
	return validator.Validate(
		cm,
		map[string][]options.Option{
			"Title": {
				options.LenGreaterThan(3),
				options.LenLessThan(256),
			},
			"Description": {
				options.IfNotNil(
					options.LenLessThan(4096),
				),
			},
			"Poster": {
				options.IfNotNil(
					options.LenLessThan(256),
				),
			},
			"MalID": {
				options.IfNotNil(
					options.GreaterThan[uint64](0),
				),
			},
			"ShikiID": {
				options.IfNotNil(
					options.GreaterThan[uint64](0),
				),
			},
		},
	)
}

func CreateModelFromRequest(request *http.Request) *CreateModel {
	model := new(CreateModel)

	err := json.NewDecoder(request.Body).Decode(model)
	if err != nil {
		logrus.Infof("Error occurred while decoding CreateModel %v", err)
	}

	return model
}
