package services

import (
	"context"

	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/pkg/models/user"
	"github.com/ilfey/hikilist-go/pkg/repositories"
	"golang.org/x/sync/errgroup"
)

type User interface {
	Create(ctx context.Context, cm *user.CreateModel) error
	Get(ctx context.Context, conds any) (*user.DetailModel, error)
	GetListModel(ctx context.Context, p *paginate.Paginator, conds any) (*user.ListModel, error)
	// FindWithPaginator(ctx context.Context, p *paginate.Paginator, conds any) ([]*user.ListItemModel, error)
	// Count(ctx context.Context, conds any) (uint, error)
	ChangeUsername(ctx context.Context, userId uint, oldUsername, newUsername string) error
	UpdateLastOnline(ctx context.Context, userId uint) error
	UpdatePassword(ctx context.Context, userId uint, hash string) error
	Delete(ctx context.Context, conds any) error
}

type UserImpl struct {
	user repositories.User
}

func NewUser(userRepo repositories.User) User {
	return &UserImpl{
		user: userRepo,
	}
}

func (s *UserImpl) Create(ctx context.Context, cm *user.CreateModel) error {
	err := s.user.Create(ctx, cm)
	if err != nil {
		logger.Debugf("Error occurred while creating user %v", err)

		return err
	}

	return nil
}

func (s *UserImpl) Get(ctx context.Context, conds any) (*user.DetailModel, error) {
	dm, err := s.user.Get(ctx, conds)
	if err != nil {
		logger.Debugf("Error occurred while getting user %v", err)

		return nil, err
	}

	return dm, nil
}

func (s *UserImpl) GetListModel(ctx context.Context, p *paginate.Paginator, conds any) (*user.ListModel, error) {
	var lm user.ListModel

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

func (s *UserImpl) FindWithPaginator(ctx context.Context, p *paginate.Paginator, conds any) ([]*user.ListItemModel, error) {
	items, err := s.user.FindWithPaginator(ctx, p, conds)
	if err != nil {
		logger.Debugf("Error occurred while finding users %v", err)

		return nil, err
	}

	return items, nil
}

func (s *UserImpl) Count(ctx context.Context, conds any) (uint, error) {
	count, err := s.user.Count(ctx, conds)
	if err != nil {
		logger.Debugf("Error occurred while counting users %v", err)

		return 0, err
	}

	return count, nil
}

func (s *UserImpl) ChangeUsername(ctx context.Context, userId uint, oldUsername string, newUsername string) error {
	err := s.user.UpdateUsername(ctx, userId, oldUsername, newUsername)
	if err != nil {
		logger.Debugf("Error occurred while updating user username %v", err)

		return err
	}

	return nil
}

func (s *UserImpl) UpdateLastOnline(ctx context.Context, userId uint) error {
	err := s.user.UpdateLastOnline(ctx, userId)
	if err != nil {
		logger.Debugf("Error occurred while updating user last online %v", err)

		return err
	}

	return nil
}

func (s *UserImpl) UpdatePassword(ctx context.Context, userId uint, hash string) error {
	err := s.user.UpdatePassword(ctx, userId, hash)
	if err != nil {
		logger.Debugf("Error occurred while updating user password %v", err)

		return err
	}

	return nil
}

func (s *UserImpl) Delete(ctx context.Context, conds any) error {
	err := s.user.Delete(ctx, conds)
	if err != nil {
		logger.Debugf("Error occurred while deleting user %v", err)

		return err
	}

	return nil
}
