package userModels

import (
	"github.com/ilfey/hikilist-go/data/entities"
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
)

type ListItemModel struct {
	ID uint `json:"id"`

	Username string `json:"username"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListModel = baseModels.ListModel[ListItemModel]

func UserListModelFromEntities(entities []*entities.User, count int64) *ListModel {
	results := make([]*ListItemModel, len(entities))

	for i, entity := range entities {
		results[i] = &ListItemModel{
			ID: entity.ID,

			Username: entity.Username,

			CreatedAt: entity.CreatedAt.String(),
			UpdatedAt: entity.UpdatedAt.String(),
		}
	}

	return &ListModel{
		Results: results,
		Count:   count,
	}
}
