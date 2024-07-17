package authService

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authConfig "github.com/ilfey/hikilist-go/config/auth"
	"github.com/ilfey/hikilist-go/data/entities"
	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	tokenRepository "github.com/ilfey/hikilist-go/repositories/token"
	userRepository "github.com/ilfey/hikilist-go/repositories/user"
	"golang.org/x/crypto/bcrypt"
)

// Сервис авторизации.
//
// Предоставляет функционал аунтентификации пользователей.
type Service interface {
	// ComparePassword сравнивает пароль пользователя.
	//
	// Возвращает true, если пароли совпадают.
	CompareUserPassword(*userModels.DetailModel, string) bool

	// GenerateTokens создает модель токенов аунтентификации пользователя.
	//
	// Возвращает модель токенов.
	GenerateTokens(*userModels.DetailModel) (*authModels.TokensModel, error)

	DeleteToken(token string) error

	// ParseToken парсит данные из токена авторизации.
	//
	// Возвращает данные токена и успех парсинга токена (true, если удалось распарсить и токен актуален).
	ParseToken(token string) (*Claims, error)

	// GetUser получение пользователя.
	//
	// Возвращает модель пользователя и транзакцию.
	GetUser(*Claims) (*userModels.DetailModel, error)

	// UpdateUserOnline обновляет время последней активности пользователя.
	//
	// Возвращает транзакцию
	UpdateUserOnline(*userModels.DetailModel) error
}

// Сервис аутентификации
type service struct {
	config *authConfig.Config
	*Dependencies
}

// Зависимости сервиса
type Dependencies struct {
	User  userRepository.Repository
	Token tokenRepository.Repository
}

// Конструктор сервиса аутентификации
func NewService(config *authConfig.Config, deps *Dependencies) Service {
	return &service{
		config:       config,
		Dependencies: deps,
	}
}

// Create Token
// func (s *AuthService) Create(model *authModels.TokenCreateModel) (*authModels.TokenDetailModel, *gorm.DB) {
// 	userEntity := &entities.User{}
// 	userEntity.ID = model.User.ID

// 	// Create entity
// 	tokenEntity := &entities.Token{
// 		Token: model.Token,
// 		User:  userEntity,
// 	}

// 	tx := s.repository.Create(tokenEntity)

// 	detailModel := authModels.TokenDetailModelFromEntity(tokenEntity)

// 	return detailModel, tx
// }

func (s *service) CompareUserPassword(model *userModels.DetailModel, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(password)) == nil
}

// GetByID Token
// func (s *AuthService) GetByID(id uint64) (*authModels.TokenDetailModel, *gorm.DB) {
// 	entity, tx := s.repository.Get(map[string]any{
// 		"ID": id,
// 	})

// 	model := authModels.TokenDetailModelFromEntity(entity)

// 	return model, tx
// }

// GetByID Token
// func (s *AuthService) GetByToken(token string) (*authModels.TokenDetailModel, *gorm.DB) {
// 	entity, tx := s.repository.Get(map[string]any{
// 		"Token": token,
// 	})

// 	model := authModels.TokenDetailModelFromEntity(entity)

// 	return model, tx
// }

// Данные, хранящиеся в токене
type Claims struct {
	jwt.RegisteredClaims
	UserID uint `json:"user_id"`
}

func (s *service) GenerateTokens(model *userModels.DetailModel) (*authModels.TokensModel, error) {
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

	// Сохранение в БД
	tokenSlice := strings.Split(refreshToken, ".")

	tx := s.Token.Create(&entities.Token{
		Token: tokenSlice[len(tokenSlice)-1],
	})
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &authModels.TokensModel{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func (s *service) DeleteToken(token string) error {
	tokenSlice := strings.Split(token, ".")

	// Get token
	entity, tx := s.Token.Get(map[string]any{
		"Token": tokenSlice[len(tokenSlice)-1],
	})
	if tx.Error != nil {
		return tx.Error
	}

	// Check if token is deleted already
	if entity.DeletedAt.Valid {
		return errors.New("token already deleted")
	}

	// Delete token
	tx = s.Token.Delete(entity)

	return tx.Error
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

func (s *service) GetUser(claims *Claims) (*userModels.DetailModel, error) {
	entity, tx := s.User.Get(map[string]any{
		"ID": claims.UserID,
	})

	if tx.Error != nil {
		return nil, tx.Error
	}

	model := userModels.DetailModelFromEntity(entity)

	return model, nil
}

func (s *service) UpdateUserOnline(user *userModels.DetailModel) error {
	currentTime := time.Now()

	user.LastOnline = &currentTime

	tx := s.User.Update(user.ToEntity())
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
