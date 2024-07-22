package collectionModels

import (
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
)

type ListItemModel struct {
	ID uint `json:"id"`

	UserID uint `json:"user_id"`

	User *userModels.ListItemModel `json:"user" gorm:""`

	Name string `json:"name"`

	Description *string `json:"description"`

	IsPublic bool `json:"is_public"`
}

type ListModel = baseModels.ListModel[ListItemModel]

func NewListModel(items []*ListItemModel) *ListModel {
	return baseModels.NewListModel(items)
}
