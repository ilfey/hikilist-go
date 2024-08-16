package user

import (
	"github.com/ilfey/hikilist-go/internal/paginate"
)

func NewPaginator(queries map[string][]string) *paginate.Paginator {
	return paginate.New(
		queries,
		paginate.WithAwaiableOrders(
			[]string{
				"",
				"id",
				"-id",
				"username",
				"-username",
			},
		),
		paginate.WithDefaultLimit(10),
		paginate.WithMaxLimit(101),
		paginate.WithDefaultOrder("-id"),
	)
}
