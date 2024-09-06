package tokenizerInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
)

type Tokenizer interface {
	Generate(userId uint64) (*agg.TokenPair, error)
	Verify(ctx context.Context, token string) (uint64, error)
	Block(ctx context.Context, token string) error
}