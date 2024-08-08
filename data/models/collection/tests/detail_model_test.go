package collection

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/collection"
	"github.com/stretchr/testify/assert"
)

func TestDetailModelGetSQL(t *testing.T) {
	var m collection.DetailModel

	sql, args, err := m.GetSQL(map[string]any{
		"id": 1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT id, title, user_id, description, is_public, created_at, updated_at FROM collections WHERE id = ?",
		sql,
	)
}
