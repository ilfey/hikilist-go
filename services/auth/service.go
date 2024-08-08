package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang-jwt/jwt/v5"
	authConfig "github.com/ilfey/hikilist-go/config/auth"
	"github.com/ilfey/hikilist-go/data/database"
	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	tokenModels "github.com/ilfey/hikilist-go/data/models/token"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	userActionModels "github.com/ilfey/hikilist-go/data/models/userAction"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/postgres"
	"github.com/rotisserie/eris"
	"golang.org/x/crypto/bcrypt"
)

// Сервис авторизации.
//
// Предоставляет функционал аунтентификации пользователей.
type Service interface {
	// ComparePassword сравнивает пароль пользователя.
	//
	// Возвращает true, если пароли совпадают.
	CompareUserPassword(dm *userModels.DetailModel, password string) bool

	// GenerateTokens создает модель токенов аунтентификации пользователя.
	//
	// Возвращает модель токенов.
	GenerateTokens(ctx context.Context, userId uint) (*authModels.TokensModel, error)

	DeleteToken(ctx context.Context, token string) error

	Logout(ctx context.Context, model *authModels.RefreshModel) error

	// ParseToken парсит данные из токена авторизации.
	//
	// Возвращает данные токена и успех парсинга токена (true, если удалось распарсить и токен актуален).
	ParseToken(token string) (*Claims, error)

	// Authorize получение пользователя.
	//
	// Возвращает модель пользователя и транзакцию.
	Authorize(ctx context.Context, claims *Claims) (*userModels.DetailModel, error)

	Register(ctx context.Context, registerModel *authModels.RegisterModel) (*userModels.CreateModel, error)

	// UpdateUserOnline обновляет время последней активности пользователя.
	//
	// Возвращает транзакцию
	UpdateUserOnline(ctx context.Context, dm *userModels.DetailModel) error
}

// Сервис аутентификации
type service struct {
	config *authConfig.Config
}

// Конструктор сервиса аутентификации
func New(
	config *authConfig.Config,
) Service {
	return &service{
		config: config,
	}
}

func (s *service) CompareUserPassword(model *userModels.DetailModel, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(password)) == nil
}

// Данные, хранящиеся в токене
type Claims struct {
	jwt.RegisteredClaims
	UserID uint `json:"user_id"`
}

func (s *service) GenerateTokens(ctx context.Context, userId uint) (*authModels.TokensModel, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: s.config.Issuer,
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),
			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
		UserID: userId,
	}

	// Создание 24 часового токена
	accessToken := s.generateToken(claims)

	claims.ExpiresAt = jwt.NewNumericDate(
		time.Now().Add(7 * 24 * time.Hour),
	)

	// Создание 7 дневного токена
	refreshToken := s.generateToken(claims)

	// Save token in database
	tokenSlice := strings.Split(refreshToken, ".")

	cm := &tokenModels.CreateModel{
		Token: tokenSlice[len(tokenSlice)-1],
	}

	err := cm.Insert(ctx)
	if err != nil {
		return nil, err
	}

	return &authModels.TokensModel{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func (s *service) DeleteToken(ctx context.Context, token string) error {
	tokenSlice := strings.Split(token, ".")

	sql, args, err := sq.Delete("tokens").
		Where(sq.Eq{"token": tokenSlice[len(tokenSlice)-1]}).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return eris.Wrap(err, "failed to build delete query")
	}

	var id uint

	err = database.Instance().QueryRow(
		ctx,
		sql,
		args...,
	).Scan(&id)
	if err != nil {
		return eris.Wrap(err, "failed to delete token")
	}

	return nil
}

func (s *service) Logout(ctx context.Context, model *authModels.RefreshModel) error {
	err := model.Validate()
	if err != nil {
		return eris.Wrap(err, "failed to validate model")
	}

	_, err = s.ParseToken(model.Refresh)
	if err != nil {
		return eris.Wrap(err, "failed to parse token")
	}

	return s.DeleteToken(ctx, model.Refresh)
}

func (s *service) generateToken(claims *Claims) string {
	return errorsx.Must(
		jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			claims,
		).SignedString(s.config.Secret),
	)
}

// Парсинг токена
func (s *service) ParseToken(token string) (*Claims, error) {
	claims := new(Claims)
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if iss, err := token.Claims.GetIssuer(); iss != s.config.Issuer || err != nil {
			return nil, errors.New("unexpected issuer")
		}

		return s.config.Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (s *service) Authorize(ctx context.Context, claims *Claims) (*userModels.DetailModel, error) {
	var dm userModels.DetailModel

	// Get user from database
	err := dm.Get(ctx, map[string]any{
		"id": claims.UserID,
	})
	if err != nil {
		return nil, eris.Wrapf(err, "failed get user with id: %d", claims.UserID)
	}

	return &dm, nil
}

func (s *service) passwordToHash(password string) string {
	hash := errorsx.Must(
		bcrypt.GenerateFromPassword(
			[]byte(password),
			bcrypt.DefaultCost,
		),
	)

	return string(hash)
}

func (s *service) Register(ctx context.Context, registerModel *authModels.RegisterModel) (*userModels.CreateModel, error) {
	err := registerModel.Validate()
	if err != nil {
		return nil, eris.Wrap(err, "failed to validate model")
	}

	cm := userModels.CreateModel{
		Username: registerModel.Username,
		Password: s.passwordToHash(registerModel.Password),
	}

	err = database.Instance().RunTx(ctx, func(tx *postgres.Transaction) error {
		// Create user
		sql, args, err := cm.InsertSQL()
		if err != nil {
			return err
		}

		err = tx.QueryRow(ctx, sql, args...).Scan(&cm.ID)
		if err != nil {
			return eris.Wrap(err, "failed to insert user")
		}

		// Create user action
		actionCm := userActionModels.NewRegisterUserAction(cm.ID)

		sql, args, err = actionCm.InsertSQL()
		if err != nil {
			return err
		}

		err = tx.QueryRow(ctx, sql, args...).Scan(&cm.ID)
		if err != nil {
			return eris.Wrap(err, "failed to insert user action")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &cm, nil
}

func (s *service) UpdateUserOnline(ctx context.Context, user *userModels.DetailModel) error {
	err := user.UpdateLastOnline(ctx)
	if err != nil {
		return eris.Wrapf(err, "failed update online for user with id: %d", user.ID)
	}

	return nil
}

func (s *service) UpdateUserPassword(ctx context.Context, userId uint, password string) error {
	hash := s.passwordToHash(password)

	dm := userModels.DetailModel{
		ID:       userId,
		Password: hash,
	}

	sql, args, err := dm.UpdatePasswordSQL()
	if err != nil {
		return err
	}

	_, err = database.Instance().Exec(ctx, sql, args...)
	if err != nil {
		return eris.Wrapf(err, "failed update password for user with id: %d", userId)
	}

	return nil
}
