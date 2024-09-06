package securityInterface

type Hasher interface {
	Hash(source string) (string, error)
	Verify(provider HashProvider, source string) bool
}
