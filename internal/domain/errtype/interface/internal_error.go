package errtypeInterface

type InternalError interface {
	error

	Internal() bool
}
