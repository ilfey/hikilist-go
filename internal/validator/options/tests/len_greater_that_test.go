package tests

import (
	"reflect"
	"testing"

	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/stretchr/testify/assert"
)

func TestLenGreaterThan(t *testing.T) {
	testCases := []struct {
		desc  string
		value any
		ok    bool
	}{
		{
			desc:  "Empty array ptr",
			value: ptr([]string{}),
			ok:    false,
		},
		{
			desc:  "Empty map ptr",
			value: ptr(map[string]string{}),
			ok:    false,
		},
		{
			desc:  "Empty string ptr",
			value: ptr(""),
			ok:    false,
		},
		{
			desc:  "Array ptr",
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
			ok:    true,
		},
		{
			desc:  "String 2 ptr",
			value: ptr("tt"),
			ok:    true,
		},

		{
			desc: "Map ptr",
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
			ok: true,
		},

		// non-ptrs
		{
			desc:  "Empty string",
			value: "",
			ok:    false,
		},
		{
			desc:  "String",
			value: "t",
			ok:    false,
		},
		{
			desc:  "String 2",
			value: "tt",
			ok:    true,
		},
		{
			desc:  "Empty array",
			value: []string{},
			ok:    false,
		},
		{
			desc:  "Array",
			value: []string{"t"},
			ok:    false,
		},
		{
			desc:  "Array 2",
			value: []string{"t", "t"},
			ok:    true,
		},
		{
			desc:  "Empty map",
			value: map[string]string{},
			ok:    false,
		},
		{
			desc:  "Map",
			value: map[string]string{"t": "t"},
			ok:    false,
		},
		{
			desc:  "Map 2",
			value: map[string]string{"t": "t", "e": "e"},
			ok:    true,
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

			isValid := validate(rValue, options.LenGreaterThan(1))

			if tC.ok {
				assert.True(t, isValid)

				return
			}

			assert.False(t, isValid)
		})
	}
}
