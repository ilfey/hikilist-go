package animecollection

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/sirupsen/logrus"
)

type RemoveAnimesModel struct {
	UserID       uint `json:"-"`
	CollectionID uint `json:"-"`

	Animes []uint `json:"animes"`
}

func (aam *RemoveAnimesModel) Validate() error {
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

func NewRemoveAnimesModelFromRequest(
	request *http.Request,
	userId uint,
	collectionId uint,
) *RemoveAnimesModel {
	model := new(RemoveAnimesModel)

	err := json.NewDecoder(request.Body).Decode(model)
	if err != nil {
		logrus.Infof("Error occurred while decoding RemoveAnimesModel %v", err)
	}

	model.UserID = userId

	model.CollectionID = collectionId

	return model
}
