package tests

import (
	"reflect"
	"testing"

	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/stretchr/testify/assert"
)

func TestLenLessThan(t *testing.T) {
	testCases := []struct {
		desc  string
		value any
		ok    bool
	}{
		{
			desc:  "Empty array ptr",
			value: ptr([]string{}),
			ok:    true,
		},
		{
			desc:  "Empty map ptr",
			value: ptr(map[string]string{}),
			ok:    true,
		},
		{
			desc:  "Empty string ptr",
			value: ptr(""),
			ok:    true,
		},
		{
			desc:  "Array 1 ptr",
			value: ptr([]string{"t"}),
			ok:    false,
		},
		{
			desc:  "String ptr",
			value: ptr("t"),
			ok:    false,
		},
		{
			desc:  "Array 2 ptr",
			value: ptr([]string{"t", "t"}),
			ok:    false,
		},
		{
			desc:  "String 2 ptr",
			value: ptr("tt"),
			ok:    false,
		},

		{
			desc: "Empty map ptr",
			value: ptr(map[string]string{
				"t": "t",
			}),
			ok: false,
		},
		{
			desc: "Map 2 ptr",
			value: ptr(map[string]string{
				"t": "t",
				"e": "e",
			}),
			ok: false,
		},

		// non-ptrs
		{
			desc:  "Empty string",
			value: "",
			ok:    true,
		},
		{
			desc:  "String",
			value: "t",
			ok:    false,
		},
		{
			desc:  "String 2",
			value: "tt",
			ok:    false,
		},
		{
			desc:  "Empty array",
			value: []string{},
			ok:    true,
		},
		{
			desc:  "Array",
			value: []string{"t"},
			ok:    false,
		},
		{
			desc:  "Array 2",
			value: []string{"t", "t"},
			ok:    false,
		},
		{
			desc:  "Empty map",
			value: map[string]string{},
			ok:    true,
		},
		{
			desc:  "Map",
			value: map[string]string{"t": "t"},
			ok:    false,
		},
		{
			desc:  "Map 2",
			value: map[string]string{"t": "t", "e": "e"},
			ok:    false,
		},

		// invalid type
		{
			desc:  "Int 0",
			value: 0,
			ok:    false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rValue := reflect.ValueOf(tC.value)

			isValid := validate(rValue, options.LenLessThan(1))

			if tC.ok {
				assert.True(t, isValid)

				return
			}

			assert.False(t, isValid)
		})
	}
}
