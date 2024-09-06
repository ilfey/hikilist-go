package errorsx

// Ignore returns `returnValue` and ignores `err`.
func Ignore[T any](returnValue T, _ error) T {
	return returnValue
}
