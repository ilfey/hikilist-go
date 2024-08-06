package collectionModels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateModelValidate(t *testing.T) {
	valids := []CreateModel{
		{
			Title: "test",
		},
		{
			Title: "very_very_very_very_very_long_name",
		},
	}

	for _, m := range valids {
		assert.NoError(t, m.Validate())
	}

	invalid := CreateModel{}

	assert.Error(t, invalid.Validate())
}

func TestCreateModelInsertSQL(t *testing.T) {
	m := CreateModel{
		Title: "test",
	}

	sql, args, err := m.insertSQL()
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"INSERT INTO collections (title,user_id,description,is_public,created_at) VALUES (?,?,?,?,?) RETURNING id",
		sql,
	)
}
