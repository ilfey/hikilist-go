package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/data/models/anime"
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
	"github.com/stretchr/testify/assert"
)

func TestPaginateNormalize(t *testing.T) {
	p := anime.Paginate{}

	p.Normalize()

	assert.Equal(t, 1, p.Page)
	assert.Equal(t, 10, p.Limit)
	assert.Equal(t, "-id", string(p.Order))
}

func TestPaginateValidate(t *testing.T) {
	p := anime.Paginate{}

	assert.NoError(t, p.Validate())

	// Test page

	p = anime.Paginate{
		Page: -1,
	}

	assert.Error(t, p.Validate())

	// Test limit

	p = anime.Paginate{
		Limit: -1,
	}

	assert.Error(t, p.Validate())

	// Test order

	awaibleOrders := []baseModels.OrderField{
		"",
		"id",
		"-id",
		"title",
		"-title",
		"episodes",
		"-episodes",
		"episodes_released",
		"-episodes_released",
	}

	for _, o := range awaibleOrders {
		p = anime.Paginate{
			Order: o,
		}

		assert.NoError(t, p.Validate())
	}

	notAwaibleOrders := []baseModels.OrderField{
		"test",
		"-test",
	}

	for _, o := range notAwaibleOrders {
		p = anime.Paginate{
			Order: o,
		}

		assert.Error(t, p.Validate())
	}
}
