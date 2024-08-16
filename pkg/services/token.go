package services

import (
	"context"

	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/pkg/models/token"
	"github.com/ilfey/hikilist-go/pkg/repositories"
)

type Token interface {
	GetDBView(token string) string

	Create(ctx context.Context, cm *token.CreateModel) error
	Get(ctx context.Context, conds any) (*token.DetailModel, error)
	Delete(ctx context.Context, conds any) error
}

type TokenImpl struct {
	token repositories.Token
}

func (s *TokenImpl) Create(ctx context.Context, cm *token.CreateModel) error {
	err := s.token.Create(ctx, cm)
	if err != nil {
		logger.Debugf("Error occurred while creating token %v", err)

		return err
	}

	return nil
}

func (s *TokenImpl) Get(ctx context.Context, conds any) (*token.DetailModel, error) {
	dm, err := s.token.Get(ctx, conds)
	if err != nil {
		logger.Debugf("Error occurred while getting token %v", err)

		return nil, err
	}

	return dm, nil
}

func (s *TokenImpl) Delete(ctx context.Context, conds any) error {
	err := s.token.Delete(ctx, conds)
	if err != nil {
		logger.Debugf("Error occurred while deleting token %v", err)

		return err
	}

	return nil
}
