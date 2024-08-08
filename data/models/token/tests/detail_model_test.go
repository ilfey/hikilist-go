package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/token"
	"github.com/stretchr/testify/assert"
)

func TestDetailModelGetSQL(t *testing.T) {
	var m token.DetailModel

	sql, args, err := m.GetSQL(map[string]any{
		"id": 1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT id, token, created_at FROM tokens WHERE id = ?",
		sql,
	)
}

func TestDetailModelDeleteSQL(t *testing.T) {
	m := token.DetailModel{
		ID: 1,
	}

	sql, args, err := m.DeleteSQL()
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"DELETE FROM tokens WHERE id = ? RETURNING id",
		sql,
	)
}
