package runner

import (
	"context"
	"time"

	"github.com/ilfey/hikilist-go/pkg/parser/extractor"
	"github.com/ilfey/hikilist-go/pkg/parser/statistic"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
)

type Extractor struct {
	logger logrus.FieldLogger
	extractor.Extractor
}

func Wrap(logger logrus.FieldLogger, ex extractor.Extractor) *Extractor {
	return &Extractor{
		logger:    logger,
		Extractor: ex,
	}
}

/*
RunAnime runs anime extractor.

Errors:

	extractor.ErrExtractorsStopped // if need to stop all extractors
	extractor.ErrStopParser // if need to stop parser
*/
func (e *Extractor) RunAnime(ctx context.Context) (statistic.Readable, error) {
	stats := NewStatistic()

	anime := NewAnimeRunner(
		e.logger.WithField("extractor", "anime"),
		e.Anime(),
		stats,
	)

	timeout := e.GetTimeout()

	for {
		select {
		case <-ctx.Done():
			return stats, ctx.Err()
		default:
			stats.AddPages(1)

			e.logger.Debugf("Tick %d", stats.Pages())

			err := anime.Tick(ctx)
			if err != nil {
				// If page is empty, animes was parsed.
				if eris.Is(err, extractor.ErrPageEmpty) {
					return stats, nil
				}

				// Stop extractor.
				if eris.Is(err, extractor.ErrExtractorStopped) {
					return stats, nil
				}

				return stats, err
			}

			// Wait timeout before next tick.
			e.logger.Infof("Waiting %s...", timeout)

			time.Sleep(timeout)
		}
	}
}
