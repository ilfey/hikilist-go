package baseModels

import (
	"encoding/json"

	"github.com/ilfey/hikilist-go/internal/utils/errorsx"
	"github.com/ilfey/hikilist-go/internal/utils/resx"
)

type ListModel[T any] struct {
	Results []*T `json:"results"`

	Count int64 `json:"count"`
}

func (m *ListModel[T]) JSON() []byte {
	return errorsx.Must(json.Marshal(m))
}

func (m *ListModel[T]) Response() *resx.Response {
	return resx.NewResponse(200, m.JSON())
}
