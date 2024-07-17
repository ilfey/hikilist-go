package validator

import (
	"reflect"
	"strings"
)

func Validate(st any, fieldsOpts map[string][]Option) ValidateError {
	errs := map[string][]string{}

	rv := reflect.ValueOf(st)

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i).Name

		fieldTagName := strings.TrimSuffix(rv.Type().Field(i).Tag.Get("json"), ",omitempty")
		if fieldTagName == "" {
			fieldTagName = field
		}

		for _, opt := range fieldsOpts[field] {
			if msg, ok := opt(rv.Field(i)); !ok {
				errs[fieldTagName] = append(errs[fieldTagName], msg)
			}
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return &errorImpl{errs}
}
