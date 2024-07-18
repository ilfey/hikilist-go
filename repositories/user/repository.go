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
	Create(*entities.User) error

	// Получение пользователя.
	//
	// Принимает фильтр.
	// Возвращает сущность и транзакцию.
	Get(v any, conds ...any) error

	// Получение пользователей.
	//
	// Возвращает список сущностей и транзакцию.
	Find(v any, conds ...any) error

	// Обновление пользователя.
	//
	// Принимает ID и сущность.
	// Возвращает транзакцию.
	Update(*entities.User) error
}

// Имплементация.

type repository struct {
	db *gorm.DB
}

// Конструктор репозитория пользователя
func New(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Создание пользователя
func (r *repository) Create(entity *entities.User) error {
	result := r.db.Create(entity)

	return result.Error
}

// Получение одного пользователя
func (r *repository) Get(v any, conds ...any) error {

	result := r.db.Model(&entities.User{}).First(v, conds...)

	return result.Error
}

// Получение пользователей
func (r *repository) Find(v any, conds ...any) error {
	result := r.db.Model(&entities.User{}).Find(v, conds...)

	return result.Error
}

// Обновление пользователя.
//
// Принимает сущность. У сущности должен быть ID.
// Возвращает транзакцию.
func (r *repository) Update(entity *entities.User) error {
	result := r.db.Save(entity)

	return result.Error
}
