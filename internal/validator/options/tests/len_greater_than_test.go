package tests

import (
	"reflect"
	"testing"

	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/stretchr/testify/assert"
)

func TestLenGreaterThan(t *testing.T) {
	tests := map[any]bool{
		// ptrs
		ptr([]string{}):             false, // array ptr
		ptr[map[string]string](nil): false, // map ptr
		ptr(""):                     false, // string ptr
		ptr([]string{"t"}):          false, // array ptr
		ptr("t"):                    false, // string ptr
		ptr([]string{"t", "t"}):     true,  // array ptr
		ptr("tt"):                   true,  // string ptr

		ptr(map[string]string{
			"t": "t",
		}): false, // map ptr
		ptr(map[string]string{
			"t": "t",
			"e": "e",
		}): true, // map ptr

		// non-ptrs
		"":   false, // string
		"t":  false, // string
		"tt": true,  // string
	}

	for value, result := range tests {
		rValue := reflect.ValueOf(value)

		assert.Equalf(t, validate(rValue, options.LenGreaterThan(1)), result, "value: %v result: %v", value, result)
	}

	// []string{}: false
	rValue := reflect.ValueOf([]string{})
	assert.Equal(t, validate(rValue, options.LenGreaterThan(1)), false)

	// []string{"t"}: false
	rValue = reflect.ValueOf([]string{"t"})
	assert.Equal(t, validate(rValue, options.LenGreaterThan(1)), false)

	// []string{"t", "t"}: true
	rValue = reflect.ValueOf([]string{"t", "t"})
	assert.Equal(t, validate(rValue, options.LenGreaterThan(1)), true)

	// map[string]string{}: false
	rValue = reflect.ValueOf(map[string]string{})
	assert.Equal(t, validate(rValue, options.LenGreaterThan(1)), false)

	// map[string]string{"t": "t"}: false
	rValue = reflect.ValueOf(map[string]string{"t": "t"})
	assert.Equal(t, validate(rValue, options.LenGreaterThan(1)), false)

	// map[string]string{"t": "t", "e": "e"}: true
	rValue = reflect.ValueOf(map[string]string{"t": "t", "e": "e"})
	assert.Equal(t, validate(rValue, options.LenGreaterThan(1)), true)
}
