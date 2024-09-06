package security

import (
	"github.com/ilfey/hikilist-go/internal/config/hasher"
	securityInterface "github.com/ilfey/hikilist-go/internal/domain/service/security/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct {
	logger loggerInterface.Logger
	config *hasher.Config
}

func NewBcryptService(logger loggerInterface.Logger, config *hasher.Config) *BcryptService {
	return &BcryptService{
		logger: logger,
		config: config,
	}
}

func (s *BcryptService) Hash(source string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(source),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", s.logger.LogPropagate(err)
	}

	return string(hash), nil
}

func (s *BcryptService) Verify(provider securityInterface.HashProvider, source string) bool {
	return bcrypt.CompareHashAndPassword([]byte(provider.GetHash()), []byte(source)) == nil
}
