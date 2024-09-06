package extractor

import (
	"github.com/gorilla/mux"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	"net/http"
)

type ParametersExtractor struct{}

func NewParametersExtractor() *ParametersExtractor {
	return &ParametersExtractor{}
}

// Parameters - will return a map with param. names as keys to values
func (e *ParametersExtractor) Parameters(r *http.Request) map[string]string {
	params := make(map[string]string)

	for name, param := range r.URL.Query() {
		params[name] = param[0]
	}

	for name, param := range mux.Vars(r) {
		params[name] = param
	}

	return params
}

// HasParameter - checking the param. is existing in request
func (e *ParametersExtractor) HasParameter(req *http.Request, param string) bool {
	if param == "" {
		return false
	}

	if _, ok := e.Parameters(req)[param]; ok {
		return true
	}

	return false
}

// GetParameter - checking if the param. is existing in request, it will be returned
func (e *ParametersExtractor) GetParameter(req *http.Request, param string) (string, error) {
	if param == "" {
		return "", errtype.NewFieldCannotBeEmptyError(param)
	}

	v, ok := e.Parameters(req)[param]
	if !ok {
		return "", errtype.NewFieldCannotBeEmptyError(param)
	}

	return v, nil
}
