package animeModels

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/rotisserie/eris"

	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
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
			"Poster": {
				options.IfNotNil(
					options.LenLessThan(256),
				),
			},
			"MalID": {
				options.IfNotNil(
					options.GreaterThan[int64](0),
				),
			},
			"ShikiID": {
				options.IfNotNil(
					options.GreaterThan[int64](0),
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
		return eris.Wrap(err, "failed to create insert sql")
	}

	return database.Instance().
		QueryRow(ctx, sql, args...).
		Scan(&cm.ID)
}

func (cm *CreateModel) insertSQL() (string, []any, error) {
	return sq.Insert("animes").
		Columns(
			"title",
			"description",
			"poster",
			"episodes",
			"episodes_released",
			"mal_id",
			"shiki_id",
			"created_at",
		).
		Values(
			cm.Title,
			cm.Description,
			cm.Poster,
			cm.Episodes,
			cm.EpisodesReleased,
			cm.MalID,
			cm.ShikiID,
			time.Now(),
		).
		Suffix("RETURNING id").
		ToSql()
}

func CreateModelFromRequest(request *http.Request) *CreateModel {
	model := new(CreateModel)

	json.NewDecoder(request.Body).Decode(model)

	return model
}
