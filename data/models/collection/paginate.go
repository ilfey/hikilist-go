package collectionModels

import (
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
	"github.com/ilfey/hikilist-go/internal/validator"
)

type Paginate struct {
	baseModels.Paginate

	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewPaginateFromQuery(queries map[string][]string) *Paginate {
	p := Paginate{}

	p.Page = p.QueryInt(queries, "page")
	p.Limit = p.QueryInt(queries, "limit")

	return &p
}

func (p Paginate) Validate() validator.ValidateError {
	return validator.Validate(p, map[string][]validator.Option{
		"Page": {
			validator.GreaterThat[int64](-1),
		},
		"Limit": {
			validator.GreaterThat[int64](-1),
			validator.LessThat[int64](101),
		},
	})
}

func (p *Paginate) Normalize() *Paginate {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.Limit == 0 {
		p.Limit = 10
	}

	// Self-return
	return p
}
