package dto

/* ===== Collection ===== */

type UserCollectionsRequestDTO struct {
	UserID uint64 `json:"-"`
}

/* ===== Create ===== */

type UserCreateRequestDTO struct {
	UserID uint64 `json:"-"`

	Username string `json:"username"`
	Password string `json:"password"`
}

/* ===== Detail ===== */

type UserDetailRequestDTO struct {
	UserID uint64
}

/* ===== Delete ===== */

type UserDeleteRequestDTO struct {
	UserID uint64 `json:"-"`

	Refresh  string `json:"refresh"`
	Password string `json:"password"`
}

/* ===== List ===== */

type UserListRequestDTO struct {
	*PaginationRequestDTO `json:"-"`
}

/* ===== Me ===== */

type UserMeRequestDTO struct {
	UserID uint64 `json:"-"`
}
