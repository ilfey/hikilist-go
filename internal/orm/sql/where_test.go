package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhere(t *testing.T) {
	assert.Equal(t, "t.id = 1", where("t", 1))

	assert.Equal(t, "id = 1", where("t", "id = 1"))

	assert.Equal(t, "t.id IN (1, 2, 3)", where("t", []int{1, 2, 3}))

	assert.Equal(t, "t.id = 1", where("t", map[string]any{"ID": 1}))
}
