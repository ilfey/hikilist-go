package resx

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/utils/errorsx"
)

type J = map[string]any

type Response struct {
	code int
	body []byte
}

func NewResponse(code int, body []byte) *Response {
	return &Response{code, body}
}

func (r *Response) JSON(w http.ResponseWriter) (int, error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.code)
	return w.Write(r.body)
}

func JSON(code int, body map[string]any) *Response {
	return NewResponse(
		code,
		[]byte(errorsx.Must(json.Marshal(body))),
	)
}
