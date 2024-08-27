// nolint: unused
package tests

import (
	"reflect"

	"github.com/ilfey/hikilist-go/internal/validator/options"
)

// Returns pointer of value.
func ptr[T any](val T) *T {
	return &val
}

// Returns nil of type T.
func _nil[T any]() *T {
	return nil
}

func validate(rValue reflect.Value, opt options.Option) bool {
	_, ok := opt(rValue)

	return ok
}
