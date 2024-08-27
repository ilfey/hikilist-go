package runner

import (
	"context"

	"github.com/ilfey/hikilist-go/pkg/parser/extractor"
	"github.com/ilfey/hikilist-go/pkg/parser/statistic"
	"github.com/sirupsen/logrus"
)

type AnimeRunner struct {
	logger logrus.FieldLogger
	extractor.AnimeExtractor
	stats statistic.RW
}

func NewAnimeRunner(logger logrus.FieldLogger, extractor extractor.AnimeExtractor, stats statistic.RW) *AnimeRunner {
	return &AnimeRunner{
		logger:         logger,
		AnimeExtractor: extractor,
		stats:          stats,
	}
}

/*
Tick fetches list and details of anime.

Returns:

	extractor.ErrPageEmpty // if list is empty
	extractor.ErrExtractorStopped // if need to stop extractor
	extractor.ErrExtractorsStopped // if need to stop all extractors
	extractor.ErrStopParser // if need to stop parser
*/
func (e *AnimeRunner) Tick(ctx context.Context) error {
	stats := e.stats

	// Fetch list of anime.
RETRY_FETCH_LIST:
	listModel, err := e.FetchList(stats.Pages())
	if err != nil {
		// Handle network error.
		action := e.OnFetchError(err)

		if action == extractor.ActionRetry {
			goto RETRY_FETCH_LIST
		}

		err = e.applyAction(action, err)
		if err != nil {
			return err
		}

		return nil
	}

	if len(listModel.IDs) == 0 {
		e.logger.Errorf("Page %d is empty", stats.Pages())

		return extractor.ErrPageEmpty
	}

	for _, id := range listModel.IDs {
		// Add to fetched.
		stats.AddFetched(1)

		// Fetch detail.
	RETRY_FETCH_DETAIL:
		detailModel, err := e.FetchById(id)
		if err != nil {
			// Handle network error.
			action := e.OnFetchError(err)

			if action == extractor.ActionRetry {
				goto RETRY_FETCH_DETAIL
			}

			err = e.applyAction(action, err)
			if err != nil {
				return err
			}

			continue
		}

		// Resolve detail.
	RETRY_RESOLVE:
		err = e.Resolve(ctx, detailModel)
		if err != nil {
			action := e.OnResolveError(err)

			if action == extractor.ActionRetry {
				goto RETRY_RESOLVE
			}

			err = e.applyAction(action, err)
			if err != nil {
				return err
			}

			continue
		}

		// Add to resolved.
		stats.AddResolved(1)
	}

	return nil
}

func (e *AnimeRunner) applyAction(action extractor.Action, err error) error {
	switch action {
	case extractor.ActionSkip:
		e.stats.AddScipped(1)
	case extractor.ActionResolved:
		e.stats.AddResolved(1)
	case extractor.ActionStopExtractors:
		return extractor.ErrExtractorsStopped
	case extractor.ActionStopExtractor:
		return extractor.ErrExtractorStopped
	case extractor.ActionStopParser:
		return extractor.ErrParserStopped
	}

	return err
}
