package animeRepository

import (
	"github.com/ilfey/hikilist-go/data/entities"
	"gorm.io/gorm"
)

// Репозиторий аниме
type Respository struct {
	db *gorm.DB
}

// Конструктор репозитория аниме
func NewRepository(db *gorm.DB) *Respository {
	return &Respository{
		db: db,
	}
}

// Создание аниме
func (r *Respository) Create(entity *entities.Anime) *gorm.DB {
	tx := r.db.Preload("Related").Create(entity).First(entity)

	return tx
}

// Получение одного аниме
func (r *Respository) Get(query map[string]any) (*entities.Anime, *gorm.DB) {
	entity := &entities.Anime{}

	tx := r.db.Preload("Related").First(entity, query)

	return entity, tx
}

// Получение аниме
func (r *Respository) Find() ([]*entities.Anime, *gorm.DB) {
	var entities []*entities.Anime
	tx := r.db.Find(&entities)

	return entities, tx
}
