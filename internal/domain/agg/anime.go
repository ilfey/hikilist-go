package agg

import "github.com/ilfey/hikilist-go/internal/domain/entity"

/* ===== Detail ===== */

type Anime struct {
	entity.Anime
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
