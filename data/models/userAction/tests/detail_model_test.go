package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/userAction"
	"github.com/stretchr/testify/assert"
)

func TestDetailModelGetSQL(t *testing.T) {
	var m userAction.DetailModel

	sql, args, err := m.GetSQL(map[string]any{
		"id": 1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT id, user_id, title, description, created_at, updated_at FROM user_actions WHERE id = ?",
		sql,
	)
}
