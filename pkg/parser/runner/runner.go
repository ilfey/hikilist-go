package runner

import (
	"context"

	"github.com/ilfey/hikilist-go/pkg/parser/extractor"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
)

type Runner struct {
	logger    *logrus.Logger
	extractor *Extractor
}

func New(logger *logrus.Logger, ex extractor.Extractor) *Runner {
	return &Runner{
		logger: logger,
		extractor: Wrap(
			logger.WithField("source", ex.Source().GetName()),
			ex,
		),
	}
}

func (r *Runner) Run(ctx context.Context) (*Result, error) {
	result := &Result{}

	stats, err := r.extractor.RunAnime(ctx)
	if err != nil {
		result.Anime = stats

		r.logger.Warnf("An error occurred while parsing anime %v", err)

		if eris.Is(err, extractor.ErrExtractorsStopped) {
			return result, nil
		}

		return result, err
	}

	r.logger.Info("Animes was parsed")

	r.logger.Infof(
		"Anime stats - %s",
		stats.String(),
	)

	return result, nil
}
