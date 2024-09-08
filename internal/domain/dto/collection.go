package dto

/* ===== AddAnime ===== */

type CollectionAddAnimeRequestDTO struct {
	UserID       uint64 `json:"-"`
	CollectionID uint64 `json:"-"`

	Animes []uint64 `json:"animes"`
}

/* ===== AnimeListFromCollection ===== */

type AnimeListFromCollectionRequestDTO struct {
	UserID       uint64 `json:"-"`
	CollectionID uint64 `json:"-"`

	*PaginationRequestDTO `json:"-"`
}

/* ===== Create ===== */

type CollectionCreateRequestDTO struct {
	ID uint64 `json:"-"`

	UserID uint64 `json:"-"`

	Title       string  `json:"title"`
	Description *string `json:"description"`
	IsPublic    *bool   `json:"is_public"`
}

type CollectionDeleteRequestDTO struct {
	UserID       uint64 `json:"-"`
	CollectionID uint64 `json:"-"`
}

/* ===== Detail ===== */

type CollectionDetailRequestDTO struct {
	CollectionID uint64 `json:"-"`
	UserID       uint64 `json:"-"`
}

/* ===== List ===== */

type CollectionListRequestDTO struct {
	*PaginationRequestDTO `json:"-"`
}

/* ===== RemoveAnime ===== */

type CollectionRemoveAnimeRequestDTO struct {
	UserID       uint64 `json:"-"`
	CollectionID uint64 `json:"-"`

	Animes []uint64 `json:"animes"`
}

/* ===== Update ===== */

type CollectionUpdateRequestDTO struct {
	CollectionID uint64 `json:"-"`

	UserID uint64 `json:"-"`

	Title       *string `json:"title"`
	Description *string `json:"description"`
	IsPublic    *bool   `json:"is_public"`
}
