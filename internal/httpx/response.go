package httpx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ilfey/hikilist-go/internal/errorsx"
)

type Response struct {
	request  *RequestBuilder
	response *http.Response

	duration time.Duration
}

// Response returns *http.Response.
func (r *Response) Response() *http.Response {
	return r.response
}

/* 
String returns string representation of response.

Returns string in format: "{METHOD} {URL} {STATUS} {DURATION}ms".
*/
func (r *Response) String() string {
	return fmt.Sprintf(
		"%s %s %d %dms",
		r.request.Method(),
		r.request.Url(),
		r.StatusCode(),
		r.Duration().Milliseconds(),
	)
}

// Duration returns duration of response.
func (r *Response) Duration() time.Duration {
	return r.duration
}

// StatusCode returns status code of response.
func (r *Response) StatusCode() int {
	return r.response.StatusCode
}

// Body returns body of response.
func (r *Response) Body() io.ReadCloser {
	return r.Response().Body
}

// Text returns text of response.
func (r *Response) Text() string {
	return string(errorsx.Ignore(io.ReadAll(r.Body())))
}

// JSON returns json of response.
func (r *Response) JSON(response any) error {
	return json.NewDecoder(r.Body()).Decode(response)
}
