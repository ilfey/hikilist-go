package animecollection

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
)

type AddAnimesModel struct {
	UserID       uint `json:"-"`
	CollectionID uint `json:"-"`

	Animes []uint `json:"animes"`
}

func (aam *AddAnimesModel) Validate() error {
	return validator.Validate(
		aam,
		map[string][]options.Option{
			"Animes": {
				options.Required(),
				options.LenGreaterThan(0),
			},
		},
	)
}

func NewAddAnimesModelFromRequest(
	request *http.Request,
	userId uint,
	collectionId uint,
) *AddAnimesModel {
	model := new(AddAnimesModel)

	json.NewDecoder(request.Body).Decode(model)

	model.UserID = userId

	model.CollectionID = collectionId

	return model
}
