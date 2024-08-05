package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateStruct(t *testing.T) {
	user := &User{
		Username: "test",
	}

	us := UpdateFromStruct(user).IgnoreZeros()

	assert.Equal(t, "UPDATE users SET users.username = 'test';", us.SQL())

}
