package animeModels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T {
	return &v
}

func TestCreateModelValidate(t *testing.T) {
	valids := []CreateModel{
		{
			Title:            "test",
			EpisodesReleased: 1,
		},
		{
			Title:       "test",
			Description: ptr("test"),
			Poster:      ptr("test"),

			Episodes:         ptr(uint(1)),
			EpisodesReleased: 0,

			MalID:   ptr(uint(1)),
			ShikiID: ptr(uint(1)),
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
		Title:            "test",
		EpisodesReleased: 1,
	}

	sql, args, err := m.insertSQL()
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"INSERT INTO animes (title,description,poster,episodes,episodes_released,mal_id,shiki_id,created_at) VALUES (?,?,?,?,?,?,?,?) RETURNING id",
		sql,
	)
}
