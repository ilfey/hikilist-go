package builderInterface

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"net/http"
)

type Pagination interface {
	BuildPaginationRequestDROFromRequest(r *http.Request) (*dto.PaginationRequestDTO, error)
}
