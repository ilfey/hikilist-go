package errorsx

func Must[T any](returnValue T, err error) T {
	if err != nil {
		panic(err)
	}

	return returnValue
}
