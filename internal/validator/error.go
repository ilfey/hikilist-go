package validator

import (
	"encoding/json"
	"fmt"

	"github.com/ilfey/hikilist-go/internal/errorsx"
)

type errorImpl struct {
	errs map[string][]string
}

type ValidateError interface {
	error

	MarshalJSON() ([]byte, error)
}

func (e *errorImpl) Error() string {
	return string(errorsx.Ignore(e.MarshalJSON()))
}

func (e *errorImpl) MarshalJSON() ([]byte, error) {
	_map := make(map[string][]string)
	for k, v := range e.errs {
		for _, msg := range v {
			_map[k] = append(_map[k], fmt.Sprintf(msg, k))
		}
	}

	return json.Marshal(_map)
}
