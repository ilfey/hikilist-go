package userModels

import (
	"time"

	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
)

type ListItemModel struct {
	ID uint `json:"id"`

	Username string `json:"username"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ListItemModel) TableName() string {
	return "users"
}

type ListModel = baseModels.ListModel[ListItemModel]

func NewListModel(items []*ListItemModel) *ListModel {
	return baseModels.NewListModel(items)
}
