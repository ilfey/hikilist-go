package collectionModels

import (
	"time"

	"github.com/ilfey/hikilist-go/data/entities"
	animeModels "github.com/ilfey/hikilist-go/data/models/anime"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
)

type DetailModel struct {
	ID uint `json:"id"`

	User *userModels.ListItemModel `json:"user"`

	Name string `json:"name"`

	Description *string `json:"description"`

	IsPublic bool `json:"is_public"`

	Animes []*animeModels.ListItemModel `json:"animes"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewDetailModelFromEntity(entity *entities.Collection) *DetailModel {
	model := DetailModel{
		ID: entity.ID,

		Name: entity.Name,

		Description: entity.Description,

		IsPublic: entity.IsPublic,

		Animes: make([]*animeModels.ListItemModel, 0, len(entity.Animes)),

		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	if entity.User != nil {
		model.User = &userModels.ListItemModel{
			ID:        entity.User.ID,
			Username:  entity.User.Username,
			CreatedAt: entity.User.CreatedAt,
			UpdatedAt: entity.User.UpdatedAt,
		}
	}

	return &model
}
