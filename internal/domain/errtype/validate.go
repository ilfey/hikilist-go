package errtype

import (
	"fmt"
	"github.com/ilfey/hikilist-go/pkg/logger"
	"github.com/pkg/errors"
	"net/http"
)

const (
	validateType = "validation"

	publicValidateLevel  = logger.ErrorLevel
	publicValidateStatus = http.StatusBadRequest

	internalValidateLevel  = logger.CriticalLevel
	internalValidateStatus = http.StatusInternalServerError
)

/* ===== FieldCannotBeEmptyError ===== */

type FieldCannotBeEmptyError struct{ publicError }

func NewFieldCannotBeEmptyError(field string) *FieldCannotBeEmptyError {
	return &FieldCannotBeEmptyError{
		publicError{
			errored{
				ErrorDetail: fmt.Sprintf("field '%s' cannot be empty", field),
				ErrorType:   validateType,
				errorLevel:  publicValidateLevel,
				errorStatus: publicValidateStatus,
			},
		},
	}
}

func IsFieldCannotBeEmptyError(err error) bool {
	return errors.As(err, &FieldCannotBeEmptyError{})
}

/* ===== FieldMustBeIntegerError ===== */

type FieldMustBeIntegerError struct{ publicError }

func NewFieldMustBeIntegerError(field string) *FieldMustBeIntegerError {
	return &FieldMustBeIntegerError{
		publicError{
			errored{
				ErrorDetail: fmt.Sprintf("field '%s' must be integer", field),
				ErrorType:   validateType,
				errorLevel:  publicValidateLevel,
				errorStatus: publicValidateStatus,
			},
		},
	}
}

func IsFieldMustBeIntegerError(err error) bool {
	return errors.As(err, &FieldMustBeIntegerError{})
}

/* ===== ValidatorError ===== */

type ValidatorError struct {
	publicError

	Expectations map[string][]string `json:"expectations"`
}

func NewValidatorError(dto string, expectations map[string][]string) *ValidatorError {
	return &ValidatorError{
		publicError{
			errored{
				ErrorDetail: fmt.Sprintf("dto '%s' validation error", dto),
				ErrorType:   validateType,
				errorLevel:  publicValidateLevel,
				errorStatus: publicValidateStatus,
			},
		},
		expectations,
	}
}

func IsValidatorError(err error) bool {
	return errors.As(err, &ValidatorError{})
}

/* ===== BodyIsEmptyError ===== */

type BodyIsEmptyError struct{ publicError }

func NewBodyIsEmptyError() *BodyIsEmptyError {
	return &BodyIsEmptyError{
		publicError{
			errored{
				ErrorDetail: "body cannot be empty",
				ErrorType:   validateType,
				errorLevel:  publicValidateLevel,
				errorStatus: publicValidateStatus,
			},
		},
	}
}

/* ===== InternalValidateError ===== */

type InternalValidateError struct{ internalError }

func NewInternalValidateError(msg string) *InternalValidateError {
	return &InternalValidateError{
		internalError{
			errored{
				ErrorDetail: msg,
				ErrorType:   validateType,
				errorLevel:  internalValidateLevel,
				errorStatus: internalValidateStatus,
			},
		},
	}
}

func IsInternalValidateError(err error) bool {
	return errors.As(err, &InternalValidateError{})
}
