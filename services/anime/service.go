package animeService

import (
	"github.com/ilfey/hikilist-go/data/entities"
	animeModels "github.com/ilfey/hikilist-go/data/models/anime"
	animeRepository "github.com/ilfey/hikilist-go/repositories/anime"
	"gorm.io/gorm"
)

// Сервис аниме
type Service struct {
	repository *animeRepository.Respository
}

// Конструктор сервиса аниме
func NewService(repository *animeRepository.Respository) *Service {
	return &Service{repository: repository}
}

// Создание аниме
func (s *Service) Create(model *animeModels.AnimeCreateModel) (*animeModels.AnimeDetailModel, *gorm.DB) {
	var related []*entities.Anime

	if model.Related != nil {
		related = make([]*entities.Anime, len(*model.Related))

		for i, item := range *model.Related {
			related[i] = &entities.Anime{}
			related[i].ID = item
		}
	}

	// Create entity
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

	detailModel := animeModels.AnimeDetailModelFromEntity(AnimeEntity)

	return detailModel, tx
}

// Получение аниме по ID
func (s *Service) GetByID(id uint64) (*animeModels.AnimeDetailModel, *gorm.DB) {
	entity, tx := s.repository.Get(map[string]any{
		"ID": id,
	})

	model := animeModels.AnimeDetailModelFromEntity(entity)

	return model, tx
}

// Получение всех аниме
func (s *Service) Find() (*animeModels.AnimeListModel, *gorm.DB) {
	entities, tx := s.repository.Find()

	model := animeModels.AnimeListModelFromEntities(entities, tx.RowsAffected)

	return model, tx
}
