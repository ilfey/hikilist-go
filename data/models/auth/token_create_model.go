package authModels

import (
	userModels "github.com/ilfey/hikilist-go/data/models/user"
)

// Модель создания токена
type TokenCreateModel struct {
	Token string                     `json:"refresh"`
	User  userModels.UserDetailModel `json:"user"`
}
