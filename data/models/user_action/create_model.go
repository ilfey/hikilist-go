package userActionModels

import "github.com/ilfey/hikilist-go/data/entities"

type CreateModel struct {
	UserID uint

	Title       string
	Description string
}

func (m *CreateModel) ToEntity() *entities.UserAction {
	return &entities.UserAction{
		UserID:      m.UserID,
		Title:       m.Title,
		Description: m.Description,
	}
}
