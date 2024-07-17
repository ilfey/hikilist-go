package authRepository

import (
	"gorm.io/gorm"
)

// WARN: Not used this package

// Репозиторий аутентификации.
type Repository interface{}

// Имплементация.

type repository struct {
	db *gorm.DB
}

// Конструктор репозитория аутентификации
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create Token
// func (r *AuthRespository) Create(entity *entities.Token) *gorm.DB {
// 	tx := r.db.Create(entity)

// 	return tx
// }

// func (r *AuthRespository) Get(query map[string]any) (*entities.Token, *gorm.DB) {
// 	entity := &entities.Token{}

// 	tx := r.db.First(entity, query)

// 	return entity, tx
// }
