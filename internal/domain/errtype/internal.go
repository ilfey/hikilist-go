package errtype

import (
	"github.com/ilfey/hikilist-go/pkg/logger"
	"net/http"
)

const (
	applicationType               = "application"
	internalServerErrorMessage    = "internal server error"
	internalServerErrorStatusCode = http.StatusInternalServerError
	internalServerErrorLevel      = logger.ErrorLevel
)

type InternalServerError struct{ publicError }

func NewInternalServerError() InternalServerError {
	return InternalServerError{
		publicError{
			errored{
				ErrorType:   applicationType,
				ErrorDetail: internalServerErrorMessage,
				errorStatus: internalServerErrorStatusCode,
				errorLevel:  internalServerErrorLevel,
			},
		},
	}
}
