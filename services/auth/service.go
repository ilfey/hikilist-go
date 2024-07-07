package authService

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authConfig "github.com/ilfey/hikilist-go/config/auth"
	"github.com/ilfey/hikilist-go/data/entities"
	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/ilfey/hikilist-go/internal/utils/errorsx"
	authRepository "github.com/ilfey/hikilist-go/repositories/auth"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Сервис аутентификации
type Service struct {
	config     *authConfig.Config
	repository *authRepository.Respository
}

// Конструктор сервиса аутентификации
func NewService(config *authConfig.Config, repository *authRepository.Respository) *Service {
	return &Service{
		config:     config,
		repository: repository,
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

// Создание пользователя
func (s *Service) CreateUser(model *authModels.RegisterModel) (*userModels.UserDetailModel, *gorm.DB) {
	hashedPassword := errorsx.Must(
		bcrypt.GenerateFromPassword(
			[]byte(model.Password),
			bcrypt.DefaultCost,
		),
	)

	// Create entity
	userEntity := &entities.User{
		Username: model.Username,
		Password: string(hashedPassword),
	}

	tx := s.repository.CreateUser(userEntity)

	detailModel := userModels.UserDetailModelFromEntity(userEntity)

	return detailModel, tx
}

func (s *Service) CompareUserPassword(model *userModels.UserDetailModel, password string) bool {
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

func (s *Service) GenerateTokens(model *userModels.UserDetailModel) *authModels.TokensModel {
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

	return &authModels.TokensModel{
		Access:  accessToken,
		Refresh: refreshToken,
	}
}

func (s *Service) generateToken(claims *Claims) string {
	return errorsx.Must(
		jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			claims,
		).SignedString(s.config.Secret),
	)
}

// Парсинг токена
func (s *Service) ParseToken(token string) (*Claims, bool) {
	claims := new(Claims)
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if iss, err := token.Claims.GetIssuer(); iss != s.config.Issuer || err != nil {
			return nil, errors.New("unexpected issuer")
		}

		return s.config.Secret, nil
	})
	if err != nil || !parsedToken.Valid {
		return nil, false
	}

	return claims, true
}
