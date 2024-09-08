package dto

/* ===== Create ===== */

type AnimeCreateRequestDTO struct {
	ID uint64 `json:"-"`

	Title            string  `json:"title"`
	Description      *string `json:"description"`
	Poster           *string `json:"poster"`
	Episodes         *uint64 `json:"episodes"`
	EpisodesReleased uint64  `json:"episodes_released"`

	MalID   *uint64 `json:"mal_id"`
	ShikiID *uint64 `json:"shiki_id"`
}

/* ===== Detail ===== */

type AnimeDetailRequestDTO struct {
	ID uint64 `json:"id"`
}

/* ===== List ===== */

type AnimeListRequestDTO struct {
	*PaginationRequestDTO `json:"-"`
}
