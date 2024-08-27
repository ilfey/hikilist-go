package extractor

import (
	"time"
)

type Extractor interface {
	GetTimeout() time.Duration

	// Anime returns AnimeExtractor.
	Anime() AnimeExtractor

	// Source returns Source.
	Source() Source
}
