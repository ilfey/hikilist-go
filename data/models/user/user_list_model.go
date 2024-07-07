package userModels

import (
	"github.com/ilfey/hikilist-go/data/entities"
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
)

type UserListItemModel struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserListModel = baseModels.ListModel[UserListItemModel]

func UserListModelFromEntities(entities []*entities.User, count int64) *UserListModel {
	results := make([]*UserListItemModel, len(entities))

	for i, entity := range entities {
		results[i] = &UserListItemModel{
			ID: entity.ID,

			Username: entity.Username,

			CreatedAt: entity.CreatedAt.String(),
			UpdatedAt: entity.UpdatedAt.String(),
		}
	}

	return &UserListModel{
		Results: results,
		Count:   count,
	}
}
