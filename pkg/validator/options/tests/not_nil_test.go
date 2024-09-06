package tests

import (
	"reflect"
	"testing"

	"github.com/ilfey/hikilist-go/pkg/validator/options"
	"github.com/stretchr/testify/assert"
)

func TestNotNil(t *testing.T) {
	testCases := []struct {
		desc  string
		value any
		ok    bool
	}{
		{
			desc:  "Int ptr",
			value: ptr(0),
			ok:    true,
		},
		{
			desc:  "Nil int ptr",
			value: _nil[int](),
			ok:    false,
		},
		// non-ptr
		{
			desc:  "Int",
			value: 0,
			ok:    false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rValue := reflect.ValueOf(tC.value)

			isValid := validate(rValue, options.NotNil())

			if tC.ok {
				assert.True(t, isValid)

				return
			}

			assert.False(t, isValid)
		})
	}
}
