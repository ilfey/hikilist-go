package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertStruct(t *testing.T) {
	user := &User{
		ID: 1,

		Username: "test",

		Password: "test",
	}

	is := InsertFromStruct(user)

	assert.Equal(
		t,
		"INSERT INTO users (id, username, email, password) VALUES (1, 'test', NULL, 'test');",
		is.SQL(),
	)
}
