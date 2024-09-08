package entity

import "time"

type Anime struct {
	ID uint64 `json:"id"`

	Title            string  `json:"title"`
	Description      *string `json:"description,omitempty"`
	Poster           *string `json:"poster,omitempty"`
	Episodes         *uint64 `json:"episodes,omitempty"`
	EpisodesReleased uint64  `json:"episodes_released"`

	MalID   *uint64 `json:"mal_id,omitempty"`
	ShikiID *uint64 `json:"shiki_id,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
