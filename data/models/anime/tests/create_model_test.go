package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/anime"
	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T {
	return &v
}

func TestCreateModelValidate(t *testing.T) {
	testCases := []struct {
		desc  string
		model anime.CreateModel
		ok    bool
	}{
		{
			desc: "Model with only title",
			model: anime.CreateModel{
				Title: "test",
			},
			ok: true,
		},
		{
			desc: "Model with title and description",
			model: anime.CreateModel{
				Title:       "test",
				Description: ptr("test"),
			},
			ok: true,
		},
		{
			desc: "Model with empty title",
			model: anime.CreateModel{
				Title: "",
			},
			ok: false,
		},
		{
			desc: "Model with long title",
			model: anime.CreateModel{
				Title: string(make([]byte, 256)),
			},
			ok: false,
		},
		{
			desc:  "Empty model",
			model: anime.CreateModel{},
			ok:    false,
		},
		{
			desc: "Model with mal_id",
			model: anime.CreateModel{
				Title: "test",
				MalID: ptr(uint(1)),
			},
			ok: true,
		},
		{
			desc: "Model with shiki_id",
			model: anime.CreateModel{
				Title:   "test",
				ShikiID: ptr(uint(1)),
			},
			ok: true,
		},
		{
			desc: "Model with mal_id and shiki_id",
			model: anime.CreateModel{
				Title:   "test",
				MalID:   ptr(uint(1)),
				ShikiID: ptr(uint(1)),
			},
			ok: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.ok {
				assert.NoError(t, tC.model.Validate())

				return
			}

			assert.Error(t, tC.model.Validate())
		})
	}
}

func TestCreateModelInsertSQL(t *testing.T) {
	m := anime.CreateModel{
		Title:            "test",
		EpisodesReleased: 1,
	}

	sql, args, err := m.InsertSQL()
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"INSERT INTO animes (title,description,poster,episodes,episodes_released,mal_id,shiki_id,created_at) VALUES (?,?,?,?,?,?,?,?) RETURNING id",
		sql,
	)
}
