package collection

import (
	"time"
)

type DetailModel struct {
	ID uint `json:"id"`

	UserID uint `json:"user_id"`

	// User *userModels.ListItemModel `json:"user"`

	Title string `json:"title"`

	Description *string `json:"description"`

	IsPublic bool `json:"is_public"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
