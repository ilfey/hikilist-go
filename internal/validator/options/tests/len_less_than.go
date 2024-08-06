package tests

import (
	"reflect"
	"testing"

	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/stretchr/testify/assert"
)

func TestLenLessThan(t *testing.T) {
	tests := map[any]bool{
		// ptrs
		ptr([]string{}):             true,  // array ptr
		ptr[map[string]string](nil): true,  // map ptr
		ptr(""):                     true,  // string ptr
		ptr([]string{"t"}):          false, // array ptr
		ptr("t"):                    false, // string ptr
		ptr([]string{"t", "t"}):     false, // array ptr
		ptr("tt"):                   false, // string ptr

		ptr(map[string]string{
			"t": "t",
		}): false, // map ptr
		ptr(map[string]string{
			"t": "t",
			"e": "e",
		}): false, // map ptr

		// non-ptrs
		"":   true,  // string
		"t":  false, // string
		"tt": false, // string
	}

	for value, result := range tests {
		rValue := reflect.ValueOf(value)

		assert.Equalf(t, validate(rValue, options.LenLessThan(1)), result, "value: %v result: %v", value, result)
	}

	// []string{}: true
	rValue := reflect.ValueOf([]string{})
	assert.Equal(t, validate(rValue, options.LenLessThan(1)), true)

	// []string{"t"}: false
	rValue = reflect.ValueOf([]string{"t"})
	assert.Equal(t, validate(rValue, options.LenLessThan(1)), false)

	// []string{"t", "t"}: false
	rValue = reflect.ValueOf([]string{"t", "t"})
	assert.Equal(t, validate(rValue, options.LenLessThan(1)), false)

	// map[string]string{}: true
	rValue = reflect.ValueOf(map[string]string{})
	assert.Equal(t, validate(rValue, options.LenLessThan(1)), true)

	// map[string]string{"t": "t"}: false
	rValue = reflect.ValueOf(map[string]string{"t": "t"})
	assert.Equal(t, validate(rValue, options.LenLessThan(1)), false)

	// map[string]string{"t": "t", "e": "e"}: false
	rValue = reflect.ValueOf(map[string]string{"t": "t", "e": "e"})
	assert.Equal(t, validate(rValue, options.LenLessThan(1)), false)
}
