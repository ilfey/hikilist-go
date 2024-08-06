package tests

import (
	"reflect"
	"testing"

	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/stretchr/testify/assert"
)

func TestGreaterThan(t *testing.T) {
	tests := map[any]bool{
		// ptrs
		_nil[int](): false, // nil int ptr
		ptr(0):      false, // int ptr
		ptr(0.0):    false, // float64 ptr
		ptr(1):      false, // int ptr
		ptr(1.):     false, // float64 ptr
		ptr(2):      true,  // int ptr
		ptr(1.1):    true,  // float64 ptr

		// non-ptrs
		0:   false, // int
		0.0: false, // float64
		1:   false, // int
		1.:  false, // float64
		2:   true,  // int
		1.1: true,  // float64
	}

	for value, result := range tests {
		rValue := reflect.ValueOf(value)

		expected := validate(rValue, options.GreaterThan[int64](1))

		assert.Equalf(t, expected, result, "value: %v kind: %v result: %v", value, rValue.Kind(), result)
	}
}
