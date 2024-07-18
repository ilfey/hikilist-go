package userActionModels

import (
	"time"

	"github.com/ilfey/hikilist-go/data/entities"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"gorm.io/gorm"
)

type DetailModel struct {
	ID uint

	UserID uint
	User   *userModels.ListItemModel

	Title       string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func DetailModelFromEntity(entity *entities.UserAction) *DetailModel {
	return &DetailModel{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func (m *DetailModel) ToEntity() *entities.UserAction {
	return &entities.UserAction{
		Model: gorm.Model{
			ID:        m.ID,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
		Title:       m.Title,
		Description: m.Description,
	}
}
