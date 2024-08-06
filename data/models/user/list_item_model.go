package userModels

import (
	"time"
)

type ListItemModel struct {
	ID uint `json:"id"`

	Username string `json:"username"`

	CreatedAt time.Time `json:"created_at"`
}
