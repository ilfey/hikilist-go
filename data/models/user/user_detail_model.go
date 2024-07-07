package userModels

import (
	"github.com/ilfey/hikilist-go/data/entities"
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
	"github.com/ilfey/hikilist-go/internal/utils/resx"
)

// Модель пользователя
type UserDetailModel struct {
	baseModels.DetailModel

	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Собрать модель `UserDetailModel` из `entities.User`
func UserDetailModelFromEntity(entity *entities.User) *UserDetailModel {
	return &UserDetailModel{
		ID: entity.ID,

		Username: entity.Username,
		Password: entity.Password,

		CreatedAt: entity.CreatedAt.String(),
		UpdatedAt: entity.UpdatedAt.String(),
	}
}

// Преобразовать в json
func (m *UserDetailModel) JSON() []byte {
	return m.DetailModel.JSON(m)
}

// Преобразовать в ответ
func (m *UserDetailModel) Response() *resx.Response {
	return m.DetailModel.Response(m)
}
