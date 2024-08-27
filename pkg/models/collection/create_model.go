package collection

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/sirupsen/logrus"
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

func NewCreateModelFromRequest(request *http.Request, userId uint) *CreateModel {
	model := new(CreateModel)

	err := json.NewDecoder(request.Body).Decode(model)
	if err != nil {
		logrus.Infof("Error occurred while decoding CreateModel %v", err)
	}

	model.UserID = userId

	return model
}
