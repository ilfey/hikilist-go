package token

import (
	"time"
)

type CreateModel struct {
	ID uint `json:"-"`

	Token     string    `json:"-"`
	CreatedAt time.Time `json:"-"`
}
