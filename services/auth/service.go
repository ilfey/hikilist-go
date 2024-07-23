package authService

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authConfig "github.com/ilfey/hikilist-go/config/auth"
	"github.com/ilfey/hikilist-go/data/entities"
	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/rotisserie/eris"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	GenerateTokens(ctx context.Context, dm *userModels.DetailModel) (*authModels.TokensModel, error)

	DeleteToken(ctx context.Context, token string) error

	// ParseToken парсит данные из токена авторизации.
	//
	// Возвращает данные токена и успех парсинга токена (true, если удалось распарсить и токен актуален).
	ParseToken(token string) (*Claims, error)

	// GetUser получение пользователя.
	//
	// Возвращает модель пользователя и транзакцию.
	GetUser(ctx context.Context, claims *Claims) (*userModels.DetailModel, error)

	// UpdateUserOnline обновляет время последней активности пользователя.
	//
	// Возвращает транзакцию
	UpdateUserOnline(ctx context.Context, dm *userModels.DetailModel) error
}

// Сервис аутентификации
type service struct {
	config *authConfig.Config
	db     *gorm.DB
}

// Конструктор сервиса аутентификации
func New(
	config *authConfig.Config,
	db *gorm.DB,
) Service {
	return &service{
		config: config,
		db:     db,
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

func (s *service) GenerateTokens(ctx context.Context, model *userModels.DetailModel) (*authModels.TokensModel, error) {
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
		UserID: model.ID,
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

	result := s.db.WithContext(ctx).
		Create(&entities.Token{
			Token: tokenSlice[len(tokenSlice)-1],
		})
	if result.Error != nil {
		return nil, eris.Wrap(result.Error, "failed create token")
	}

	return &authModels.TokensModel{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func (s *service) DeleteToken(ctx context.Context, token string) error {
	tokenSlice := strings.Split(token, ".")

	var entity entities.Token

	return s.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			// Get token from database
			result := tx.First(&entity, map[string]any{
				"Token": tokenSlice[len(tokenSlice)-1],
			})
			if result.Error != nil {
				if eris.Is(result.Error, gorm.ErrRecordNotFound) {
					return eris.Wrap(result.Error, "token not found")
				}

				return eris.Wrap(result.Error, "failed get token")
			}

			// Check if token is deleted already
			if entity.DeletedAt.Valid {
				return eris.New("token already deleted")
			}

			// Delete token
			result = tx.Delete(entity)
			if result.Error != nil {
				return eris.Wrap(result.Error, "failed delete token")
			}

			return nil
		})
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

func (s *service) GetUser(ctx context.Context, claims *Claims) (*userModels.DetailModel, error) {
	var dm userModels.DetailModel

	// Get user from database
	result := s.db.WithContext(ctx).
		Model(&entities.User{}).
		First(&dm, claims.UserID)
	if result.Error != nil {
		return nil, eris.Wrapf(result.Error, "failed get user with id: %d", claims.UserID)
	}

	return &dm, nil
}

func (s *service) UpdateUserOnline(ctx context.Context, user *userModels.DetailModel) error {
	currentTime := time.Now()

	user.LastOnline = &currentTime

	result := s.db.WithContext(ctx).
		Save(user.ToEntity())
	if result.Error != nil {
		return eris.Wrapf(result.Error, "failed update online for user with id: %d", user.ID)
	}

	return nil
}
