package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/user"
	"github.com/stretchr/testify/assert"
)

func TestDetailModelGetSQL(t *testing.T) {
	var m user.DetailModel

	sql, args, err := m.GetSQL(map[string]any{
		"id": 1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT id, username, password, last_online, created_at FROM users WHERE id = ?",
		sql,
	)
}

func TestDetailModelUpdatePasswordSQL(t *testing.T) {
	m := user.DetailModel{
		ID: 1,
	}

	sql, args, err := m.UpdatePasswordSQL()
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"UPDATE users SET password = ? WHERE id = ?",
		sql,
	)
}

func TestDetailModelUpdateLastOnlineSQL(t *testing.T) {
	m := user.DetailModel{
		ID: 1,
	}

	sql, args, err := m.UpdateLastOnlineSQL()
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"UPDATE users SET last_online = ? WHERE id = ?",
		sql,
	)
}

func TestDetailModelDeleteSQL(t *testing.T) {
	m := user.DetailModel{
		ID: 1,
	}

	sql, args, err := m.DeleteSQL()
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"DELETE FROM users WHERE id = ? RETURNING id",
		sql,
	)
}
