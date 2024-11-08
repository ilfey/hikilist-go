package auth

import (
	"context"
	"github.com/ilfey/hikilist-database/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/enum"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	securityInterface "github.com/ilfey/hikilist-go/internal/domain/service/security/interface"
	tokenizerInterface "github.com/ilfey/hikilist-go/internal/domain/service/tokenizer/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strings"
)

type Auth struct {
	log loggerInterface.Logger

	hasher securityInterface.Hasher

	userRepo  repositoryInterface.User
	tokenizer tokenizerInterface.Tokenizer
}

func NewAuth(
	log loggerInterface.Logger,

	hasher securityInterface.Hasher,
	tokenizer tokenizerInterface.Tokenizer,

	user repositoryInterface.User,
) *Auth {
	return &Auth{
		log: log,

		hasher:    hasher,
		tokenizer: tokenizer,

		userRepo: user,
	}
}

func (s *Auth) IsAuthed(request *http.Request) (uint64, error) {
	// Detail token from header.
	header := request.Header.Get(enum.AccessTokenHeaderKey)
	if header == "" {
		return 0, s.log.Propagate(errtype.NewAuthFailedError("token not provided"))
	}

	// Check Bearer prefix.
	if !strings.HasPrefix(header, "Bearer ") {
		return 0, s.log.Propagate(errtype.NewAuthFailedError("invalid token prefix"))
	}

	// Remove Bearer prefix.
	token := strings.TrimPrefix(header, "Bearer ")

	// Verify token.
	userId, err := s.tokenizer.Verify(request.Context(), token)
	if err != nil {
		return 0, s.log.Propagate(err)
	}

	return userId, nil
}

func (s *Auth) ChangePassword(ctx context.Context, userId uint64, password string) error {
	hash, err := s.hasher.Hash(password)
	if err != nil {
		return s.log.Propagate(err)
	}

	err = s.userRepo.UpdatePassword(ctx, userId, hash)
	if err != nil {
		return s.log.Propagate(err)
	}

	return nil
}

func (s *Auth) DeleteUser(ctx context.Context, deleteDTO *dto.UserDeleteRequestDTO) error {
	user, err := s.userRepo.Get(ctx, map[string]any{
		"id": deleteDTO.UserID,
	})
	if err != nil {
		return err
	}

	if !s.hasher.Verify(user, deleteDTO.Password) {
		return s.log.Propagate(errtype.NewPasswordNotMatchError())
	}

	g := errgroup.Group{}

	// Delete refresh token
	g.Go(func() error {
		err := s.tokenizer.Block(ctx, deleteDTO.Refresh)
		if err != nil {
			return s.log.Propagate(err)
		}

		return nil
	})

	// Delete account
	g.Go(func() error {
		err := s.userRepo.Delete(ctx, deleteDTO.UserID)
		if err != nil {
			return s.log.Propagate(err)
		}

		return nil
	})

	return g.Wait()
}

func (s *Auth) Login(ctx context.Context, login *dto.AuthLoginRequestDTO) (*agg.TokenPair, error) {
	user, err := s.userRepo.Get(ctx, map[string]any{
		"username": login.Username,
	})
	if err != nil {
		return nil, s.log.Propagate(err)
	}

	if !s.hasher.Verify(user, login.Password) {
		return nil, s.log.Propagate(errtype.NewAuthFailedError("invalid credentials"))
	}

	tokensModel, err := s.tokenizer.Generate(user.ID)
	if err != nil {
		return nil, s.log.Propagate(err)
	}

	return tokensModel, nil
}

func (s *Auth) Register(ctx context.Context, registerModel *dto.AuthRegisterRequestDTO) (*dto.UserCreateRequestDTO, error) {
	hash, err := s.hasher.Hash(registerModel.Password)
	if err != nil {
		return nil, s.log.Propagate(err)
	}

	cm := dto.UserCreateRequestDTO{
		Username: registerModel.Username,
		Password: hash,
	}

	err = s.userRepo.Create(ctx, &cm)
	if err != nil {
		return nil, s.log.Propagate(err)
	}

	return &cm, nil
}

func (s *Auth) Refresh(ctx context.Context, refresh *dto.AuthRefreshRequestDTO) (*agg.TokenPair, error) {
	// Detail userId
	userId, err := s.tokenizer.Verify(ctx, refresh.Refresh)
	if err != nil {
		return nil, s.log.Propagate(err)
	}

	// Block old refresh token
	err = s.tokenizer.Block(ctx, refresh.Refresh)
	if err != nil {
		return nil, s.log.Propagate(err)
	}

	// Generate new tokens
	tokensModel, err := s.tokenizer.Generate(userId)
	if err != nil {
		return nil, s.log.Propagate(err)
	}

	return tokensModel, nil
}

func (s *Auth) Logout(ctx context.Context, logoutModel *dto.AuthLogoutRequestDTO) error {
	err := s.tokenizer.Block(ctx, logoutModel.Refresh)
	if err != nil {
		return s.log.Propagate(err)
	}

	return nil
}
