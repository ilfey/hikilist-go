package action

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/action/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"golang.org/x/sync/errgroup"
)

type Action struct {
	log loggerInterface.Logger

	action repositoryInterface.Action
}

func NewAction(log loggerInterface.Logger, actionRepo repositoryInterface.Action) actionInterface.Action {
	return &Action{
		log: log,

		action: actionRepo,
	}
}

func (s *Action) GetListDTO(ctx context.Context, p *dto.ActionListRequestDTO, conds any) (*agg.ActionList, error) {
	var lm agg.ActionList

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
		return nil, s.log.LogPropagate(err)
	}

	return &lm, nil
}

func (s *Action) FindWithPaginator(ctx context.Context, p *dto.ActionListRequestDTO, conds any) ([]*agg.ActionListItem, error) {
	items, err := s.action.Find(ctx, p, conds)
	if err != nil {
		return nil, s.log.LogPropagate(err)
	}

	return items, nil
}

func (s *Action) Count(ctx context.Context, conds any) (uint64, error) {
	count, err := s.action.Count(ctx, conds)
	if err != nil {
		return 0, s.log.LogPropagate(err)
	}

	return count, nil
}
