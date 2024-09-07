package dto

/* ===== Register ===== */

type AuthRegisterRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/* ===== Login ===== */

type AuthLoginRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/* ===== Refresh ===== */

type AuthRefreshRequestDTO struct {
	Refresh string `json:"refresh"`
}

/* ===== Logout ===== */

type AuthLogoutRequestDTO struct {
	Refresh string `json:"refresh"`
}
