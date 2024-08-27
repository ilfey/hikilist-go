package models

import "time"

type AnimeDetailModel struct {
	ID uint `json:"id"`

	Title       string  `json:"title"`
	Description *string `json:"description"`

	Poster *string `json:"poster"`

	Episodes *uint `json:"episodes"`

	EpisodesReleased *uint `json:"episodes_released"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
