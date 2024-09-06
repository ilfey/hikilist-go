package repositoryInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/pkg/postgres"
)

type AnimeCollection interface {
	WithTx(tx postgres.RW) AnimeCollection

	AddAnimes(ctx context.Context, aam *dto.CollectionAddAnimeRequestDTO) error

	RemoveAnimes(ctx context.Context, ram *dto.CollectionRemoveAnimeRequestDTO) error
}
