package collection

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
)

type CreateModel struct {
	ID uint `json:"-"`

	UserID uint `json:"-"`

	Title       string  `json:"title"`
	Description *string `json:"description"`
	IsPublic    *bool   `json:"is_public"`
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
		},
	)
}

func NewCreateModelFromRequest(request *http.Request) *CreateModel {
	model := new(CreateModel)

	json.NewDecoder(request.Body).Decode(model)

	return model
}
