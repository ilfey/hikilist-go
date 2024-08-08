package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/user"
	"github.com/stretchr/testify/assert"
)

func TestListModelFillResultsSQL(t *testing.T) {
	var lm user.ListModel

	p := user.Paginate{}

	sql, args, err := lm.FillResultsSQL(p.Normalize(), nil)
	assert.NoError(t, err)
	assert.Nil(t, args)

	assert.Equal(
		t,
		"SELECT id, username, created_at FROM users ORDER BY id DESC LIMIT 10 OFFSET 0",
		sql,
	)

	sql, args, err = lm.FillResultsSQL(p.Normalize(), map[string]any{"username": "test"})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT id, username, created_at FROM users WHERE username = ? ORDER BY id DESC LIMIT 10 OFFSET 0",
		sql,
	)
}

func TestListModelFillCountSQL(t *testing.T) {
	var lm user.ListModel

	sql, args, err := lm.FillCountSQL(nil)
	assert.NoError(t, err)
	assert.Nil(t, args)

	assert.Equal(
		t,
		"SELECT COUNT(*) FROM users",
		sql,
	)

	sql, args, err = lm.FillCountSQL(map[string]any{"username": "test"})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT COUNT(*) FROM users WHERE username = ?",
		sql,
	)
}