package action

import (
	"github.com/ilfey/hikilist-go/internal/paginate"
)

func NewPaginator(queries map[string][]string) *paginate.Paginator {
	return paginate.New(
		queries,
		paginate.WithDefaultLimit(10),
		paginate.WithMaxLimit(100),
	)
}
