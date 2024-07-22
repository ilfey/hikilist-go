package userService

import (
	"github.com/ilfey/hikilist-go/data/entities"
	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"
	userRepository "github.com/ilfey/hikilist-go/repositories/user"
	userActionRepository "github.com/ilfey/hikilist-go/repositories/user_action"
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
	user       userRepository.Repository
	userAction userActionRepository.Repository
}

// Конструктор сервиса пользователя
func New(
	user userRepository.Repository,
	userAction userActionRepository.Repository,
) Service {
	return &service{
		user:       user,
		userAction: userAction,
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

	err := s.user.Create(userEntity)
	if err != nil {
		return nil, err
	}

	detailModel := userModels.DetailModelFromEntity(userEntity)

	err = s.createRegisterAction(detailModel)
	if err != nil {
		logger.Errorf("Failed to create register action: %v", err)

		return detailModel, nil
	}

	return detailModel, nil
}

func (s *service) Get(conds ...any) (*userModels.DetailModel, error) {
	var model userModels.DetailModel

	err := s.user.Get(&model, conds...)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// Получение всех пользователей
func (s *service) Find(conds ...any) (*userModels.ListModel, error) {
	var items []*userModels.ListItemModel

	err := s.user.Find(&items, conds...)
	if err != nil {
		return nil, err
	}

	model := userModels.NewListModel(items)

	return model, nil
}

func (s *service) createRegisterAction(user *userModels.DetailModel) error {
	return s.userAction.Create(&entities.UserAction{
		UserID:      user.ID,
		Title:       "Регистрация аккаунта",
		Description: "Начало вашего пути на сайте Hikilist",
	})
}
