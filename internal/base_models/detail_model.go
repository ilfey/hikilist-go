package baseModels

import (
	"encoding/json"

	"github.com/ilfey/hikilist-go/internal/utils/errorsx"
	"github.com/ilfey/hikilist-go/internal/utils/resx"
)

type DetailModel struct{}

func (DetailModel) JSON(m interface{}) []byte {
	return errorsx.Must(json.Marshal(m))
}

func (DetailModel) Response(m interface {
	JSON() []byte
}) *resx.Response {
	return resx.NewResponse(200, m.JSON())
}
