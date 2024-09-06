package agg

import "time"

/* ===== List ===== */

type ActionList struct {
	Results []*ActionListItem `json:"results"`

	Count *uint64 `json:"count,omitempty"`
}

/* ===== ListItem ===== */

type ActionListItem struct {
	ID uint64 `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at"`
}
