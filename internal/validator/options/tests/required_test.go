package tests

import (
	"reflect"
	"testing"

	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/stretchr/testify/assert"
)

func TestRequired(t *testing.T) {
	tests := map[any]bool{
		// zero ptrs
		_nil[int]():              false, // nil int ptr
		ptr(0):                   false, // int ptr
		ptr(0.):                  false, // float64 ptr
		ptr(""):                  false, // string ptr
		ptr(false):               false, // bool ptr
		ptr([]string{}):          false, // array ptr
		ptr(map[string]string{}): false, // map ptr

		// non-zero ptrs
		ptr(1):                           true, // int ptr
		ptr(1.1):                         true, // float64 ptr
		ptr("t"):                         true, // string ptr
		ptr(true):                        true, // bool ptr
		ptr([]string{"t"}):               true, // array ptr
		ptr(map[string]string{"t": "t"}): true, // map ptr

		// zeros
		0:     false, // int
		0.0:   false, // float64
		"":    false, // string
		false: false, // bool

		// non-zeros
		1:    true, // int
		1.1:  true, // float64
		"t":  true, // string
		true: true, // bool
	}

	for value, result := range tests {
		rValue := reflect.ValueOf(value)

		expected := validate(rValue, options.Required())

		assert.Equalf(t, expected, result, "value: %v kind: %v result: %v", value, rValue.Kind(), result)
	}

	// []string{}: false, // array
	rValue := reflect.ValueOf([]string{})
	assert.Equalf(t, validate(rValue, options.Required()), false, "value: %v result: %v", rValue, false)

	// map[string]string{}: false, // map
	rValue = reflect.ValueOf(map[string]string{})
	assert.Equalf(t, validate(rValue, options.Required()), false, "value: %v result: %v", rValue, false)

	// []string{"t"}: true, // array
	rValue = reflect.ValueOf([]string{"t"})
	assert.Equalf(t, validate(rValue, options.Required()), true, "value: %v result: %v", rValue, true)

	// map[string]string{"t": "t"}: true, // map
	rValue = reflect.ValueOf(map[string]string{"t": "t"})
	assert.Equalf(t, validate(rValue, options.Required()), true, "value: %v result: %v", rValue, true)
}
