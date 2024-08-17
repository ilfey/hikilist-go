package parser

import (
	"context"
	"time"

	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/pkg/models/anime"
	"github.com/ilfey/hikilist-go/pkg/parser/shiki/service"
	"github.com/ilfey/hikilist-go/pkg/services"
	"github.com/rotisserie/eris"
)

var ErrPageEmpty = eris.New("page is empty")

type Parser struct {
	Shiki service.Service
	Anime services.Anime
}

func (p *Parser) Parse(ctx context.Context, page uint) error {
	animes, err := p.Shiki.ParseAnimes(page)
	if err != nil {
		return err
	}

	// Handle empty page.
	if len(animes) == 0 {
		return ErrPageEmpty
	}

	for _, anime := range animes {
		logger.Infof("Saving shikiID %v", *anime.ID)

		// Save anime.
		err := p.saveAnime(ctx, anime)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) saveAnime(ctx context.Context, sdm *anime.ShikiDetailModel) error {
	// Validate shikimori model.
	err := sdm.Validate()
	if err != nil {
		// Handle validate error.
		if validator.IsValidateError(err) {
			logger.Warnf("Error occurred on validating shiki detail model (ShikiId: %d scipped) %v", *sdm.ID, err)

			return nil
		}

		// Unhandled error.
		logger.Errorf("Failed to validate create model %v", err)

		return err
	}

	// Serialize.
	createModel := sdm.ToCreateModel()

	// Validate create model.
	err = createModel.Validate()
	if err != nil {
		// Handle validate error.
		if validator.IsValidateError(err) {
			logger.Warnf("Error occurred on validating create model (ShikiId: %d scipped) %v", *sdm.ID, err)

			return nil
		}

		// Unhandled error.
		logger.Errorf("Failed to validate create model %v", err)

		return err
	}

	// Save in database
	err = p.Anime.Create(ctx, createModel)
	if err != nil {
		// Unhandled error.
		logger.Errorf("Failed to save anime (ShikiId: %d) %v", *sdm.ID, err)

		return err
	}

	return nil
}

// Run parser
func (parser *Parser) Run(ctx context.Context) chan errorsx.Result[uint] {
	var (
		resultCh      = make(chan errorsx.Result[uint], 1)
		page     uint = 1
		err      error
	)

	go func() {
		defer func(page uint, err error) {
			resultCh <- errorsx.Wrap(page, err)
		}(page, err)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				err = parser.Parse(ctx, page)
				if err != nil {
					return
				}

				page++

				time.Sleep(time.Minute)
			}
		}
	}()

	return resultCh
}
