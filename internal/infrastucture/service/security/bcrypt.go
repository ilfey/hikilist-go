package security

import (
	"github.com/ilfey/hikilist-go/internal/config/hasher"
	securityInterface "github.com/ilfey/hikilist-go/internal/domain/service/security/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct {
	log loggerInterface.Logger
	cfg *hasher.Config
}

func NewBcryptService(log loggerInterface.Logger, cfg *hasher.Config) *BcryptService {
	return &BcryptService{
		log: log,
		cfg: cfg,
	}
}

func (s *BcryptService) Hash(source string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(source),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", s.log.Propagate(err)
	}

	return string(hash), nil
}

func (s *BcryptService) Verify(provider securityInterface.HashProvider, source string) bool {
	return bcrypt.CompareHashAndPassword([]byte(provider.GetHash()), []byte(source)) == nil
}
