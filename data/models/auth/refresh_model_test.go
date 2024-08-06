package authModels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefreshModelValidate(t *testing.T) {
	invalid := RefreshModel{}

	assert.Error(t, invalid.Validate())

	valid := RefreshModel{
		Refresh: "test",
	}

	assert.NoError(t, valid.Validate())
}
