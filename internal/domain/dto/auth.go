package dto

/* ===== ChangePassword ===== */

// TODO: Remove request DTO
type AuthChangeUsernameRequestDTO struct {
	NewUsername string `json:"new_username"`
	Password    string `json:"password"`
}

/* ===== ChangePassword ===== */

// TODO: Remove request DTO
type AuthChangePasswordRequestDTO struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

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
