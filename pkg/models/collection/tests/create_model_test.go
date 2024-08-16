package collection

import (
	"testing"

	"github.com/ilfey/hikilist-go/pkg/models/collection"
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
				Title:    "test",
				IsPublic: ptr(true),
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
