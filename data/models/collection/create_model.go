package collectionModels

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/validator"
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

func (CreateModel) TableName() string {
	return "collections"
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
			"Title": {
				validator.Required(),
				validator.LenGreaterThat(3),
				validator.LenLessThat(256),
			},
		},
	)
}

func (cm *CreateModel) Insert(ctx context.Context) error {
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
