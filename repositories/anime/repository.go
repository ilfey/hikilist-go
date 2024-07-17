package animeRepository

import (
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
	Create(*entities.Anime) *gorm.DB

	// Get получение аниме.
	//
	// Принимает фильтры.
	// Возвращает сущность и транзакцию.
	Get(map[string]any) (*entities.Anime, *gorm.DB)

	// Find получение аниме.
	//
	// Возвращает массив сущностей и транзакцию.
	Find() ([]*entities.Anime, *gorm.DB)
}

// Имплементация
//
// Прелоадинг (Preload) нужен для заполнения полей сущности.

// Репозиторий аниме
type repository struct {
	db *gorm.DB
}

// Конструктор репозитория аниме
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Создание аниме
func (r *repository) Create(entity *entities.Anime) *gorm.DB {

	tx := r.db.Preload("Related").Create(entity).First(entity)

	return tx
}

// Получение одного аниме
func (r *repository) Get(query map[string]any) (*entities.Anime, *gorm.DB) {
	entity := &entities.Anime{}

	tx := r.db.Preload("Related").First(entity, query)

	return entity, tx
}

// Получение аниме
func (r *repository) Find() ([]*entities.Anime, *gorm.DB) {
	var entities []*entities.Anime
	tx := r.db.Find(&entities)

	return entities, tx
}
