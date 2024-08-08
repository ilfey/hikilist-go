package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/userAction"
	"github.com/stretchr/testify/assert"
)

func TestListModelFillResultsSQL(t *testing.T) {
	var lm userAction.ListModel

	p := userAction.Paginate{}

	sql, args, err := lm.FillResultsSQL(p.Normalize(), nil)
	assert.NoError(t, err)
	assert.Nil(t, args)

	assert.Equal(
		t,
		"SELECT id, title, description, created_at FROM user_actions LIMIT 10 OFFSET 0",
		sql,
	)

	sql, args, err = lm.FillResultsSQL(p.Normalize(), map[string]any{"user_id": 1})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT id, title, description, created_at FROM user_actions WHERE user_id = ? LIMIT 10 OFFSET 0",
		sql,
	)
}

func TestListModelFillCountSQL(t *testing.T) {
	var lm userAction.ListModel

	sql, args, err := lm.FillCountSQL(nil)
	assert.NoError(t, err)
	assert.Nil(t, args)

	assert.Equal(
		t,
		"SELECT COUNT(*) FROM user_actions",
		sql,
	)

	sql, args, err = lm.FillCountSQL(map[string]any{"user_id": 1})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT COUNT(*) FROM user_actions WHERE user_id = ?",
		sql,
	)
}
