package animeService

import (
	"github.com/ilfey/hikilist-go/data/entities"
	animeModels "github.com/ilfey/hikilist-go/data/models/anime"
	animeRepository "github.com/ilfey/hikilist-go/repositories/anime"
	"gorm.io/gorm"
)

// Сервис аниме.
//
// Предоставляет управление базой данных аниме, через модели.
// В методах в качестве *gorm.DB возвращается транзакция операции.
type Service interface {
	// Create создает аниме.
	//
	// Возвращает модель созданного аниме и транзакцию.
	Create(*animeModels.CreateModel) (*animeModels.DetailModel, *gorm.DB)

	// Get получение аниме.
	//
	// Принимает фильтры для поиска аниме.
	// Возвращает модель аниме и транзакцию.
	Get(query map[string]any) (*animeModels.DetailModel, *gorm.DB)

	// GetByID получение аниме по ID.
	//
	// Возвращает модель аниме и транзакцию.
	GetByID(uint64) (*animeModels.DetailModel, *gorm.DB)

	// Find получение списка аниме.
	//
	// Возвращает модель списка аниме и транзакцию.
	Find() (*animeModels.AnimeListModel, *gorm.DB)
}

// Имплементация.

// Сервис аниме.
type service struct {
	repository animeRepository.Repository
}

// Конструктор сервиса аниме
func NewService(repository animeRepository.Repository) Service {
	return &service{repository: repository}
}

// Создание аниме
func (s *service) Create(model *animeModels.CreateModel) (*animeModels.DetailModel, *gorm.DB) {
	var related []*entities.Anime // Массив связанных аниме (пустые сущности с указанными id)

	if model.Related != nil { // Если в модели указан массив связанных аниме
		related = make([]*entities.Anime, len(*model.Related))

		for i, item := range *model.Related { // Заполняем массив моделями
			related[i] = &entities.Anime{}
			related[i].ID = item
		}
	}

	// Создаем сущность
	AnimeEntity := &entities.Anime{
		Title:            model.Title,
		Description:      model.Description,
		Poster:           model.Poster,
		Episodes:         model.Episodes,
		EpisodesReleased: model.EpisodesReleased,

		MalID:   model.MalID,
		ShikiID: model.ShikiID,

		Related: related,
	}

	tx := s.repository.Create(AnimeEntity)

	detailModel := animeModels.DetailModelFromEntity(AnimeEntity)

	return detailModel, tx
}

func (s *service) Get(query map[string]any) (*animeModels.DetailModel, *gorm.DB) {
	entity, tx := s.repository.Get(query)

	model := animeModels.DetailModelFromEntity(entity)

	return model, tx
}

// Получение аниме по ID
func (s *service) GetByID(id uint64) (*animeModels.DetailModel, *gorm.DB) {
	return s.Get(map[string]any{
		"ID": id,
	})
}

// Получение всех аниме
func (s *service) Find() (*animeModels.AnimeListModel, *gorm.DB) {
	entities, tx := s.repository.Find()

	model := animeModels.AnimeListModelFromEntities(entities, tx.RowsAffected)

	return model, tx
}
