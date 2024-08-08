package async

import "sync"

// Promise
type Promise[T any] struct {
	value chan T // Value
	cache *T     // Cached value

	err error // Error

	onthens   []func(T)     // On then functions
	oncatches []func(error) // On catch functions

	mu sync.Mutex
}

/*
New is a promise constructor.

`fn` is a function that will be called as goroutine.

Returns promise.
*/
func New[T any](fn func() (T, error)) *Promise[T] {
	promise := Promise[T]{
		value: make(chan T),
		mu:    sync.Mutex{},
	}

	go func() {
		value, err := fn()

		// Call hooks

		promise.mu.Lock()

		if err == nil {
			for _, onthen := range promise.onthens {
				onthen(value)
			}

			promise.onthens = []func(T){}
		}

		if err != nil {
			for _, oncatch := range promise.oncatches {
				oncatch(err)
			}

			promise.oncatches = []func(error){}
		}

		promise.mu.Unlock()

		promise.err = err

		promise.cache = &value

		promise.value <- value

		close(promise.value)
	}()

	return &promise
}

/*
Await is a func that waits until the promise is resolved.

Returns value and error.
*/
func (p *Promise[T]) Await() (T, error) {
	if p.cache != nil {
		return *p.cache, p.err
	}

	return p.Wait(), p.err
}

/*
Wait is a func that waits until the promise is resolved.

Returns value.
*/
func (p *Promise[T]) Wait() T {
	if p.cache != nil {
		return *p.cache
	}

	return <-p.value
}

/*
Then is a func that calls when the promise is resolved.

Returns promise.
*/
func (p *Promise[T]) Then(callback func(T)) *Promise[T] {
	if p.cache == nil && p.err == nil {
		p.mu.Lock()

		p.onthens = append(p.onthens, callback)

		p.mu.Unlock()

		return p
	}

	if p.cache != nil && p.err == nil {
		callback(*p.cache)
	}

	return p
}

/*
Then is a func that calls when the promise is rejected.

Returns promise.
*/
func (p *Promise[T]) Catch(callback func(error)) *Promise[T] {
	if p.cache == nil && p.err == nil {
		p.mu.Lock()

		p.oncatches = append(p.oncatches, callback)

		p.mu.Unlock()

		return p
	}

	if p.err != nil {
		callback(p.err)
	}

	return p
}
