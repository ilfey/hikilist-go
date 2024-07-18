package animeModels

import (
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
)

// Модель элемента списка аниме
type ListItemModel struct {
	ID uint `json:"id"`

	Title            string  `json:"title"`
	Poster           *string `json:"poster"`
	Episodes         *uint   `json:"episodes"`
	EpisodesReleased uint    `json:"episodes_released"`
}

// Модель списка аниме
type ListModel = baseModels.ListModel[ListItemModel]

func NewListModel(items []*ListItemModel) *ListModel {
	return baseModels.NewListModel(items)
}
