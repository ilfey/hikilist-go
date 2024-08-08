package tests

import (
	"reflect"
	"testing"

	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/stretchr/testify/assert"
)

func TestRequired(t *testing.T) {
	testCases := []struct {
		desc  string
		value any
		ok    bool
	}{
		{
			desc:  "Nil int ptr",
			value: _nil[int](),
			ok:    false,
		},
		{
			desc:  "Int 0 ptr",
			value: ptr(0),
			ok:    false,
		},
		{
			desc:  "Float64 0.0 ptr",
			value: ptr(0.0),
			ok:    false,
		},
		{
			desc:  "Int 1 ptr",
			value: ptr(1),
			ok:    false,
		},
		{
			desc:  "Float64 1.0 ptr",
			value: ptr(1.0),
			ok:    false,
		},
		{
			desc:  "Int 2 ptr",
			value: ptr(2),
			ok:    true,
		},
		{
			desc:  "Float64 1.1 ptr",
			value: ptr(1.1),
			ok:    true,
		},
		// non-ptrs
		{
			desc:  "Int 0",
			value: 0,
			ok:    false,
		},
		{
			desc:  "Float64 0.0",
			value: 0.0,
			ok:    false,
		},
		{
			desc:  "Int 1",
			value: 1,
			ok:    false,
		},
		{
			desc:  "Float64 1.0",
			value: 1.0,
			ok:    false,
		},
		{
			desc:  "Int 2",
			value: 2,
			ok:    true,
		},
		{
			desc:  "Float64 1.1",
			value: 1.1,
			ok:    true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rValue := reflect.ValueOf(tC.value)

			isValid := validate(rValue, options.GreaterThan[int64](1))

			if tC.ok {
				assert.True(t, isValid)

				return
			}

			assert.False(t, isValid)
		})
	}
}
