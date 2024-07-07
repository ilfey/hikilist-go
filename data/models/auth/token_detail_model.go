package authModels

import (
	"github.com/ilfey/hikilist-go/data/entities"
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
	"github.com/ilfey/hikilist-go/internal/utils/resx"

	userModels "github.com/ilfey/hikilist-go/data/models/user"
)

// Модель токена
type TokenDetailModel struct {
	baseModels.DetailModel

	ID uint `json:"id"`

	Token string `json:"-"`

	User *userModels.UserDetailModel `json:"user"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

// Собрать модель `TokenDetailModel` из `entities.Token`
func TokenDetailModelFromEntity(entity *entities.Token) *TokenDetailModel {
	return &TokenDetailModel{
		ID: entity.ID,

		Token: entity.Token,
		User:  userModels.UserDetailModelFromEntity(entity.User),

		CreatedAt: entity.CreatedAt.String(),
		UpdatedAt: entity.UpdatedAt.String(),
		DeletedAt: entity.DeletedAt.Time.String(),
	}
}

// Преобразовать в json
func (m *TokenDetailModel) JSON() []byte {
	return m.DetailModel.JSON(m)
}

// Преобразовать в ответ
func (m *TokenDetailModel) Response() *resx.Response {
	return m.DetailModel.Response(m)
}
