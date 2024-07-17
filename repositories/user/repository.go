package userRepository

import (
	"github.com/ilfey/hikilist-go/data/entities"
	"gorm.io/gorm"
)

// Репозиторий пользователя.
//
// Предоставляет функционал управления пользователями через сущности.
// Следует использовать абстракцию в виде сервиса,
// реализующего данный функционал через модели.
type Repository interface {
	// Создание пользователя.
	//
	// Принимает сущность.
	// Возвращает транзакцию.
	Create(*entities.User) *gorm.DB

	// Получение пользователя.
	//
	// Принимает фильтр.
	// Возвращает сущность и транзакцию.
	Get(map[string]any) (*entities.User, *gorm.DB)

	// Получение пользователей.
	//
	// Возвращает список сущностей и транзакцию.
	Find() ([]*entities.User, *gorm.DB)

	// Обновление пользователя.
	//
	// Принимает ID и сущность.
	// Возвращает транзакцию.
	Update(*entities.User) *gorm.DB

}

// Имплементация.

type repository struct {
	db *gorm.DB
}

// Конструктор репозитория пользователя
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Создание пользователя
func (r *repository) Create(entity *entities.User) *gorm.DB {
	tx := r.db.Create(entity)

	return tx
}

// Получение одного пользователя
func (r *repository) Get(query map[string]any) (*entities.User, *gorm.DB) {
	entity := &entities.User{}

	tx := r.db.First(entity, query)

	return entity, tx
}

// Получение пользователей
func (r *repository) Find() ([]*entities.User, *gorm.DB) {
	var entities []*entities.User
	tx := r.db.Find(&entities)

	return entities, tx
}

// Обновление пользователя.
//
// Принимает сущность. У сущности должен быть ID.
// Возвращает транзакцию.
func (r *repository) Update(entity *entities.User) *gorm.DB {
	tx := r.db.Save(entity)

	return tx
}