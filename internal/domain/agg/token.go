package agg

import (
	"time"
)

/* ===== Create ===== */

type TokenCreate struct {
	ID uint64 `json:"-"`

	Token     string    `json:"-"`
	CreatedAt time.Time `json:"-"`
}

/* ===== Detail ===== */

type TokenDetail struct {
	ID uint64 `json:"-"`

	Token string `json:"-"`

	CreatedAt time.Time `json:"created_at"`
}

/* ===== Pair ===== */

type TokenPair struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
