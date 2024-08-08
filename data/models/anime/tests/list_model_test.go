package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/anime"
	"github.com/stretchr/testify/assert"
)

func TestListModelFillResultsSQLWithoutConds(t *testing.T) {
	var m anime.ListModel

	p := anime.Paginate{}

	assert.NoError(t, p.Validate())

	sql, args, err := m.FillResultsSQL(p.Normalize(), nil)
	assert.NoError(t, err)
	assert.Nil(t, args)

	assert.Equal(
		t,
		"SELECT id, title, poster, episodes, episodes_released FROM animes ORDER BY id DESC LIMIT 10 OFFSET 0",
		sql,
	)
}

func TestListModelFillResultsSQLWithConds(t *testing.T) {
	var m anime.ListModel

	p := anime.Paginate{}

	assert.NoError(t, p.Validate())

	sql, args, err := m.FillResultsSQL(p.Normalize(), map[string]any{"episodes": 12})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT id, title, poster, episodes, episodes_released FROM animes WHERE episodes = ? ORDER BY id DESC LIMIT 10 OFFSET 0",
		sql,
	)
}

func TestListModelFillCountSQLWithoutConds(t *testing.T) {
	var m anime.ListModel

	sql, args, err := m.FillCountSQL(nil)
	assert.NoError(t, err)
	assert.Nil(t, args)

	assert.Equal(
		t,
		"SELECT COUNT(*) FROM animes",
		sql,
	)
}

func TestListModelFillCountSQLWithConds(t *testing.T) {
	var m anime.ListModel

	sql, args, err := m.FillCountSQL(map[string]any{"episodes": 12})
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT COUNT(*) FROM animes WHERE episodes = ?",
		sql,
	)
}

func TestListModelFillFromCollectionResultsSQL(t *testing.T) {
	var m anime.ListModel

	p := anime.Paginate{}

	assert.NoError(t, p.Validate())

	sql, args, err := m.FillFromCollectionResultsSQL(p.Normalize(), 1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT id, title, poster, episodes, episodes_released FROM animes_collections "+
			"JOIN animes ON animes.id = animes_collections.anime_id "+
			"WHERE collection_id = (SELECT id FROM collections WHERE id = ? AND (is_public = TRUE OR user_id = ?)) "+
			"ORDER BY id DESC LIMIT 10 OFFSET 0",
		sql,
	)
}

func TestListModelFillFromCollectionCountSQL(t *testing.T) {
	var m anime.ListModel

	sql, args, err := m.FillFromCollectionCountSQL(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"SELECT COUNT(*) FROM animes_collections "+
			"JOIN collections ON collections.id = animes_collections.collection_id "+
			"WHERE collection_id = ? AND (is_public = TRUE OR user_id = ?)",
		sql,
	)
}
