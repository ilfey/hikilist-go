package book

import (
	"fmt"
	"<package_name>/internal/entities/book"
)

type Repository struct {
	database []book.Entity
}

// Repository constructor
func NewRepository() *Repository {
	return &Repository{
		database: make([]book.Entity, 0),
	}
}

// Create book
func (repository *Repository) Create(entity book.Entity) {
	repository.database = append(repository.database, entity)
	fmt.Println(repository.database)
}