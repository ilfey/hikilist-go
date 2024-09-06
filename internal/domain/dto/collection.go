package dto

import "github.com/ilfey/hikilist-go/internal/domain/types"

/* ===== AddAnime ===== */

type CollectionAddAnimeRequestDTO struct {
	UserID       uint64 `json:"-"`
	CollectionID uint64 `json:"-"`

	Animes []uint64 `json:"animes"`
}

type AnimeListFromCollectionRequestDTO struct {
	UserID       uint64 `json:"-"`
	CollectionID uint64 `json:"-"`

	Page  uint64      `json:"page"`
	Limit uint64      `json:"limit"`
	Order types.Order `json:"order"`
}

/* ===== Create ===== */

type CollectionCreateRequestDTO struct {
	ID uint64 `json:"-"`

	UserID uint64 `json:"-"`

	Title       string  `json:"title"`
	Description *string `json:"description"`
	IsPublic    *bool   `json:"is_public"`
}

/* ===== Detail ===== */

type CollectionDetailRequestDTO struct {
	CollectionID uint64 `json:"-"`
	UserID       uint64 `json:"-"`
}

/* ===== List ===== */

type CollectionListRequestDTO struct {
	Page  uint64      `json:"page"`
	Limit uint64      `json:"limit"`
	Order types.Order `json:"order"`
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
