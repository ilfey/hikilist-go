package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/token"
	"github.com/stretchr/testify/assert"
)

func TestCreateModelInsertSQL(t *testing.T) {
	m := token.CreateModel{}

	sql, args, err := m.InsertSQL()
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"INSERT INTO tokens (token,created_at) VALUES (?,?) RETURNING id",
		sql,
	)
}
