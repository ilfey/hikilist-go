package userModels

import (
	"time"

	"github.com/ilfey/hikilist-go/data/entities"
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
	"gorm.io/gorm"
)

// Модель пользователя
type DetailModel struct {
	baseModels.DetailModel

	ID uint `json:"id"`

	Username string `json:"username"`
	Password string `json:"-"`

	LastOnline *time.Time `json:"last_online"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Собрать модель `UserDetailModel` из `entities.User`
func DetailModelFromEntity(entity *entities.User) *DetailModel {
	return &DetailModel{
		ID: entity.ID,

		Username: entity.Username,
		Password: entity.Password,

		LastOnline: entity.LastOnline,

		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

// Преобразовать в json
func (m *DetailModel) ToJSON() []byte {
	return m.DetailModel.ToJSON(m)
}

func (m *DetailModel) ToEntity() *entities.User {
	return &entities.User{
		Model: gorm.Model{
			ID: m.ID,

			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},

		Username: m.Username,
		Password: m.Password,

		LastOnline: m.LastOnline,
	}
}
