package userService

import (
	"github.com/ilfey/hikilist-go/data/entities"
	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	userRepository "github.com/ilfey/hikilist-go/repositories/user"
	"golang.org/x/crypto/bcrypt"
)

// Сервис пользователей.
//
// Предоставляет функционал получения пользователей
type Service interface {
	// Create создание пользователя.
	//
	// Возвращает модель созданного пользователя и транзакцию.
	Create(model *authModels.RegisterModel) (*userModels.DetailModel, error)

	// GetByID получение пользователя по ID.
	//
	// Возвращает модель пользователя и транзакцию.
	Get(...any) (*userModels.DetailModel, error)

	// Find получение списка пользователей.
	//
	// Возвращает модель списка пользователей и транзакцию.
	Find(...any) (*userModels.ListModel, error)
}

// Сервис пользователя
type service struct {
	repository userRepository.Repository
}

// Конструктор сервиса пользователя
func New(repository userRepository.Repository) Service {
	return &service{
		repository,
	}
}

// Создание пользователя
func (s *service) Create(model *authModels.RegisterModel) (*userModels.DetailModel, error) {
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

	err := s.repository.Create(userEntity)
	if err != nil {
		return nil, err
	}

	detailModel := userModels.DetailModelFromEntity(userEntity)

	return detailModel, nil
}

func (s *service) Get(conds ...any) (*userModels.DetailModel, error) {
	var model userModels.DetailModel

	err := s.repository.Get(&model, conds...)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// Получение всех пользователей
func (s *service) Find(conds ...any) (*userModels.ListModel, error) {
	var items []*userModels.ListItemModel

	err := s.repository.Find(&items, conds...)
	if err != nil {
		return nil, err
	}

	model := userModels.NewListModel(items)

	return model, nil
}
