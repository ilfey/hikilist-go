package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/pkg/models/auth"
	"github.com/stretchr/testify/assert"
)

func TestChangePasswordModelValidate(t *testing.T) {
	testCases := []struct {
		desc  string
		model auth.ChangePasswordModel
		ok    bool
	}{
		{
			desc: "Expected valid model",
			model: auth.ChangePasswordModel{
				OldPassword: "test12",
				NewPassword: "test123",
			},
			ok: true,
		},
		{
			desc: "Model without old password",
			model: auth.ChangePasswordModel{
				OldPassword: "",
				NewPassword: "test123",
			},
			ok: false,
		},
		{
			desc: "Model without new password",
			model: auth.ChangePasswordModel{
				OldPassword: "test12",
				NewPassword: "",
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
