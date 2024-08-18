package parser

import (
	"context"
	"time"

	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/pkg/models/anime"
	"github.com/ilfey/hikilist-go/pkg/parser/shiki/service"
	"github.com/ilfey/hikilist-go/pkg/services"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
)

var ErrPageEmpty = eris.New("page is empty")

type Parser struct {
	Logger *logrus.Logger

	Shiki service.Service
	Anime services.Anime
}

func (parser *Parser) Parse(ctx context.Context, page uint) error {
	animes, err := parser.Shiki.ParseAnimes(page)
	if err != nil {
		return err
	}

	// Handle empty page.
	if len(animes) == 0 {
		return ErrPageEmpty
	}

	for _, anime := range animes {
		parser.Logger.Infof("Saving shikiID %v", *anime.ID)

		// Save anime.
		err := parser.saveAnime(ctx, anime)
		if err != nil {
			return err
		}
	}

	return nil
}

func (parser *Parser) saveAnime(ctx context.Context, sdm *anime.ShikiDetailModel) error {
	// Validate shikimori model.
	err := sdm.Validate()
	if err != nil {
		// Handle validate error.
		if validator.IsValidateError(err) {
			parser.Logger.Warnf("Error occurred on validating shiki detail model (ShikiId: %d scipped) %v", *sdm.ID, err)

			return nil
		}

		// Unhandled error.
		parser.Logger.Errorf("Failed to validate create model %v", err)

		return err
	}

	// Serialize.
	createModel := sdm.ToCreateModel()

	// Validate create model.
	err = createModel.Validate()
	if err != nil {
		// Handle validate error.
		if validator.IsValidateError(err) {
			parser.Logger.Warnf("Error occurred on validating create model (ShikiId: %d scipped) %v", *sdm.ID, err)

			return nil
		}

		// Unhandled error.
		parser.Logger.Errorf("Failed to validate create model %v", err)

		return err
	}

	// Save in database
	err = parser.Anime.Create(ctx, createModel)
	if err != nil {
		// Unhandled error.
		parser.Logger.Errorf("Failed to save anime (ShikiId: %d) %v", *sdm.ID, err)

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
