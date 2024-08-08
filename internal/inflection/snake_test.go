package inflection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnake(t *testing.T) {
	testCases := []struct {
		desc     string
		expected string
		actual   string
	}{
		{
			desc:     "LowerCase",
			expected: "test",
			actual:   "test",
		},
		{
			desc:     "First char upper case",
			expected: "test",
			actual:   "Test",
		},
		{
			desc:     "First char upper case with number",
			expected: "test123",
			actual:   "Test123",
		},
		{
			desc:     "CamelCase",
			expected: "test_test",
			actual:   "TestTest",
		},
		{
			desc:     "CamelCase with number",
			expected: "test_test123",
			actual:   "TestTest123",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expected, Snake(tC.actual))
		})
	}
}
