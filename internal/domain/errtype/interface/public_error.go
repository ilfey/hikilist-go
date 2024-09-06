package errtypeInterface

type PublicError interface {
	error
	Status() int
	Type() string
	Public() bool
}
