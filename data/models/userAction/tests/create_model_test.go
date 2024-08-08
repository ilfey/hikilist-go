package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/userAction"
	"github.com/stretchr/testify/assert"
)

func TestCreateModelInsertSQL(t *testing.T) {
	m := userAction.CreateModel{
		UserID:      1,
		Title:       "test",
		Description: "test",
	}

	sql, args, err := m.InsertSQL()
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"INSERT INTO user_actions (user_id,title,description,created_at) VALUES (?,?,?,?) RETURNING id",
		sql,
	)
}
