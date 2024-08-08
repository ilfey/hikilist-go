package collection

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/collection"
	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T {
	return &v
}

func TestCreateModelValidate(t *testing.T) {
	testCases := []struct {
		desc  string
		model collection.CreateModel
		ok    bool
	}{
		{
			desc: "Model with only title",
			model: collection.CreateModel{
				Title: "test",
			},
			ok: true,
		},
		{
			desc: "Model with title and description",
			model: collection.CreateModel{
				Title:       "test",
				Description: ptr("test"),
			},
			ok: true,
		},
		{
			desc: "Model with long title",
			model: collection.CreateModel{
				Title: string(make([]byte, 256)),
			},
			ok: false,
		},
		{
			desc:  "Empty model",
			model: collection.CreateModel{},
			ok:    false,
		},
		{
			desc: "Model with title and is_public",
			model: collection.CreateModel{
				Title:       "test",
				IsPublic:    ptr(true),
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
	m := collection.CreateModel{
		Title: "test",
	}

	sql, args, err := m.InsertSQL()
	assert.NoError(t, err)
	assert.NotNil(t, args)

	assert.Equal(
		t,
		"INSERT INTO collections (title,user_id,description,is_public,created_at) VALUES (?,?,?,?,?) RETURNING id",
		sql,
	)
}
