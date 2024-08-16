package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/pkg/models/auth"
	"github.com/stretchr/testify/assert"
)

func TestChangeUsernameModelValidate(t *testing.T) {
	testCases := []struct {
		desc  string
		model auth.ChangeUsernameModel
		ok    bool
	}{
		{
			desc: "Expected valid model",
			model: auth.ChangeUsernameModel{
				Username: "test",
			},
			ok: true,
		},
		{
			desc: "Model without new username",
			model: auth.ChangeUsernameModel{
				Username: "",
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
