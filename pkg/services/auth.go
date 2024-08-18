package services

import (
	"context"
	"strings"
	"time"

	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/sirupsen/logrus"

	"github.com/ilfey/hikilist-go/pkg/claims"
	config "github.com/ilfey/hikilist-go/pkg/config/auth"
	"github.com/ilfey/hikilist-go/pkg/repositories"

	"github.com/ilfey/hikilist-go/pkg/models/auth"
	"github.com/ilfey/hikilist-go/pkg/models/token"
	"github.com/ilfey/hikilist-go/pkg/models/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rotisserie/eris"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
)

type Auth interface {
	CompareUserPassword(dm *user.DetailModel, password string) bool
	ChangePassword(ctx context.Context, userId uint, password string) error
	DeleteUser(ctx context.Context, userId uint, refresh string) error

	// DeleteToken(ctx context.Context, token string) error
	RefreshTokens(ctx context.Context, refresh string) (*auth.TokensModel, error)
	GenerateTokens(ctx context.Context, userId uint) (*auth.TokensModel, error)
	Logout(ctx context.Context, model *auth.LogoutModel) error
	ParseToken(token string) (*claims.Claims, error)
	Register(ctx context.Context, registerModel *auth.RegisterModel) (*user.CreateModel, error)
}

type AuthImpl struct {
	logger logrus.FieldLogger

	config    *config.Config
	userRepo  repositories.User
	tokenRepo repositories.Token
}

func NewAuth(
	logger logrus.FieldLogger,

	config *config.Config,
	userRepo repositories.User,
	tokenRepo repositories.Token,
) Auth {
	return &AuthImpl{
		logger: logger,

		config:    config,
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

/* ****** USERS ****** */

func (s *AuthImpl) CompareUserPassword(model *user.DetailModel, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(password)) == nil
}

func (s *AuthImpl) ChangePassword(ctx context.Context, userId uint, password string) error {
	hash := s.passwordToHash(password)

	return s.userRepo.UpdatePassword(ctx, userId, hash)
}

func (s *AuthImpl) DeleteUser(ctx context.Context, userId uint, refresh string) error {
	g := errgroup.Group{}

	// Delete refresh token
	g.Go(func() error {
		return s.DeleteToken(ctx, refresh)
	})

	// Delete account
	g.Go(func() error {
		err := s.userRepo.Delete(ctx, userId)
		if err != nil {
			s.logger.Debugf("Error occurred while deleting user %v", err)

			return err
		}

		return nil
	})

	return g.Wait()
}

/* ****** TOKENS ****** */

func (s *AuthImpl) Register(ctx context.Context, registerModel *auth.RegisterModel) (*user.CreateModel, error) {
	cm := user.CreateModel{
		Username: registerModel.Username,
		Password: s.passwordToHash(registerModel.Password),
	}

	err := s.userRepo.Create(ctx, &cm)
	if err != nil {
		return nil, err
	}

	return &cm, nil
}

func (s *AuthImpl) RefreshTokens(ctx context.Context, refresh string) (*auth.TokensModel, error) {
	claims, err := s.ParseToken(refresh)
	if err != nil {
		return nil, err
	}

	tm := auth.TokensModel{
		Access:  s.generateAccess(claims.UserID),
		Refresh: s.generateRefresh(claims.UserID),
	}

	g := errgroup.Group{}

	// Delete old refresh token
	g.Go(func() error {
		return s.DeleteToken(ctx, refresh)
	})

	// Save new refresh token
	g.Go(func() error {
		return s.saveRefresh(ctx, tm.Refresh)
	})

	err = g.Wait()
	if err != nil {
		return nil, err
	}

	return &tm, nil
}

func (s *AuthImpl) Logout(ctx context.Context, logoutModel *auth.LogoutModel) error {
	_, err := s.ParseToken(logoutModel.Refresh)
	if err != nil {
		s.logger.Debugf("Error occurred on token parsing %v", err)

		return eris.Wrap(err, "failed to parse token")
	}

	return s.DeleteToken(ctx, logoutModel.Refresh)
}

func (s *AuthImpl) ParseToken(token string) (*claims.Claims, error) {
	claims := new(claims.Claims)
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if iss, err := token.Claims.GetIssuer(); iss != s.config.Issuer || err != nil {
			return nil, eris.New("unexpected issuer")
		}

		return s.config.Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, eris.New("invalid token")
	}

	return claims, nil
}

func (s *AuthImpl) GenerateTokens(ctx context.Context, userId uint) (*auth.TokensModel, error) {
	// Generate tokens
	tm := auth.TokensModel{
		Access:  s.generateAccess(userId),
		Refresh: s.generateRefresh(userId),
	}

	// Save refresh
	err := s.saveRefresh(ctx, tm.Refresh)
	if err != nil {
		return nil, err
	}

	return &tm, nil
}

func (s *AuthImpl) DeleteToken(ctx context.Context, token string) error {
	err := s.tokenRepo.Delete(ctx, map[string]any{
		"token": s.getTokenDBView(token),
	})
	if err != nil {
		s.logger.Debugf("Error occurred while deleting token %v", err)

		return err
	}

	return nil
}

/* ****** UTILS ****** */

func (s *AuthImpl) generateAccess(userId uint) string {
	now := time.Now()

	claims := claims.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: s.config.Issuer,
			ExpiresAt: jwt.NewNumericDate(
				now.Add(time.Duration(s.config.AccessLifeTime) * time.Hour),
			),
			IssuedAt: jwt.NewNumericDate(
				now,
			),
		},
		UserID: userId,
	}

	return s.generate(&claims)
}

func (s *AuthImpl) generateRefresh(userId uint) string {
	now := time.Now()

	claims := claims.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: s.config.Issuer,
			ExpiresAt: jwt.NewNumericDate(
				now.Add(time.Duration(s.config.RefreshLifeTime) * time.Hour),
			),
			IssuedAt: jwt.NewNumericDate(
				now,
			),
		},
		UserID: userId,
	}

	return s.generate(&claims)
}

func (s *AuthImpl) saveRefresh(ctx context.Context, refresh string) error {
	cm := token.CreateModel{
		Token: s.getTokenDBView(refresh),
	}

	err := s.tokenRepo.Create(ctx, &cm)
	if err != nil {
		s.logger.Debugf("Error occured while creating token %v", err)

		return err
	}

	return nil
}

func (s *AuthImpl) generate(claims *claims.Claims) string {
	return errorsx.Must(
		jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			claims,
		).SignedString(s.config.Secret),
	)
}

func (s *AuthImpl) passwordToHash(password string) string {
	hash := errorsx.Must(
		bcrypt.GenerateFromPassword(
			[]byte(password),
			bcrypt.DefaultCost,
		),
	)

	return string(hash)
}

func (s *AuthImpl) getTokenDBView(token string) string {
	tokenSlice := strings.Split(token, ".")

	return tokenSlice[len(tokenSlice)-1]
}
