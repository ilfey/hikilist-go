package animeModels

import (
	"testing"

	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
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
		"SELECT id, title, poster, episodes, episodes_released FROM animes ORDER BY id DESC LIMIT 10 OFFSET 0",
		sql,
	)

	p = Paginate{
		Order: baseModels.OrderField("episodes_released"),
		Page:  2,
		Limit: 5,
	}

	assert.NoError(t, p.Validate())

	sql, args, err = m.fillResultsSQL(p.Normalize(), map[string]any{"episodes": 12})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT id, title, poster, episodes, episodes_released FROM animes WHERE episodes = ? ORDER BY episodes_released ASC LIMIT 5 OFFSET 5",
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
		"SELECT COUNT(*) FROM animes",
		sql,
	)

	sql, args, err = m.fillCountSQL(map[string]any{"episodes": 12})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT COUNT(*) FROM animes WHERE episodes = ?",
		sql,
	)
}
