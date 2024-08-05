package animeModels

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/orm"
	"github.com/ilfey/hikilist-go/internal/validator"
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

	// Related *[]uint `json:"related"`

	CreatedAt time.Time `json:"-"`
}

func (CreateModel) TableName() string {
	return "animes"
}

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

func (cm *CreateModel) Insert(ctx context.Context) error {
	cm.CreatedAt = time.Now()

	id, err := orm.Insert(cm).
		Ignore("ID").
		Exec(ctx, database.Instance())

	if err != nil {
		return err
	}

	cm.ID = id

	return nil
}
