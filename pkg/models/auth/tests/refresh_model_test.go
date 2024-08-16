package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/pkg/models/auth"
	"github.com/stretchr/testify/assert"
)

func TestRefreshModelValidate(t *testing.T) {
	testCases := []struct {
		desc  string
		model auth.RefreshModel
		ok    bool
	}{
		{
			desc: "Expected valid model",
			model: auth.RefreshModel{
				Refresh: "test",
			},
			ok: true,
		},
		{
			desc: "Model without refresh",
			model: auth.RefreshModel{
				Refresh: "",
			},
			ok: false,
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
