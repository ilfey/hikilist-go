package anime

import (
	"github.com/ilfey/hikilist-go/internal/paginate"
)

func NewPaginator(queries map[string][]string) *paginate.Paginator {
	return paginate.New(
		queries,
		paginate.WithDefaultOrder("-id"),
		paginate.WithDefaultLimit(10),
		paginate.WithMaxLimit(100),
		paginate.WithAwaiableOrders(
			[]string{
				"",
				"id",
				"-id",
				"title",
				"-title",
				"episodes",
				"-episodes",
				"episodes_released",
				"-episodes_released",
			},
		),
	)
}
