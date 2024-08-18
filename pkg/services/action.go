package services

import (
	"context"
	"fmt"

	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/pkg/models/action"
	"github.com/ilfey/hikilist-go/pkg/repositories"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type Action interface {
	// CreateRegisterAction(ctx context.Context, userId uint) error
	// CreateCollectionCreateAction(ctx context.Context, userId uint, collectionTitle string) error
	// CreateUpdateUsernameAction(ctx context.Context, userId uint, oldUsername, newUsername string) error

	// Create(ctx context.Context, cm *action.CreateModel) error
	GetListModel(ctx context.Context, p *paginate.Paginator, conds any) (*action.ListModel, error)
	// FindWithPaginator(ctx context.Context, p *paginate.Paginator, conds any) ([]*action.ListItemModel, error)
	// Count(ctx context.Context, conds any) (uint, error)
}

type ActionImpl struct {
	logger logrus.FieldLogger

	action repositories.Action
}

func NewAction(logger logrus.FieldLogger, actionRepo repositories.Action) Action {
	return &ActionImpl{
		logger: logger,

		action: actionRepo,
	}
}

func (s *ActionImpl) CreateRegisterAction(ctx context.Context, userId uint) error {
	cm := action.CreateModel{
		UserID:      userId,
		Title:       "Регистрация аккаунта",
		Description: "Это начало вашего пути на сайте Hikilist.",
	}

	return s.Create(ctx, &cm)
}

func (s *ActionImpl) CreateCollectionCreateAction(ctx context.Context, userId uint, collectionTitle string) error {
	cm := action.CreateModel{
		UserID:      userId,
		Title:       "Создание коллекции",
		Description: fmt.Sprintf("Вы создали коллекцию \"%s\".", collectionTitle),
	}

	return s.Create(ctx, &cm)
}

func (s *ActionImpl) CreateUpdateUsernameAction(ctx context.Context, userId uint, oldUsername, newUsername string) error {
	cm := action.CreateModel{
		UserID: userId,
		Title:  "Обновление никнейма",
		Description: fmt.Sprintf(
			"%s останется в прошлом. Продолжим путь как %s.",
			oldUsername,
			newUsername,
		),
	}

	return s.Create(ctx, &cm)
}

func (s *ActionImpl) Create(ctx context.Context, cm *action.CreateModel) error {
	err := s.action.Create(ctx, cm)
	if err != nil {
		s.logger.Debugf("Error occurred while creating action %v", err)

		return err
	}

	return nil
}

func (s *ActionImpl) GetListModel(ctx context.Context, p *paginate.Paginator, conds any) (*action.ListModel, error) {
	var lm action.ListModel

	g := errgroup.Group{}

	g.Go(func() error {
		items, err := s.FindWithPaginator(ctx, p, conds)
		if err != nil {
			return err
		}

		lm.Results = items

		return nil
	})

	g.Go(func() error {
		count, err := s.Count(ctx, conds)
		if err != nil {
			return err
		}

		lm.Count = &count

		return nil
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	return &lm, nil
}

func (s *ActionImpl) FindWithPaginator(ctx context.Context, p *paginate.Paginator, conds any) ([]*action.ListItemModel, error) {
	items, err := s.action.FindWithPaginator(ctx, p, conds)
	if err != nil {
		s.logger.Debugf("Error occurred while finding actions %v", err)

		return nil, err
	}

	return items, nil
}

func (s *ActionImpl) Count(ctx context.Context, conds any) (uint, error) {
	count, err := s.action.Count(ctx, conds)
	if err != nil {
		s.logger.Debugf("Error occurred while counting actions %v", err)

		return 0, err
	}

	return count, nil
}
