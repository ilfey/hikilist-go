package requestbuilder

import (
	"github.com/ilfey/hikilist-go/internal/async"
	"github.com/ilfey/hikilist-go/internal/httpx"
)

type Request struct {
	*httpx.RequestBuilder

	requestHooks  []func(*httpx.RequestBuilder)
	responseHooks []func(*httpx.Response)
}

func (r *Request) execRequestHooks(req *httpx.RequestBuilder) {
	for _, hook := range r.requestHooks {
		hook(req)
	}
}

func (r *Request) execResponseHooks(res *httpx.Response) {
	for _, hook := range r.responseHooks {
		hook(res)
	}
}

func (r *Request) Do() (*httpx.Response, error) {

	r.execRequestHooks(r.RequestBuilder)

	res, err := r.RequestBuilder.Do()
	if err != nil {
		return nil, err
	}

	r.execResponseHooks(res)

	return res, nil
}

func (r *Request) Async() *async.Promise[*httpx.Response] {
	r.execRequestHooks(r.RequestBuilder)

	promise := r.RequestBuilder.Async()

	promise.Then(func(res *httpx.Response) {
		r.execResponseHooks(res)
	})

	return promise
}
