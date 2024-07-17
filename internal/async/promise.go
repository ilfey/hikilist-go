package async

// Promise
type Promise[T any] struct {
	value chan T // Value
	cache *T     // Cached value

	err error // Error

	onthen  func(T)     // On then function
	oncatch func(error) // On catch function
}

/*
New is a promise constructor.

`fn` is a function that will be called as goroutine.

Returns promise.
*/
func New[T any](fn func() (T, error)) *Promise[T] {
	promise := Promise[T]{
		value: make(chan T),
	}

	go func() {
		value, err := fn()

		promise.err = err

		promise.value <- value

		close(promise.value)

		// Call hooks
		if promise.onthen != nil && err == nil {
			promise.onthen(value)
		}

		if promise.oncatch != nil && err != nil {
			promise.oncatch(err)
		}
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

	value := <-p.value

	p.cache = &value

	return value, p.err
}

/*
Then is a func that calls when the promise is resolved.
This function is blocked until the promise is resolved.

Returns promise.
*/
func (p *Promise[T]) Then(callback func(T)) *Promise[T] {
	p.Await()

	if p.err == nil {
		callback(*p.cache)
	}

	return p
}

/*
Then is a func that calls when the promise is rejected.
This function is blocked until the promise is resolved.

Returns promise.
*/
func (p *Promise[T]) Catch(callback func(error)) *Promise[T] {
	p.Await()

	if p.err != nil {
		callback(p.err)
	}

	return p
}
