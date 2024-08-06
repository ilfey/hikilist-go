package collectionModels

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/rotisserie/eris"
)

type CreateModel struct {
	ID uint `json:"-"`

	UserID uint `json:"-"`

	Title       string  `json:"title"`
	Description *string `json:"description"`
	IsPublic    *bool   `json:"is_public"`

	CreatedAt time.Time `json:"-"`
}

func (cm CreateModel) Validate() error {
	return validator.Validate(
		cm,
		map[string][]options.Option{
			"Title": {
				options.Required(),
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

func (cm *CreateModel) Insert(ctx context.Context) error {
	err := cm.Validate()
	if err != nil {
		return eris.Wrap(err, "failed to validate model")
	}

	sql, args, err := cm.insertSQL()
	if err != nil {
		return eris.Wrap(err, "failed to build insert query")
	}

	err = database.Instance().QueryRow(ctx, sql, args...).Scan(&cm.ID)
	if err != nil {
		return eris.Wrap(err, "failed to insert collection")
	}

	return nil
}

func (cm *CreateModel) insertSQL() (string, []any, error) {
	return sq.Insert("collections").
		Columns(
			"title",
			"user_id",
			"description",
			"is_public",
			"created_at",
		).
		Values(
			cm.Title,
			cm.Description,
			cm.IsPublic,
			cm.UserID,
			time.Now(),
		).
		Suffix("RETURNING id").
		ToSql()
}
func NewCreateModelFromRequest(request *http.Request) *CreateModel {
	model := new(CreateModel)

	json.NewDecoder(request.Body).Decode(model)

	return model
}
