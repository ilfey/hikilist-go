package errtype

import (
	"fmt"
	"github.com/ilfey/hikilist-go/pkg/logger"
	"net/http"
)

const (
	repositoryType = "application"

	publicRepositoryLevel  = logger.ErrorLevel
	publicRepositoryStatus = http.StatusNotFound

	internalRepositoryLevel  = logger.CriticalLevel
	internalRepositoryStatus = http.StatusInternalServerError
)

type EntityAlreadyExistsError struct{ publicError }

func NewEntityAlreadyExistsError(entity string, by string) *EntityAlreadyExistsError {
	return &EntityAlreadyExistsError{
		publicError{
			errored{
				ErrorDetail: fmt.Sprintf("%v already exists by given '%s'", entity, by),
				ErrorType:   repositoryType,
				errorLevel:  publicRepositoryLevel,
				errorStatus: publicRepositoryStatus,
			},
		},
	}
}

/* ===== EntityNotFoundError ===== */

type EntityNotFoundError struct{ publicError }

func NewEntityNotFoundError(where, entity string, by string) *EntityNotFoundError {
	return &EntityNotFoundError{
		publicError{
			errored{
				ErrorDetail: fmt.Sprintf("%v: '%v' not found by given '%s'", where, entity, by),
				ErrorType:   repositoryType,
				errorLevel:  publicRepositoryLevel,
				errorStatus: publicRepositoryStatus,
			},
		},
	}
}

/* ===== InternalRepositoryError ===== */

type InternalRepositoryError struct{ internalError }

func NewInternalRepositoryError(msg string) *InternalRepositoryError {
	return &InternalRepositoryError{
		internalError{
			errored{
				ErrorDetail: msg,
				ErrorType:   repositoryType,
				errorLevel:  internalRepositoryLevel,
				errorStatus: internalRepositoryStatus,
			},
		},
	}
}
