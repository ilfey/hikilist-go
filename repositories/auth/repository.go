package authRepository

import (
	"github.com/ilfey/hikilist-go/data/entities"
	"gorm.io/gorm"
)

// Репозиторий аутентификации
type Respository struct {
	db *gorm.DB
}

// Конструктор репозитория аутентификации
func NewRepository(db *gorm.DB) *Respository {
	return &Respository{
		db: db,
	}
}

// Create Token
// func (r *AuthRespository) Create(entity *entities.Token) *gorm.DB {
// 	tx := r.db.Create(entity)

// 	return tx
// }

// Создание пользователя
func (r *Respository) CreateUser(entity *entities.User) *gorm.DB {
	tx := r.db.Create(entity)

	return tx
}

// func (r *AuthRespository) Get(query map[string]any) (*entities.Token, *gorm.DB) {
// 	entity := &entities.Token{}

// 	tx := r.db.First(entity, query)

// 	return entity, tx
// }
