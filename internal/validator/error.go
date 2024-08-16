package validator

import (
	"encoding/json"
	"fmt"

	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/rotisserie/eris"
)

type ValidateError struct {
	errs map[string][]string
}

func (e *ValidateError) Error() string {
	return string(errorsx.Ignore(e.MarshalJSON()))
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

	return eris.As(err, &vErr)
}
