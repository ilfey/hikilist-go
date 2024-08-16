package paginate

import (
	"strconv"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
)

type Paginator struct {
	Opts *Options `json:"-"`

	Page  int        `json:"page"`
	Limit int        `json:"limit"`
	Order OrderField `json:"order"`
}

func New(queries map[string][]string, options ...Option) *Paginator {
	p := Paginator{}

	opts := DefaultOptions

	for _, opt := range options {
		opt(&opts)
	}

	p.Opts = &opts

	p.Page = p.QueryInt(queries, "page")
	p.Limit = p.QueryInt(queries, "limit")
	p.Order = p.QueryOrder(queries, "order")

	// Normalize
	if p.Order == "" {
		p.Order = OrderField(p.Opts.DefaultOrder)
	}

	if p.Page == 0 {
		p.Page = 1
	}

	if p.Limit == 0 {
		p.Limit = p.Opts.DefaultLimit
	}

	return &p
}

func (p *Paginator) Validate() error {
	return validator.Validate(p, map[string][]options.Option{
		"Page": {
			options.GreaterThan[int64](-1),
		},
		"Limit": {
			options.GreaterThan[int64](-1),
			options.LessThan(p.Opts.MaxLimit + 1),
		},
		"Order": {
			options.InList(p.Opts.AwaiableOrders),
		},
	})
}

func (Paginator) QueryInt(q map[string][]string, key string) int {
	valueStrings, ok := q[key]
	if !ok {
		return 0
	}

	valueString := valueStrings[0]

	value, err := strconv.Atoi(valueString)
	if err != nil {
		return 0
	}

	return value
}

func (Paginator) QueryOrder(q map[string][]string, key string) OrderField {
	valueStrings, ok := q[key]
	if !ok {
		return OrderField("")
	}

	return OrderField(valueStrings[0])
}

func (p *Paginator) GetOffset(page, limit int) int {
	return (page - 1) * limit
}
