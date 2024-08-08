package collection

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/collection"
	"github.com/stretchr/testify/assert"
)

func TestListModelFillResultsSQLWithoutConds(t *testing.T) {
	var m collection.ListModel

	p := collection.Paginate{}

	assert.NoError(t, p.Validate())

	sql, args, err := m.FillResultsSQL(p.Normalize(), nil)
	assert.NoError(t, err)
	assert.Nil(t, args)

	assert.Equal(
		t,
		"SELECT id, user_id, title, created_at, updated_at FROM collections ORDER BY id DESC LIMIT 10 OFFSET 0",
		sql,
	)
}


func TestListModelFillResultsSQLWithConds(t *testing.T) {
	var m collection.ListModel

	p := collection.Paginate{}

	sql, args, err := m.FillResultsSQL(p.Normalize(), map[string]any{"is_public": true})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT id, user_id, title, created_at, updated_at FROM collections WHERE is_public = ? ORDER BY id DESC LIMIT 10 OFFSET 0",
		sql,
	)
}

func TestListModelFillCountSQLWithoutConds(t *testing.T) {
	var m collection.ListModel

	sql, args, err := m.FillCountSQL(nil)
	assert.NoError(t, err)
	assert.Nil(t, args)

	assert.Equal(
		t,
		"SELECT COUNT(*) FROM collections",
		sql,
	)
}

func TestListModelFillCountSQLWithConds(t *testing.T) {
	var m collection.ListModel

	sql, args, err := m.FillCountSQL(map[string]any{"is_public": true})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT COUNT(*) FROM collections WHERE is_public = ?",
		sql,
	)
}
