package extractor

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/rotisserie/eris"
)

var (
	ErrPageEmpty = eris.New("page is empty")

	ErrExtractorStopped  = eris.New("extractor was cancelled")
	ErrExtractorsStopped = eris.New("extractors was cancelled")
	ErrParserStopped     = eris.New("parsing was cancelled")
)

type NetworkError struct {
	StatusCode int
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf(
		"network error %d %s",
		e.StatusCode,
		strings.ToLower(http.StatusText(e.StatusCode)),
	)
}
