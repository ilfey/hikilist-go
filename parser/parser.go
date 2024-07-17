package parser

import (
	"github.com/ilfey/hikilist-go/internal/logger"
	shikiService "github.com/ilfey/hikilist-go/parser/shikimori"
	animeService "github.com/ilfey/hikilist-go/services/anime"
)

type Parser struct {
	Anime animeService.Service
	Shiki *shikiService.Service
}

func (p *Parser) Parse() error {

	animes, err := p.Shiki.ParseAnimes()
	if err != nil {
		return err
	}

	for _, anime := range animes {
		logger.Debugf("Saving shikiID: %v", *anime.ShikiID)

		_, tx := p.Anime.Create(anime)
		if tx.Error != nil {
			logger.Errorf("Failed to save shikiID: %v, error: %v", *anime.ShikiID, tx.Error)
		}
	}

	return nil
}

// Run parser
func (parser *Parser) Run() error {
	return parser.Parse()
}
