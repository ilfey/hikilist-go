package errtype

import "github.com/ilfey/hikilist-go/pkg/logger"

// errored is a base error struct, implements sketch of necessary fields and base functionality.
type errored struct {
	ErrorDetail string       `json:"detail"` // Error is error interface implementation.
	ErrorType   string       `json:"type"`   // Type of part of application: validation for example.
	errorStatus int          // Status can be represented as http or stdout exit status and so on.
	errorLevel  logger.Level // Level of an error, represents as logger.IOTA, see logger.errors file.
}

func (e errored) Error() string {
	return e.ErrorDetail
}
func (e errored) Status() int {
	return e.errorStatus
}
func (e errored) Level() logger.Level {
	return e.errorLevel
}
func (e errored) Type() string {
	return e.ErrorType
}

/* ===== PublicError ===== */

type publicError struct{ errored }

func (e publicError) Public() bool {
	return true
}

/* ===== InternalError ===== */

type internalError struct{ errored }

func (e internalError) Internal() bool {
	return true
}
