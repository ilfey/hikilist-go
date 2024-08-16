package collection

import (
	"github.com/ilfey/hikilist-go/internal/paginate"
)

func NewPaginatorFromQuery(queries map[string][]string) *paginate.Paginator {
	return paginate.New(
		queries,
		paginate.WithDefaultLimit(10),
		paginate.WithMaxLimit(100),
		paginate.WithDefaultOrder("-id"),
	)
}
