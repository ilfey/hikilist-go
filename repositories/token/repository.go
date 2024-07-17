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
	Create(*entities.Token) *gorm.DB

	// Получение токена.
	//
	// Принимает фильтр.
	// Возвращает сущность и транзакцию.
	Get(map[string]any) (*entities.Token, *gorm.DB)

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
	Delete(*entities.Token) *gorm.DB
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
func (r *repository) Create(entity *entities.Token) *gorm.DB {
	tx := r.db.Create(entity)

	return tx
}

// Получение одного пользователя
func (r *repository) Get(query map[string]any) (*entities.Token, *gorm.DB) {
	entity := &entities.Token{}

	tx := r.db.First(entity, query)

	return entity, tx
}

// Получение токенов.
func (r *repository) Find() ([]*entities.Token, *gorm.DB) {
	var entities []*entities.Token
	tx := r.db.Find(&entities)

	return entities, tx
}

// Обновление токена.
//
// Принимает сущность. У сущности должен быть ID.
// Возвращает транзакцию.
func (r *repository) Update(entity *entities.Token) *gorm.DB {
	tx := r.db.Save(entity)

	return tx
}

// Удаление токена.
func (r *repository) Delete(entity *entities.Token) *gorm.DB {
	tx := r.db.Delete(entity)

	return tx
}