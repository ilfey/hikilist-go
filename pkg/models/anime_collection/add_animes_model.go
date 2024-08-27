package animecollection

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/sirupsen/logrus"
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

	err := json.NewDecoder(request.Body).Decode(model)
	if err != nil {
		logrus.Infof("Error occurred while decoding AddAnimesModel %v", err)
	}

	model.UserID = userId

	model.CollectionID = collectionId

	return model
}
