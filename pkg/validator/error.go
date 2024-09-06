package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ilfey/hikilist-go/pkg/errorsx"
)

type ValidateError struct {
	errs map[string][]string
}

func (e *ValidateError) Error() string {
	return string(errorsx.Ignore(e.MarshalJSON()))
}

func (e *ValidateError) GetErrors() map[string][]string {
	return e.errs
}

func (e *ValidateError) MarshalJSON() ([]byte, error) {
	_map := make(map[string][]string)
	for k, v := range e.errs {
		for _, msg := range v {
			_map[k] = append(_map[k], fmt.Sprintf(msg, k))
		}
	}

	return json.Marshal(_map)
}

func IsValidateError(err error) bool {
	var vErr *ValidateError

	return errors.As(err, &vErr)
}
