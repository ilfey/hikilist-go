package builderInterface

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"net/http"
)

type Action interface {
	BuildListRequestDTOFromRequest(r *http.Request) (*dto.ActionListRequestDTO, error)
}
