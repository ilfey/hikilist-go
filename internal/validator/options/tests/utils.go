package tests

import (
	"reflect"

	"github.com/ilfey/hikilist-go/internal/validator/options"
)

func ptr[T any](val T) *T {
	return &val
}

func _nil[T any]() *T {
	return nil
}

func validate(rValue reflect.Value, opt options.Option) bool {
	_, ok := opt(rValue)

	return ok
}
