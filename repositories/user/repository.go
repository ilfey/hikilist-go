package userRepository

import (
	"github.com/ilfey/hikilist-go/data/entities"
	"gorm.io/gorm"
)

// Репозиторий пользователя
type Respository struct {
	db *gorm.DB
}

// Конструктор репозитория пользователя
func NewRepository(db *gorm.DB) *Respository {
	return &Respository{
		db: db,
	}
}

// Create user
// func (r *UserRespository) Create(entity *entities.User) *gorm.DB {
// 	tx := r.db.Create(entity)

// 	return tx
// }

// Получение одного пользователя
func (r *Respository) Get(query map[string]any) (*entities.User, *gorm.DB) {
	entity := &entities.User{}

	tx := r.db.First(entity, query)

	return entity, tx
}

// Получение пользователей
func (r *Respository) Find() ([]*entities.User, *gorm.DB) {
	var entities []*entities.User
	tx := r.db.Find(&entities)

	return entities, tx
}
