package errorsx

// Ignore returns `returnValue` and ignores `err`.
func Ignore[T any](returnValue T, err error) T {
	return returnValue
}
