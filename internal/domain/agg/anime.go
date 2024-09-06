package agg

import "time"

/* ===== Detail ===== */

type AnimeDetail struct {
	ID uint64 `json:"id"`

	Title            string  `json:"title"`
	Description      *string `json:"description"`
	Poster           *string `json:"poster"`
	Episodes         *uint64 `json:"episodes"`
	EpisodesReleased uint64  `json:"episodes_released"`

	MalID   *uint64 `json:"mal_id"`
	ShikiID *uint64 `json:"shiki_id"`

	Related []*AnimeListItem `json:"related"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/* ===== List ===== */

type AnimeList struct {
	Results []*AnimeListItem `json:"results"`

	Count *uint64 `json:"count,omitempty"`
}

/* ===== ListItem ===== */

type AnimeListItem struct {
	ID uint64 `json:"id"`

	Title            string  `json:"title"`
	Poster           *string `json:"poster"`
	Episodes         *uint64 `json:"episodes"`
	EpisodesReleased uint64  `json:"episodes_released"`
}
