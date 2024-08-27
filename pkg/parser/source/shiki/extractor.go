package shiki

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/ilfey/hikilist-go/pkg/parser/extractor"
	"github.com/ilfey/hikilist-go/pkg/parser/source/shiki/config"
	"github.com/sirupsen/logrus"
)

type ShikiExtractor struct {
	config *config.Config // Extractor config.

	source  *ShikiSource         // Source.
	animeEx *ShikiAnimeExtractor // Anime extractor.

	client *resty.Client // HTTP client.
}

func New(logger *logrus.Logger, config *config.Config) extractor.Extractor {
	source := &ShikiSource{
		name:       "Shiki",
		columnName: "shiki_id",
	}

	client := newClient(
		logger.WithField("extractor", source.GetName()),
		config,
	)

	animeEx := newAnimeExtractor(client, config.Anime)

	return &ShikiExtractor{
		config: config,

		source: source,

		client: client,

		animeEx: animeEx,
	}
}

/* Impliment extractor.Extractor interface. */

func (e *ShikiExtractor) GetTimeout() time.Duration {
	return e.config.TickTimeout * time.Millisecond
}

// Anime returns AnimeExtractor.
func (e *ShikiExtractor) Anime() extractor.AnimeExtractor {
	return e.animeEx
}

// Source returns Source.
func (e *ShikiExtractor) Source() extractor.Source {
	return e.source
}
