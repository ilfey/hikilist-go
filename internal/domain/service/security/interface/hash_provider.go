package securityInterface

type HashProvider interface {
	GetHash() string
}
