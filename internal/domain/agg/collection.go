package agg

import "time"

/* ===== Detail ===== */

type CollectionDetail struct {
	ID uint64 `json:"id"`

	UserID uint64 `json:"user_id"`

	// User *userModels.CollectionListItem `json:"user"`

	Title string `json:"title"`

	Description *string `json:"description"`

	IsPublic bool `json:"is_public"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/* ===== List ===== */

type CollectionList struct {
	Results []*CollectionListItem `json:"results"`

	Count *uint64 `json:"count,omitempty"`
}

/* ===== ListItem ===== */

type CollectionListItem struct {
	ID uint64 `json:"id"`

	UserID uint64 `json:"user_id"`

	Title string `json:"title"`

	Description *string `json:"description"`
}
