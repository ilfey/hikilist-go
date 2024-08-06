package collectionModels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListModelFillResultsSQL(t *testing.T) {
	var m ListModel

	p := Paginate{}

	assert.NoError(t, p.Validate())

	sql, args, err := m.fillResultsSQL(p.Normalize(), nil)
	assert.NoError(t, err)
	assert.Nil(t, args)

	assert.Equal(
		t,
		"SELECT id, user_id, title, created_at, updated_at FROM collections ORDER BY id DESC LIMIT 10 OFFSET 0",
		sql,
	)

	p = Paginate{
		Page:  2,
		Limit: 5,
	}

	assert.NoError(t, p.Validate())

	sql, args, err = m.fillResultsSQL(p.Normalize(), map[string]any{"is_public": true})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT id, user_id, title, created_at, updated_at FROM collections WHERE is_public = ? ORDER BY id DESC LIMIT 5 OFFSET 5",
		sql,
	)
}

func TestListModelFillCountSQL(t *testing.T) {
	var m ListModel

	sql, args, err := m.fillCountSQL(nil)
	assert.NoError(t, err)
	assert.Nil(t, args)

	assert.Equal(
		t,
		"SELECT COUNT(*) FROM collections",
		sql,
	)

	sql, args, err = m.fillCountSQL(map[string]any{"is_public": true})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT COUNT(*) FROM collections WHERE is_public = ?",
		sql,
	)
}