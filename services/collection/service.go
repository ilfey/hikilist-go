package collectionService

import (
	collectionModels "github.com/ilfey/hikilist-go/data/models/collection"
	"github.com/ilfey/hikilist-go/internal/logger"
	collectionRepository "github.com/ilfey/hikilist-go/repositories/collection"
)

type Service interface {
	Create(model *collectionModels.CreateModel) (*collectionModels.DetailModel, error)
	Get(conds ...any) (*collectionModels.DetailModel, error)
	Find(conds ...any) ([]*collectionModels.ListItemModel, error)
	Paginate(p *collectionModels.Paginate, whereArgs ...any) (*collectionModels.ListModel, error)
	Count(whereArgs ...any) (int64, error)
}

type service struct {
	repository collectionRepository.Repository
}

func New(repository collectionRepository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(model *collectionModels.CreateModel) (*collectionModels.DetailModel, error) {
	entity := model.ToEntity()

	err := s.repository.Create(entity)
	if err != nil {
		return nil, err
	}

	detailModel := collectionModels.NewDetailModelFromEntity(entity)

	return detailModel, nil
}

func (s *service) Get(conds ...any) (*collectionModels.DetailModel, error) {
	var model collectionModels.DetailModel

	err := s.repository.Get(&model, conds...)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (s *service) Find(conds ...any) ([]*collectionModels.ListItemModel, error) {
	var data []*collectionModels.ListItemModel

	err := s.repository.Find(&data, conds...)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *service) Paginate(p *collectionModels.Paginate, whereArgs ...any) (*collectionModels.ListModel, error) {
	var items []*collectionModels.ListItemModel

	err := s.repository.ScopedFind(&items, p.Scope, whereArgs...)
	if err != nil {
		return nil, err
	}

	model := collectionModels.NewListModel(items)

	count, err := s.repository.Count(whereArgs...)
	if err != nil {
		logger.Errorf("Error getting count: %v", err)
		return model, nil
	}

	model.WithCount(count)

	return model, nil
}

func (s *service) Count(whereArgs ...any) (int64, error) {
	return s.repository.Count(whereArgs...)
}
