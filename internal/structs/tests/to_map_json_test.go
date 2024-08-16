package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/internal/structs"
	"github.com/stretchr/testify/assert"
)

type TestJSONStruct struct {
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
	Field3 string `json:"-"`
}

func TestToMapJSON(t *testing.T) {
	st := TestJSONStruct{1, "test", "more of a test"}

	_, err := structs.ToMapJSON("test")
	assert.Error(t, err)

	var result map[string]any

	result, err = structs.ToMapJSON(st)
	assert.NoError(t, err)

	assert.Equal(t, result["field1"], st.Field1)
	assert.Equal(t, result["field2"], st.Field2)
	_, ok := result["field3"]

	assert.False(t, ok)
}
