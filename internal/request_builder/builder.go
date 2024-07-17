package requestbuilder

import (
	"net/http"
	"strings"

	"github.com/ilfey/hikilist-go/internal/httpx"
)

type RequestBuilder struct {
	baseUrl string
	client  *http.Client

	requestHooks  []func(*httpx.RequestBuilder)
	responseHooks []func(*httpx.Response)
}

// NewRequestBuilder constructor
func NewRequestBuilder(baseUrl string, client *http.Client) *RequestBuilder {
	return &RequestBuilder{
		baseUrl: baseUrl,
		client:  client,
	}
}

/*
AddRequestHook adds request hook.

This hook will be called before request is sent.

Returns builder.
*/
func (b *RequestBuilder) AddRequestHook(hook func(*httpx.RequestBuilder)) {
	b.requestHooks = append(b.requestHooks, hook)
}

/*
AddResponseHook adds response hook.

This hook will be called after request is sent.

Returns builder.
*/
func (b *RequestBuilder) AddResponseHook(hook func(*httpx.Response)) {
	b.responseHooks = append(b.responseHooks, hook)
}

// Get returns GET request.
func (b *RequestBuilder) Get(path string) *Request {
	return b.wrap(
		httpx.Request(http.MethodGet, b.buildUrl(path)).
			Client(b.client),
	)
}

// Sub returns new sub request builder.
func (b *RequestBuilder) Sub(path string) *RequestBuilder {
	return &RequestBuilder{
		baseUrl: b.buildUrl(path),
		client:  b.client,

		requestHooks:  b.requestHooks,
		responseHooks: b.responseHooks,
	}
}

func (b *RequestBuilder) wrap(r *httpx.RequestBuilder) *Request {
	return &Request{
		RequestBuilder: r,

		requestHooks:  b.requestHooks,
		responseHooks: b.responseHooks,
	}
}

func (b *RequestBuilder) buildUrl(path string) string {
	if b.baseUrl == "" || b.baseUrl == "/" {
		return b.baseUrl
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return strings.TrimSuffix(b.baseUrl, "/") + path
}
