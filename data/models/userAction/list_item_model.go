package userAction

import "time"

type ListItemModel struct {
	ID uint `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at"`
}
