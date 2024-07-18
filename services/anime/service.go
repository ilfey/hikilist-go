package animeService

import (
	"errors"

	"github.com/ilfey/hikilist-go/data/entities"
	animeModels "github.com/ilfey/hikilist-go/data/models/anime"
	"github.com/ilfey/hikilist-go/internal/logger"
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
	Create(*animeModels.CreateModel) (*animeModels.DetailModel, error)

	// Get получение аниме.
	//
	// Принимает фильтры для поиска аниме.
	// Возвращает модель аниме и транзакцию.
	Get(...any) (*animeModels.DetailModel, error)

	// Find получение списка аниме.
	//
	// Возвращает модель списка аниме и транзакцию.
	Find(...any) (*animeModels.ListModel, error)
	Paginate(p *animeModels.Paginate, whereArgs ...any) (*animeModels.ListModel, error)

	ResolveShiki(*animeModels.ShikiDetailModel) error
}

// Имплементация.

// Сервис аниме.
type service struct {
	repository animeRepository.Repository
}

// Конструктор сервиса аниме
func New(repository animeRepository.Repository) Service {
	return &service{
		repository: repository,
	}
}

// Создание аниме
func (s *service) Create(model *animeModels.CreateModel) (*animeModels.DetailModel, error) {
	entity := model.ToEntity()

	err := s.repository.Create(entity)
	if err != nil {
		return nil, err
	}

	detailModel := animeModels.DetailModelFromEntity(entity)

	return detailModel, nil
}

func (s *service) ResolveShiki(model *animeModels.ShikiDetailModel) error {
	return s.repository.Transaction(func(tx *gorm.DB) error {
		var searchEntity entities.Anime

		if err := tx.First(&searchEntity, map[string]any{
			"shiki_id": model.ID,
		}).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			createEntity := model.ToEntity()

			return tx.Create(createEntity).Error
		}

		// Если время обновления существующей записи больше времени создания новой записи
		if searchEntity.UpdatedAt.After(*model.UpdatedAt) {
			return nil
		}

		updateEntity := model.ToEntity()

		updateEntity.ID = searchEntity.ID

		return tx.Model(&updateEntity).Updates(updateEntity).Error
	})
}

func (s *service) Get(conds ...any) (*animeModels.DetailModel, error) {
	var model animeModels.DetailModel

	err := s.repository.Get(&model, conds...)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// Получение всех аниме
func (s *service) Find(conds ...any) (*animeModels.ListModel, error) {
	var items []*animeModels.ListItemModel

	err := s.repository.Find(&items, conds...)
	if err != nil {
		return nil, err
	}

	model := animeModels.NewListModel(items)

	return model, nil
}

func (s *service) Paginate(p *animeModels.Paginate, whereArgs ...any) (*animeModels.ListModel, error) {
	var items []*animeModels.ListItemModel

	err := s.repository.ScopedFind(&items, p.Scope, whereArgs...)
	if err != nil {
		return nil, err
	}

	model := animeModels.NewListModel(items)

	count, err := s.repository.Count(whereArgs...)
	if err != nil {
		logger.Errorf("Error getting count: %v", err)

		return model, nil
	}

	model.WithCount(count)

	return model, nil
}
