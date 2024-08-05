package slicex

func Reduce[T any, R any](slice []T, reducer func(R, T) R, initialValue R) R {
	res := initialValue

	for _, item := range slice {
		res = reducer(res, item)
	}

	return res
}
