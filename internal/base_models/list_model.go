package baseModels

import (
	"encoding/json"

	"github.com/ilfey/hikilist-go/internal/errorsx"
)

type ListModel[T any] struct {
	Results []*T `json:"results"`

	Count int64 `json:"count"`
}

func (m *ListModel[T]) ToJSON() []byte {
	return errorsx.Must(json.Marshal(m))
}
