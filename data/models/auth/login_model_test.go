package authModels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginModelValidate(t *testing.T) {
	invalids := []LoginModel{
		{
			Username: "test",
			Password: "test",
		},
		{
			Username: "test",
			Password: "very_very_very_very_very_long_password",
		},
		{
			Username: "test",
			Password: "",
		},
		{
			Username: "",
			Password: "test",
		},
		{
			Username: "",
			Password: "",
		},
	}

	for _, m := range invalids {
		assert.Error(t, m.Validate())
	}

	valids := []LoginModel{
		{
			Username: "test",
			Password: "test123",
		},
		{
			Username: "test123",
			Password: "long_password",
		},
	}

	for _, m := range valids {
		assert.NoError(t, m.Validate())
	}
}
