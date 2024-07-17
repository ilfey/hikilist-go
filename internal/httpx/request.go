package httpx

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/ilfey/hikilist-go/internal/async"
	"github.com/ilfey/hikilist-go/internal/errorsx"
)

// RequestBuilder struct
type RequestBuilder struct {
	method string
	href   *url.URL
	body   io.Reader
	header map[string][]string

	client *http.Client
}

// Request is constructor for RequestBuilder.
func Request(method, href string) *RequestBuilder {
	return &RequestBuilder{
		method: method,
		href:   errorsx.Must(url.Parse(href)),
		header: map[string][]string{},
	}
}

/*
String returns string representation of request.

Returns string in format: "{METHOD} {URL}".
*/
func (r *RequestBuilder) String() string {
	return fmt.Sprintf(
		"%s %s",
		r.method,
		r.href.String(),
	)
}

// Method returns method of request.
func (r *RequestBuilder) Method() string {
	return r.method
}

// Url returns url of request.
func (r *RequestBuilder) Url() *url.URL {
	return r.href
}

// Body sets body of request.
func (r *RequestBuilder) Body(body io.Reader) *RequestBuilder {
	r.body = body

	return r
}

// Query adds query param.
func (r *RequestBuilder) Query(key, val string) *RequestBuilder {
	q := r.href.Query()
	q.Set(key, val)
	r.href.RawQuery = q.Encode()

	return r
}

// QueryMap adds query params from map.
func (r *RequestBuilder) QueryMap(values map[string]string) *RequestBuilder {
	query := r.href.Query()
	for k, v := range values {
		query.Set(k, v)
	}

	r.href.RawQuery = query.Encode()

	return r
}

// Header adds header.
func (r *RequestBuilder) Header(key, value string) *RequestBuilder {
	r.header[key] = append(r.header[key], value)

	return r
}

// HeaderMap adds headers from map.
func (r *RequestBuilder) HeaderMap(headers map[string]string) *RequestBuilder {
	for k, v := range headers {
		r.header[k] = []string{v}
	}

	return r
}

// Client sets client.
func (r *RequestBuilder) Client(client *http.Client) *RequestBuilder {
	r.client = client

	return r
}

// Build returns *http.Request.
func (r *RequestBuilder) Build() *http.Request {
	request := errorsx.Must(http.NewRequest(r.method, r.href.String(), r.body))

	request.Header = r.header

	return request
}

// Async returns promise with *Response.
func (r *RequestBuilder) Async() *async.Promise[*Response] {
	if r.client == nil {
		r.client = http.DefaultClient
	}

	return async.New(func() (*Response, error) {
		startTime := time.Now()

		res, err := r.client.Do(r.Build())
		if err != nil {
			return nil, err
		}

		response := Response{
			request:  r,
			response: res,
			duration: time.Since(startTime),
		}

		return &response, nil
	})
}

// Do sends request and returns *Response.
func (r *RequestBuilder) Do() (*Response, error) {
	if r.client == nil {
		r.client = http.DefaultClient
	}

	startTime := time.Now()

	res, err := r.client.Do(r.Build())
	if err != nil {
		return nil, err
	}

	response := Response{
		request:  r,
		response: res,
		duration: time.Since(startTime),
	}

	return &response, nil
}
