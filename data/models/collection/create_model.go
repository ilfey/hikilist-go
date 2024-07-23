package collectionModels

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/data/entities"
	"github.com/ilfey/hikilist-go/internal/validator"
)

type CreateModel struct {
	UserID      uint    `json:"-"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	IsPublic    *bool   `json:"is_public"`
}

func NewCreateModelFromRequest(request *http.Request) *CreateModel {
	model := new(CreateModel)

	json.NewDecoder(request.Body).Decode(model)

	return model
}

func (model CreateModel) Validate() validator.ValidateError {
	return validator.Validate(
		model,
		map[string][]validator.Option{
			"Name": {
				validator.Required(),
				validator.LenGreaterThat(3),
				validator.LenLessThat(256),
			},
		},
	)
}

func (m *CreateModel) ToEntity() *entities.Collection {
	entity := entities.Collection{
		UserID:      m.UserID,
		Name:        m.Name,
		Description: m.Description,
		IsPublic:    bool	(m.IsPublic != nil && *m.IsPublic),
	}

	return &entity
}
