package animeRepository

import (
	"database/sql"

	"github.com/ilfey/hikilist-go/data/entities"
	"gorm.io/gorm"
)

// Репозиторий аниме.
//
// Предоставляет управления аниме через сущности.
// Следует использовать абстракцию в виде сервиса,
// реализующего данный функционал через модели.
type Repository interface {
	// Create создание аниме.
	//
	// Принимает сущность, возвращает транзакцию.
	// Возвращает транзакцию.
	// Также обновляет переданной поля сущности.
	Create(*entities.Anime) error

	// Get получение аниме.
	//
	// Принимает фильтры.
	// Возвращает сущность и транзакцию.
	Get(v any, conds ...any) error

	// Find получение аниме.
	//
	// Возвращает массив сущностей и транзакцию.
	Find(v any, conds ...any) error
	ScopedFind(v any, scope func(*gorm.DB) *gorm.DB, whereArgs ...any) error
	Count(whereArgs ...any) (int64, error)

	Transaction(func(tx *gorm.DB) error, ...*sql.TxOptions) error
}

// Имплементация
//
// Прелоадинг (Preload) нужен для заполнения полей сущности.

// Репозиторий аниме
type repository struct {
	db *gorm.DB
}

// Конструктор репозитория аниме
func New(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Создание аниме
func (r *repository) Create(entity *entities.Anime) error {
	result := r.db.Preload("Related").Create(entity).First(entity)

	return result.Error
}

// Получение одного аниме
func (r *repository) Get(v any, conds ...any) error {
	result := r.db.Model(&entities.Anime{}).Preload("Related").First(v, conds...)

	return result.Error
}

// Получение аниме
func (r *repository) Find(v any, conds ...any) error {
	result := r.db.Model(&entities.Anime{}).Find(v, conds...)

	return result.Error
}

func (r *repository) ScopedFind(v any, scope func(*gorm.DB) *gorm.DB, whereArgs ...any) error {
	result := r.db.Model(&entities.Anime{}).Scopes(scope).Where(whereArgs).Find(v)

	return result.Error
}

func (r *repository) Count(whereArgs ...any) (int64, error) {
	var count int64

	result := r.db.Model(&entities.Anime{}).Where(whereArgs).Count(&count)

	return count, result.Error
}

func (r *repository) Transaction(fn func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return r.db.Transaction(fn, opts...)
}
