package extractorInterface

import "net/http"

type RequestParams interface {
	Parameters(r *http.Request) map[string]string
	HasParameter(req *http.Request, param string) bool
	GetParameter(req *http.Request, param string) (string, error)
}
