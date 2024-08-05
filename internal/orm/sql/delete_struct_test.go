package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteStruct(t *testing.T) {
	ds := DeleteFromStruct(&User{})

	assert.Equal(t, "DELETE FROM users;", ds.SQL())
}

func TestDeleteStructWhere(t *testing.T) {
	ds := DeleteFromStruct(&User{})

	ds.Where("id = 1")

	assert.Equal(t, "DELETE FROM users WHERE id = 1;", ds.SQL())
}