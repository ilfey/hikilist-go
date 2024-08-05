package slicex

func Map[T any, R any](slice []T, mapper func(T) R) []R {
	res := make([]R, 0, len(slice))

	for _, item := range slice {
		res = append(res, mapper(item))
	}

	return res
}
