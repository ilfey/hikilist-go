package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/user"
	"github.com/stretchr/testify/assert"
)

func TestCreateModelInsertSQL(t *testing.T) {
	m := user.CreateModel{}

	sql, args, err := m.InsertSQL()
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"INSERT INTO users (username,password,created_at) VALUES (?,?,?) RETURNING id",
		sql,
	)
}
