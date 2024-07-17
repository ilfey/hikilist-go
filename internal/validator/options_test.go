package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptr[T any](val T) *T {
	return &val
}

func _nil[T any]() *T {
	return nil
}

func validate(rValue reflect.Value, opt Option) bool {
	_, ok := opt(rValue)

	return ok
}

func TestNotNil(t *testing.T) {
	tests := map[any]bool{
		// zero ptrs
		ptr(0):      true,  // any ptr
		_nil[int](): false, // nil int ptr
	}

	for value, result := range tests {
		rValue := reflect.ValueOf(value)

		assert.Equalf(t, validate(rValue, NotNil()), result, "value: %v result: %v", value, result)
	}
}

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

		expected := validate(rValue, Required())

		assert.Equalf(t, expected, result, "value: %v kind: %v result: %v", value, rValue.Kind(), result)
	}

	// []string{}: false, // array
	rValue := reflect.ValueOf([]string{})
	assert.Equalf(t, validate(rValue, Required()), false, "value: %v result: %v", rValue, false)

	// map[string]string{}: false, // map
	rValue = reflect.ValueOf(map[string]string{})
	assert.Equalf(t, validate(rValue, Required()), false, "value: %v result: %v", rValue, false)

	// []string{"t"}: true, // array
	rValue = reflect.ValueOf([]string{"t"})
	assert.Equalf(t, validate(rValue, Required()), true, "value: %v result: %v", rValue, true)

	// map[string]string{"t": "t"}: true, // map
	rValue = reflect.ValueOf(map[string]string{"t": "t"})
	assert.Equalf(t, validate(rValue, Required()), true, "value: %v result: %v", rValue, true)
}

func TestLessThat(t *testing.T) {
	tests := map[any]bool{
		// ptrs
		_nil[int](): false, // nil int ptr
		ptr(0):      true,  // int ptr
		ptr(0.0):    true,  // float64 ptr
		ptr(1):      false, // int ptr
		ptr(1.):     false, // float64 ptr
		ptr(2):      false, // int ptr
		ptr(1.1):    false, // float64 ptr

		// non-ptrs
		0:   true,  // int
		0.0: true,  // float64
		1:   false, // int
		1.:  false, // float64
		2:   false, // int
		1.1: false, // float64
	}

	for value, result := range tests {
		rValue := reflect.ValueOf(value)

		expected := validate(rValue, LessThat[int64](1))

		assert.Equalf(t, expected, result, "value: %v kind: %v result: %v", value, rValue.Kind(), result)
	}
}

func TestGreaterThat(t *testing.T) {
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

		expected := validate(rValue, GreaterThat[int64](1))

		assert.Equalf(t, expected, result, "value: %v kind: %v result: %v", value, rValue.Kind(), result)
	}
}

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

		assert.Equalf(t, validate(rValue, LenGreaterThat(1)), result, "value: %v result: %v", value, result)
	}

	// []string{}: false
	rValue := reflect.ValueOf([]string{})
	assert.Equal(t, validate(rValue, LenGreaterThat(1)), false)

	// []string{"t"}: false
	rValue = reflect.ValueOf([]string{"t"})
	assert.Equal(t, validate(rValue, LenGreaterThat(1)), false)

	// []string{"t", "t"}: true
	rValue = reflect.ValueOf([]string{"t", "t"})
	assert.Equal(t, validate(rValue, LenGreaterThat(1)), true)

	// map[string]string{}: false
	rValue = reflect.ValueOf(map[string]string{})
	assert.Equal(t, validate(rValue, LenGreaterThat(1)), false)

	// map[string]string{"t": "t"}: false
	rValue = reflect.ValueOf(map[string]string{"t": "t"})
	assert.Equal(t, validate(rValue, LenGreaterThat(1)), false)

	// map[string]string{"t": "t", "e": "e"}: true
	rValue = reflect.ValueOf(map[string]string{"t": "t", "e": "e"})
	assert.Equal(t, validate(rValue, LenGreaterThat(1)), true)
}

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

		assert.Equalf(t, validate(rValue, LenLessThat(1)), result, "value: %v result: %v", value, result)
	}

	// []string{}: true
	rValue := reflect.ValueOf([]string{})
	assert.Equal(t, validate(rValue, LenLessThat(1)), true)

	// []string{"t"}: false
	rValue = reflect.ValueOf([]string{"t"})
	assert.Equal(t, validate(rValue, LenLessThat(1)), false)

	// []string{"t", "t"}: false
	rValue = reflect.ValueOf([]string{"t", "t"})
	assert.Equal(t, validate(rValue, LenLessThat(1)), false)

	// map[string]string{}: true
	rValue = reflect.ValueOf(map[string]string{})
	assert.Equal(t, validate(rValue, LenLessThat(1)), true)

	// map[string]string{"t": "t"}: false
	rValue = reflect.ValueOf(map[string]string{"t": "t"})
	assert.Equal(t, validate(rValue, LenLessThat(1)), false)

	// map[string]string{"t": "t", "e": "e"}: false
	rValue = reflect.ValueOf(map[string]string{"t": "t", "e": "e"})
	assert.Equal(t, validate(rValue, LenLessThat(1)), false)
}
