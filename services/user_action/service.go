package userActionService

import (
	userActionModels "github.com/ilfey/hikilist-go/data/models/user_action"
	"github.com/ilfey/hikilist-go/internal/logger"
	userActionRepository "github.com/ilfey/hikilist-go/repositories/user_action"
)

type Service interface {
	Create(*userActionModels.CreateModel) (*userActionModels.DetailModel, error)
	Get(conds ...any) (*userActionModels.DetailModel, error)
	Paginate(p *userActionModels.Paginate, whereArgs ...any) (*userActionModels.ListModel, error)
}

type service struct {
	repository userActionRepository.Repository
}

func New(repository userActionRepository.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(model *userActionModels.CreateModel) (*userActionModels.DetailModel, error) {
	entity := model.ToEntity()

	err := s.repository.Create(entity)
	if err != nil {
		return nil, err
	}

	return userActionModels.DetailModelFromEntity(entity), nil
}

func (s *service) Get(conds ...any) (*userActionModels.DetailModel, error) {
	var model userActionModels.DetailModel

	err := s.repository.Get(&model, conds...)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (s *service) Paginate(p *userActionModels.Paginate, whereArgs ...any) (*userActionModels.ListModel, error) {
	var items []*userActionModels.ListItemModel

	err := s.repository.ScopedFind(&items, p.Scope, whereArgs...)
	if err != nil {
		return nil, err
	}

	model := userActionModels.NewListModel(items)

	count, err := s.repository.Count(whereArgs...)
	if err != nil {
		logger.Errorf("Error getting count: %v", err)
		return model, nil
	}

	model.WithCount(count)

	return model, nil
}
