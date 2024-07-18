package parser

import (
	"time"

	"github.com/ilfey/hikilist-go/internal/logger"
	shikiService "github.com/ilfey/hikilist-go/parser/shikimori"
	"github.com/ilfey/hikilist-go/parser/shikimori/api/anime"
	animeService "github.com/ilfey/hikilist-go/services/anime"
)

type Parser struct {
	Anime animeService.Service
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
			logger.Debugf("Saving shikiID: %v", *anime.ID)

			err := p.Anime.ResolveShiki(anime)
			if err != nil {
				logger.Errorf("Failed to save shikiID: %v, error: %v", *anime.ID, err)
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
