package userService

import (
	"github.com/ilfey/hikilist-go/data/entities"
	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	userRepository "github.com/ilfey/hikilist-go/repositories/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Сервис пользователей.
//
// Предоставляет функционал получения пользователей
type Service interface {
	// Create создание пользователя.
	//
	// Возвращает модель созданного пользователя и транзакцию.
	Create(model *authModels.RegisterModel) (*userModels.DetailModel, *gorm.DB)
	
	// GetByID получение пользователя по ID.
	//
	// Возвращает модель пользователя и транзакцию.
	GetByID(uint64) (*userModels.DetailModel, *gorm.DB)

	// GetByUsername получение пользователя по Username.
	//
	// Возвращает модель пользователя и транзакцию.
	GetByUsername(string) (*userModels.DetailModel, *gorm.DB)

	// Find получение списка пользователей.
	//
	// Возвращает модель списка пользователей и транзакцию.
	Find() (*userModels.ListModel, *gorm.DB)
}

// Сервис пользователя
type service struct {
	repository userRepository.Repository
}

// Конструктор сервиса пользователя
func NewService(repository userRepository.Repository) Service {
	return &service{
		repository,
	}
}

// Создание пользователя
func (s *service) Create(model *authModels.RegisterModel) (*userModels.DetailModel, *gorm.DB) {
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

	tx := s.repository.Create(userEntity)

	detailModel := userModels.DetailModelFromEntity(userEntity)

	return detailModel, tx
}

// Получение пользователя по ID
func (s *service) GetByID(id uint64) (*userModels.DetailModel, *gorm.DB) {
	entity, tx := s.repository.Get(map[string]any{
		"ID": id,
	})

	model := userModels.DetailModelFromEntity(entity)

	return model, tx
}

// Получение пользользователя по Username
func (s *service) GetByUsername(username string) (*userModels.DetailModel, *gorm.DB) {
	entity, tx := s.repository.Get(map[string]any{
		"Username": username,
	})

	model := userModels.DetailModelFromEntity(entity)

	return model, tx
}

// Получение всех пользователей
func (s *service) Find() (*userModels.ListModel, *gorm.DB) {
	entities, tx := s.repository.Find()

	model := userModels.UserListModelFromEntities(entities, tx.RowsAffected)

	return model, tx
}
