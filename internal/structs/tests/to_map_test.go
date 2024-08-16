package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/internal/structs"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Field1 int
	Field2 string
	Field3 string
}

func TestToMap(t *testing.T) {
	st := TestStruct{1, "test", "more of a test"}

	_, err := structs.ToMap("test")
	assert.Error(t, err)

	var result map[string]any

	result, err = structs.ToMap(st)
	assert.NoError(t, err)

	assert.Equal(t, result["Field1"], st.Field1)
	assert.Equal(t, result["Field2"], st.Field2)
	assert.Equal(t, result["Field3"], st.Field3)
}
