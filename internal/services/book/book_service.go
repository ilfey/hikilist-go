package book

import (
	"github.com/google/uuid"
	bookEntities "<package_name>/internal/entities/book"
	bookModels "<package_name>/internal/models/book"
	"<package_name>/internal/repositories/book"
)

type Service struct {
	repository *book.Repository
}

// Service constructor
func NewService(repository *book.Repository) *Service {
	return &Service{repository: repository}
}

// Create book
func (service *Service) Create(model *bookModels.CreateModel) {
	// Create entity
	bookEntity := bookEntities.Entity{
		Uuid: uuid.New(),
		Name: model.Name,
	}

	service.repository.Create(bookEntity)
}