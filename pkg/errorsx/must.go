package errorsx

// Must returns `returnValue` and panics if `err` is not `nil`.
func Must[T any](returnValue T, err error) T {
	if err != nil {
		panic(err)
	}

	return returnValue
}
