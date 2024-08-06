package tests

import (
	"reflect"
	"testing"

	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/stretchr/testify/assert"
)

func TestNotNil(t *testing.T) {
	tests := map[any]bool{
		// zero ptrs
		ptr(0):      true,  // any ptr
		_nil[int](): false, // nil int ptr
	}

	for value, result := range tests {
		rValue := reflect.ValueOf(value)

		assert.Equalf(t, validate(rValue, options.NotNil()), result, "value: %v result: %v", value, result)
	}
}
