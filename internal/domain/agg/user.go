package agg

import "time"

/* ===== Detail ===== */

type UserDetail struct {
	ID uint64 `json:"id"`

	Username string `json:"username"`
	Password string `json:"-"`

	LastOnline *time.Time `json:"last_online"`

	CreatedAt time.Time `json:"created_at"`
}

func (model *UserDetail) GetHash() string {
	return model.Password
}

/* ===== List ===== */

type UserList struct {
	Results []*UserListItem `json:"results"` // results
	Count   *uint64         `json:"count"`
}

/* ===== ListItem ===== */

type UserListItem struct {
	ID uint64 `json:"id"`

	Username string `json:"username"`

	CreatedAt time.Time `json:"created_at"`
}
