package errtype

import (
	"fmt"
	"github.com/ilfey/hikilist-go/pkg/logger"
	"net/http"
)

const (
	authErrType           = "authorization"
	internalAuthErrLevel  = logger.CriticalLevel
	internalAuthErrStatus = http.StatusInternalServerError
	publicAuthErrLevel    = logger.InfoLevel
	publicAuthErrStatus   = http.StatusForbidden
)

/* ===== PasswordNotMatchError ===== */

type PasswordNotMatchError struct{ publicError }

func NewPasswordNotMatchError() *PasswordNotMatchError {
	return &PasswordNotMatchError{
		publicError{
			errored{
				ErrorDetail: "authorization failed: password not match",
				ErrorType:   authErrType,
				errorStatus: publicAuthErrStatus,
				errorLevel:  publicAuthErrLevel,
			},
		},
	}
}

/* ===== AuthFailedError ===== */

type AuthFailedError struct{ publicError }

func NewAuthFailedError(msg string) *AuthFailedError {
	baseMsg := "authorization failed"
	if msg != "" {
		msg = fmt.Sprintf("%v: %v", baseMsg, msg)
	} else {
		msg = baseMsg
	}

	return &AuthFailedError{
		publicError{
			errored{
				ErrorDetail: msg,
				ErrorType:   authErrType,
				errorLevel:  publicAuthErrLevel,
				errorStatus: publicAuthErrStatus,
			},
		},
	}
}

/* ===== TokenIsInvalidError ===== */

type TokenIsInvalidError struct{ publicError }

func NewAccessTokenIsInvalidError() *TokenIsInvalidError {
	return &TokenIsInvalidError{
		publicError{
			errored{
				ErrorDetail: "authorization failed: provided token is invalid",
				ErrorType:   authErrType,
				errorStatus: publicAuthErrStatus,
				errorLevel:  publicAuthErrLevel,
			},
		},
	}
}

/* ===== RefreshTokenWasBlockedError ===== */

type RefreshTokenWasBlockedError struct{ publicError }

func NewRefreshTokenWasBlockedError() *RefreshTokenWasBlockedError {
	return &RefreshTokenWasBlockedError{
		publicError{
			errored{
				ErrorDetail: "authorization failed: token was blocked",
				ErrorType:   authErrType,
				errorStatus: publicAuthErrStatus,
				errorLevel:  publicAuthErrLevel,
			},
		},
	}
}

///* ===== TokenAlgoWasNotMatchedError ===== */
//
//type TokenAlgoWasNotMatchedError struct{ internalError }
//
//func NewTokenAlgoWasNotMatchedInternalError(token string) *TokenAlgoWasNotMatchedError {
//	return &TokenAlgoWasNotMatchedError{
//		internalError{
//			errored{
//				ErrorDetail: fmt.Sprintf("token '%v' algo was not matched", token),
//				ErrorType:    authErrType,
//				errorStatus:  internalAuthErrStatus,
//				errorLevel:   internalAuthErrLevel,
//			},
//		},
//	}
//}

/* ===== TokenUnexpectedSigningMethodError ===== */

type TokenUnexpectedSigningMethodError struct{ internalError }

func NewTokenUnexpectedSigningMethodInternalError(token string, algo interface{}) *TokenUnexpectedSigningMethodError {
	return &TokenUnexpectedSigningMethodError{
		internalError{
			errored{
				ErrorDetail: fmt.Sprintf("unexpected signing algo '%v' for token '%v'", algo, token),
				ErrorType:   authErrType,
				errorStatus: internalAuthErrStatus,
				errorLevel:  internalAuthErrLevel,
			},
		},
	}
}

/* ===== TokenInvalidError ===== */

type TokenInvalidError struct{ internalError }

func NewTokenInvalidInternalError(token string) *TokenInvalidError {
	return &TokenInvalidError{
		internalError{
			errored{
				ErrorDetail: fmt.Sprintf("token '%v' is not valid", token),
				ErrorType:   authErrType,
				errorStatus: internalAuthErrStatus,
				errorLevel:  internalAuthErrLevel,
			},
		},
	}
}

/* ===== TokenIssuerWasNotMatchedError ===== */

type TokenIssuerWasNotMatchedError struct{ internalError }

func NewTokenIssuerWasNotMatchedInternalError(token string) *TokenIssuerWasNotMatchedError {
	return &TokenIssuerWasNotMatchedError{
		internalError{
			errored{
				ErrorDetail: fmt.Sprintf("token '%v' issuer was not matched", token),
				ErrorType:   authErrType,
				errorStatus: internalAuthErrStatus,
				errorLevel:  internalAuthErrLevel,
			},
		},
	}
}
