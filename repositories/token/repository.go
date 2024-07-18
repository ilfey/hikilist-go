package tokenRepository

import (
	"github.com/ilfey/hikilist-go/data/entities"
	"gorm.io/gorm"
)

// Репозиторий токена.
//
// Предоставляет функционал управления токенами авторизации через сущности.
// Следует использовать абстракцию в виде сервиса,
// реализующего данный функционал через модели.
type Repository interface {
	// Создание токена.
	//
	// Принимает сущность.
	// Возвращает транзакцию.
	Create(*entities.Token) error

	// Получение токена.
	//
	// Принимает фильтр.
	// Возвращает сущность и транзакцию.
	Get(v any, conds ...any) error
	// Find(v any, conds ...any) error

	// // Получение токенов.
	// //
	// // Возвращает список сущностей и транзакцию.
	// Find() ([]*entities.Token, *gorm.DB)

	// // Обновление токена.
	// //
	// // Принимает сущность.
	// // Возвращает транзакцию.
	// Update(*entities.Token) *gorm.DB

	// Удаление токена.
	//
	// Принимает сущность.
	// Возвращает транзакцию.
	Delete(...any) error
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
func (r *repository) Create(entity *entities.Token) error {
	result := r.db.Create(entity)

	return result.Error
}

// Получение одного пользователя
func (r *repository) Get(v any, conds ...any) error {
	result := r.db.Model(&entities.Token{}).First(v, conds...)

	return result.Error
}

// Удаление токена.
func (r *repository) Delete(conds ...any) error {
	result := r.db.Delete(&entities.Token{}, conds...)

	return result.Error
}
