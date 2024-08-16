package structs

import (
	"reflect"
	"strings"

	"github.com/rotisserie/eris"
)

func ToMapJSON(st any) (map[string]any, error) {
	rValue := reflect.ValueOf(st)

	if rValue.Kind() == reflect.Pointer {
		rValue = rValue.Elem()
	}

	if rValue.Kind() != reflect.Struct {
		return nil, eris.New("argument is not a struct")
	}

	rType := rValue.Type()

	result := make(map[string]any)

	for i := 0; i < rValue.NumField(); i++ {
		field := rType.Field(i)

		jsonTag := field.Tag.Get("json")

		if jsonTag == "-" {
			continue
		}

		key := field.Name

		fieldTagName := strings.TrimSuffix(jsonTag, ",omitempty")

		if fieldTagName != "" {
			key = fieldTagName
		}

		result[key] = rValue.Field(i).Interface()
	}

	return result, nil
}
