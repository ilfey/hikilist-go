package userAction

import (
	baseModels "github.com/ilfey/hikilist-go/internal/base_models"
	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
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

func (p Paginate) Validate() error {
	return validator.Validate(p, map[string][]options.Option{
		"Page": {
			options.GreaterThan[int64](-1),
		},
		"Limit": {
			options.GreaterThan[int64](-1),
			options.LessThan[int64](101),
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
