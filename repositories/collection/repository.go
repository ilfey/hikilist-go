package collectionRepository

import (
	"github.com/ilfey/hikilist-go/data/entities"
	"gorm.io/gorm"
)

type Repository interface {
	Create(entity *entities.Collection) error
	Get(v any, conds ...any) error
	Find(v any, conds ...any) error
	ScopedFind(v any, scope func(*gorm.DB) *gorm.DB, whereArgs ...any) error
	Count(whereArgs ...any) (int64, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(entity *entities.Collection) error {
	result := r.db.Create(entity).Preload("User").First(entity)

	return result.Error
}

func (r *repository) Get(v any, conds ...any) error {
	result := r.db.Model(&entities.Collection{}).First(v, conds...)

	return result.Error
}

func (r *repository) Find(v any, conds ...any) error {
	result := r.db.Model(&entities.Collection{}).Find(v, conds...)

	return result.Error
}

func (r *repository) ScopedFind(v any, scope func(*gorm.DB) *gorm.DB, whereArgs ...any) error {
	result := r.db.Model(&entities.Collection{}).
		Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("*")
		}).
		Scopes(scope).
		Where(whereArgs).
		Find(v)

	return result.Error
}

func (r *repository) Count(whereArgs ...any) (int64, error) {
	var count int64

	result := r.db.Model(&entities.Collection{}).Where(whereArgs).Count(&count)

	return count, result.Error
}
