package userModels

import (
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
)

type ListItemModel struct {
	ID uint `json:"id"`

	Username string `json:"username"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListModel = baseModels.ListModel[ListItemModel]

func NewListModel(items []*ListItemModel) *ListModel {
	return baseModels.NewListModel(items)
}
