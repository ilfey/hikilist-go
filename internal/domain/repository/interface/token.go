package repositoryInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/pkg/postgres"
)

type Token interface {
	WithTx(tx postgres.RW) Token

	Create(ctx context.Context, cm *agg.TokenCreate) error

	Has(ctx context.Context, token string) (bool, error)

	Delete(ctx context.Context, conds any) error
}
