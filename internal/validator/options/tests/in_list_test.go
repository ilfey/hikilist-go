package tests

import (
	"reflect"
	"testing"

	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/stretchr/testify/assert"
)

func TestInList(t *testing.T) {
	rValue := reflect.ValueOf("t")
	assert.Equal(t, validate(rValue, options.InList([]string{"t"})), true)
	assert.Equal(t, validate(rValue, options.InList([]string{"t", "e"})), true)
	assert.Equal(t, validate(rValue, options.InList([]string{"e"})), false)

	rValue = reflect.ValueOf(int64(1))
	assert.Equal(t, validate(rValue, options.InList([]int64{1})), true)
	assert.Equal(t, validate(rValue, options.InList([]int64{1, 2})), true)
	assert.Equal(t, validate(rValue, options.InList([]int64{2})), false)

	rValue = reflect.ValueOf(float64(1.1))
	assert.Equal(t, validate(rValue, options.InList([]float64{1.1})), true)
	assert.Equal(t, validate(rValue, options.InList([]float64{1.1, 2.2})), true)
	assert.Equal(t, validate(rValue, options.InList([]float64{2.2})), false)
}
