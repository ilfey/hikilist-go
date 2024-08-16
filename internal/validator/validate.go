package validator

import (
	"reflect"
	"strings"

	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/rotisserie/eris"
)

func Validate(st any, fieldsOpts map[string][]options.Option) error {
	errs := map[string][]string{}

	rValue := reflect.ValueOf(st)

	if rValue.Kind() == reflect.Pointer {
		rValue = rValue.Elem()
	}

	if rValue.Kind() != reflect.Struct {
		return eris.New("argument is not a struct")
	}

	for i := 0; i < rValue.NumField(); i++ {
		field := rValue.Type().Field(i)

		jsonTag := field.Tag.Get("json")

		key := field.Name

		fieldTagName := strings.TrimSuffix(jsonTag, ",omitempty")

		if fieldTagName != "" {
			key = fieldTagName
		}

		// Get opts by field name
		opts, ok := fieldsOpts[field.Name]
		if !ok {
			// Get opts by field tag name
			opts, ok = fieldsOpts[fieldTagName]
			if !ok {
				// Field not found
				continue
			}
		}

		for _, opt := range opts {
			if msg, ok := opt(rValue.Field(i)); !ok {
				errs[key] = append(errs[key], msg)
			}
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return &ValidateError{errs}
}
