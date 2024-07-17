package baseModels

import (
	"encoding/json"

	"github.com/ilfey/hikilist-go/internal/errorsx"
)

type DetailModel struct{}

func (DetailModel) ToJSON(m interface{}) []byte {
	return errorsx.Must(json.Marshal(m))
}
