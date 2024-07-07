package userService

import (
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	userRepository "github.com/ilfey/hikilist-go/repositories/user"
	"gorm.io/gorm"
)

// Сервис пользователя
type Service struct {
	repository *userRepository.Respository
}

// Конструктор сервиса пользователя
func NewService(repository *userRepository.Respository) *Service {
	return &Service{repository: repository}
}

// Получение пользователя по ID
func (s *Service) GetByID(id uint64) (*userModels.UserDetailModel, *gorm.DB) {
	entity, tx := s.repository.Get(map[string]any{
		"ID": id,
	})

	model := userModels.UserDetailModelFromEntity(entity)

	return model, tx
}

// Получение пользользователя по Username
func (s *Service) GetByUsername(username string) (*userModels.UserDetailModel, *gorm.DB) {
	entity, tx := s.repository.Get(map[string]any{
		"Username": username,
	})

	model := userModels.UserDetailModelFromEntity(entity)

	return model, tx
}

// Получение всех пользователей
func (s *Service) Find() (*userModels.UserListModel, *gorm.DB) {
	entities, tx := s.repository.Find()

	model := userModels.UserListModelFromEntities(entities, tx.RowsAffected)

	return model, tx
}
