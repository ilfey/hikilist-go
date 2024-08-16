package repositories

import (
	"reflect"
	"strings"

	"github.com/ilfey/hikilist-go/internal/postgres"
)

type DBRW interface {
	postgres.Read
	postgres.Write
}

func updateModelToMap(st any) map[string]any {
	rValue := reflect.ValueOf(st)

	if rValue.Kind() == reflect.Pointer {
		rValue = rValue.Elem()
	}

	if rValue.Kind() != reflect.Struct {
		panic("argument is not a struct")
	}

	rType := rValue.Type()

	result := make(map[string]any)

	for i := 0; i < rValue.NumField(); i++ {
		field := rType.Field(i)

		jsonTag := field.Tag.Get("json")

		if jsonTag == "-" || rValue.Field(i).IsNil() {
			continue
		}

		key := field.Name

		fieldTagName := strings.TrimSuffix(jsonTag, ",omitempty")

		if fieldTagName != "" {
			key = fieldTagName
		}

		result[key] = rValue.Field(i).Interface()
	}

	return result
}
