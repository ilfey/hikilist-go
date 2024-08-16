package structs

import (
	"reflect"

	"github.com/rotisserie/eris"
)

func ToMap(st any) (map[string]any, error) {
	rValue := reflect.ValueOf(st)

	if rValue.Kind() == reflect.Pointer {
		rValue = rValue.Elem()
	}

	if rValue.Kind() != reflect.Struct {
		return nil, eris.New("argument is not a struct")
	}

	rType := rValue.Type()

	result := map[string]any{}

	for i := 0; i < rValue.NumField(); i++ {
		fieldName := rType.Field(i).Name

		result[fieldName] = rValue.Field(i).Interface()
	}

	return result, nil
}
