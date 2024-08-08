package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/auth"
	"github.com/stretchr/testify/assert"
)

func TestRegisterModelValidate(t *testing.T) {
	testCases := []struct {
		desc  string
		model auth.RegisterModel
		ok    bool
	}{
		{
			desc: "Expected valid model",
			model: auth.RegisterModel{
				Username: "test",
				Password: "test123",
			},
			ok: true,
		},
		{
			desc: "Model without username",
			model: auth.RegisterModel{
				Username: "",
				Password: "test123",
			},
			ok: false,
		},
		{
			desc: "Model without password",
			model: auth.RegisterModel{
				Username: "test",
				Password: "",
			},
			ok: false,
		},
		{
			desc: "Model without username and password",
			model: auth.RegisterModel{
				Username: "",
				Password: "",
			},
			ok: false,
		},
		{
			desc: "Model with username and password longer than 255 characters",
			model: auth.RegisterModel{
				Username: "test",
				Password: string(make([]byte, 256)),
			},
			ok: false,
		},
		{
			desc: "Model with username longer than 255 characters",
			model: auth.RegisterModel{
				Username: string(make([]byte, 256)),
				Password: "test123",
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
