package animeModels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetailModelGetSQL(t *testing.T) {
	var m DetailModel

	sql, args, err := m.getSQL(map[string]any{
		"id": 1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT id, title, description, poster, episodes, episodes_released, mal_id, shiki_id, created_at FROM animes WHERE id = ?",
		sql,
	)
}
