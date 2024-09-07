package errtype

import (
	"fmt"
	"github.com/ilfey/hikilist-go/pkg/logger"
	"net/http"
)

const (
	validateType = "validation"

	publicValidateLevel  = logger.ErrorLevel
	publicValidateStatus = http.StatusBadRequest
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
