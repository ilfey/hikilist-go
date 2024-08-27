package extractor

import (
	"context"

	"github.com/ilfey/hikilist-go/pkg/parser/models"
)

type AnimeExtractor interface {
	// FetchById returns anime detail model.
	FetchById(id uint) (*models.AnimeDetailModel, error)

	// FetchList returns anime list model.
	FetchList(page uint) (*models.AnimeListModel, error)

	// Resolve resolves anime detail model.
	Resolve(ctx context.Context, detailModel *models.AnimeDetailModel) error

	OnFetchError(err error) Action

	OnResolveError(err error) Action
}
