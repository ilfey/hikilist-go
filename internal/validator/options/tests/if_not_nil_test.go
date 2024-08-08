package tests

import (
	"reflect"
	"testing"

	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/stretchr/testify/assert"
)

func TestIfNotNilWithOneOpt(t *testing.T) {
	testCases := []struct {
		desc  string
		value any
		ok    bool
	}{
		{
			desc:  "Nil int ptr",
			value: _nil[int](),
			ok:    true,
		},
		{
			desc:  "Int 0 ptr",
			value: ptr(0),
			ok:    true,
		},
		{
			desc:  "Float64 0.0 ptr",
			value: ptr(0.0),
			ok:    true,
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
			ok:    false,
		},
		{
			desc:  "Float64 1.1 ptr",
			value: ptr(1.1),
			ok:    false,
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
			ok:    false,
		},
		{
			desc:  "Float64 1.1",
			value: 1.1,
			ok:    false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rValue := reflect.ValueOf(tC.value)

			isValid := validate(rValue, options.IfNotNil(options.LessThan[int64](1)))

			if tC.ok {
				assert.True(t, isValid)

				return
			}

			assert.False(t, isValid)
		})
	}
}

func TestIfNotNilWithTwoOpts(t *testing.T) {
	testCases := []struct {
		desc  string
		value any
		ok    bool
	}{
		{
			desc:  "Nil int ptr",
			value: _nil[int](),
			ok:    true,
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
			desc:  "Int -1 ptr",
			value: ptr(-1),
			ok:    true,
		},
		{
			desc:  "Float64 0.5 ptr",
			value: ptr(0.5),
			ok:    true,
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
			ok:    false,
		},
		{
			desc:  "Float64 1.1 ptr",
			value: ptr(1.1),
			ok:    false,
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
			ok:    false,
		},
		{
			desc:  "Float64 1.1",
			value: 1.1,
			ok:    false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rValue := reflect.ValueOf(tC.value)

			isValid := validate(rValue, options.IfNotNil(
				options.LessThan[int64](1),
				options.Required(),
			))

			if tC.ok {
				assert.True(t, isValid)

				return
			}

			assert.False(t, isValid)
		})
	}
}
