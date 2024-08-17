package collection

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
)

type UpdateModel struct {
	ID uint `json:"-"`

	UserID uint `json:"-"`

	Title       *string `json:"title"`
	Description *string `json:"description"`
	IsPublic    *bool   `json:"is_public"`
}

func (um UpdateModel) Validate() error {
	return validator.Validate(
		um,
		map[string][]options.Option{
			"Title": {
				options.IfNotNil(
					options.LenGreaterThan(3),
					options.LenLessThan(256),
				),
			},
			"Description": {
				options.IfNotNil(
					options.LenLessThan(4096),
				),
			},
		},
	)
}

func NewUpdateModelFromRequest(request *http.Request, userId, collectionId uint) *UpdateModel {
	model := new(UpdateModel)

	json.NewDecoder(request.Body).Decode(model)

	model.ID = collectionId
	model.UserID = userId

	return model
}
