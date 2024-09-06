package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ilfey/hikilist-go/pkg/validator/options"
)

// Opts key may be a struct field name or a json tag.
type Opts map[string][]options.Option

/*
Validate validates struct by fields.
Validate returns map of errors and success.

Struct must be a struct or pointer to struct.

Usage:

	type MyStruct struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	validator.Validate(&MyStruct{}, validator.Opts{
		"Name": {options.Required()},
		"age":  {options.Required()},
	})
*/
func Validate(st any, fieldsOpts Opts) (map[string][]string, bool) {
	errs := map[string][]string{}

	rValue := reflect.ValueOf(st)

	if rValue.Kind() == reflect.Pointer {
		rValue = rValue.Elem()
	}

	if rValue.Kind() != reflect.Struct {
		return nil, false
	}

	for i := 0; i < rValue.NumField(); i++ {
		field := rValue.Type().Field(i)

		jsonTag := field.Tag.Get("json")

		key := field.Name

		fieldTagName := strings.TrimSuffix(jsonTag, ",omitempty")

		if fieldTagName != "" {
			key = fieldTagName
		}

		// Detail opts by field name
		opts, ok := fieldsOpts[field.Name]
		if !ok {
			// Detail opts by field tag name
			opts, ok = fieldsOpts[fieldTagName]
			if !ok {
				// Field not found
				continue
			}
		}

		for _, opt := range opts {
			if msg, ok := opt(rValue.Field(i)); !ok {
				errs[key] = append(errs[key], fmt.Sprintf(msg, key))
			}
		}
	}

	if len(errs) == 0 {
		return nil, true
	}

	return errs, false
}
