package parser

import (
	"time"

	"github.com/ilfey/hikilist-go/internal/logger"
	shikiService "github.com/ilfey/hikilist-go/parser/shikimori"
	"github.com/ilfey/hikilist-go/parser/shikimori/api/anime"
)

type Parser struct {
	Shiki *shikiService.Service
}

func (p *Parser) Parse() (uint64, error) {
	page := uint64(1)

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		// Wait for ticker
		<-ticker.C

		animes, err := p.Shiki.ParseAnimes(anime.PageOption(page))
		if err != nil {
			return page, err
		}

		if len(animes) == 0 {
			break
		}

		for _, anime := range animes {
			logger.Debugf("Resolving shikiID: %v", *anime.ID)

			err := anime.Resolve()
			if err != nil {
				logger.Errorf("Failed to resolve shikiID: %v, error: %v", *anime.ID, err)
			}
		}

		page++
	}

	return page, nil
}

// Run parser
func (parser *Parser) Run() error {
	_, err := parser.Parse()

	return err
}
