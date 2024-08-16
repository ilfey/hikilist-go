package token

import (
	"time"
)

type DetailModel struct {
	ID uint `json:"-"`

	Token string `json:"-"`

	CreatedAt time.Time `json:"created_at"`
}
