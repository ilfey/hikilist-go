package user

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/user/interface"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"golang.org/x/sync/errgroup"
)

type CRUDService struct {
	logger    loggerInterface.Logger
	user      repositoryInterface.User
	validator validatorInterface.User
}

func NewCRUDService(container diInterface.ServiceContainer) (userInterface.CRUD, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	userRepository, err := container.GetUserRepository()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	return &CRUDService{
		logger: log,
		user:   userRepository,
	}, nil
}

func (s *CRUDService) Create(ctx context.Context, req *dto.UserCreateRequestDTO) error {
	// Validate request.
	if err := s.validator.ValidateCreateRequestDTO(req); err != nil {
		return s.logger.LogPropagate(err)
	}

	err := s.user.Create(ctx, req)
	if err != nil {
		return s.logger.LogPropagate(err)
	}

	return nil
}

func (s *CRUDService) Detail(ctx context.Context, req *dto.UserDetailRequestDTO) (*agg.UserDetail, error) {
	// Validate request.
	if err := s.validator.ValidateDetailRequestDTO(req); err != nil {
		return nil, s.logger.LogPropagate(err)
	}

	detail, err := s.user.Get(ctx, map[string]any{
		"id": req.UserID,
	})
	if err != nil {
		return nil, s.logger.LogPropagate(err)
	}

	return detail, nil
}

func (s *CRUDService) List(ctx context.Context, req *dto.UserListRequestDTO) (*agg.UserList, error) {
	// Validate request.
	if err := s.validator.ValidateListRequestDTO(req); err != nil {
		return nil, s.logger.LogPropagate(err)
	}

	var lm agg.UserList

	g := &errgroup.Group{}

	g.Go(func() error {
		items, err := s.user.Find(ctx, req, nil)
		if err != nil {
			return err
		}

		lm.Results = items

		return nil
	})

	g.Go(func() error {
		count, err := s.user.Count(ctx, nil)
		if err != nil {
			return err
		}

		lm.Count = &count

		return nil
	})

	err := g.Wait()
	if err != nil {
		return nil, s.logger.LogPropagate(err)
	}

	return &lm, nil
}

func (s *CRUDService) ChangeUsername(ctx context.Context, userId uint64, oldUsername string, newUsername string) error {
	err := s.user.UpdateUsername(ctx, userId, oldUsername, newUsername)
	if err != nil {
		return s.logger.LogPropagate(err)
	}

	return nil
}

func (s *CRUDService) UpdateLastOnline(ctx context.Context, userId uint64) error {
	err := s.user.UpdateLastOnline(ctx, userId)
	if err != nil {
		return s.logger.LogPropagate(err)
	}

	return nil
}

func (s *CRUDService) UpdatePassword(ctx context.Context, userId uint64, hash string) error {
	err := s.user.UpdatePassword(ctx, userId, hash)
	if err != nil {
		return s.logger.LogPropagate(err)
	}

	return nil
}

func (s *CRUDService) Delete(ctx context.Context, conds any) error {
	err := s.user.Delete(ctx, conds)
	if err != nil {
		return s.logger.LogPropagate(err)
	}

	return nil
}
