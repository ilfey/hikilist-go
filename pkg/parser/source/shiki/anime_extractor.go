package shiki

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/pkg/parser/extractor"
	"github.com/ilfey/hikilist-go/pkg/parser/models"
	"github.com/ilfey/hikilist-go/pkg/parser/source/shiki/config"
	shikiModels "github.com/ilfey/hikilist-go/pkg/parser/source/shiki/models"
	"github.com/rotisserie/eris"
)

type ShikiAnimeExtractor struct {
	config *config.AnimeConfig
	client *resty.Client
}

func newAnimeExtractor(client *resty.Client, config *config.AnimeConfig) *ShikiAnimeExtractor {
	return &ShikiAnimeExtractor{
		config: config,
		client: client,
	}
}

/* Impliment extractor.AnimeExtractor interface. */

// FetchById returns anime detail model.
func (e *ShikiAnimeExtractor) FetchById(id uint) (*models.AnimeDetailModel, error) {
	var result shikiModels.AnimeDetailModel

	_, err := e.client.R().
		SetResult(&result).
		SetPathParams(map[string]string{
			"id": fmt.Sprintf("%d", id),
		}).
		Get("/api/animes/{id}")
	if err != nil {
		return nil, err
	}

	// Validate result.
	err = result.Validate()
	if err != nil {
		return nil, err
	}

	detailModel := models.AnimeDetailModel{
		ID: *result.ID,

		Title:       *result.Russian,
		Description: result.Description,

		Poster: result.Image.Original,

		Episodes:         result.Episodes,
		EpisodesReleased: result.EpisodesAired,

		UpdatedAt: *result.UpdatedAt,
	}

	if result.CompareStatus("released") {
		detailModel.EpisodesReleased = result.Episodes
	}

	return &detailModel, nil
}

/*
FetchList returns anime list model.

https://shikimori.one/api/doc/1.0/animes/index
*/
func (e *ShikiAnimeExtractor) FetchList(page uint) (*models.AnimeListModel, error) {
	var result []*shikiModels.AnimeListItemModel

	_, err := e.client.R().
		SetResult(&result).
		SetQueryParams(map[string]string{
			"page":     fmt.Sprintf("%d", page),
			"limit":    fmt.Sprintf("%d", e.config.Limit),
			"order":    e.config.Order,
			"censored": e.config.Censored,
		}).
		Get("/api/animes")
	if err != nil {
		return nil, err
	}

	var listModel models.AnimeListModel

	listModel.IDs = make([]uint, len(result))

	for i, item := range result {
		// Validate result.
		err = item.Validate()
		if err != nil {
			return nil, err
		}

		listModel.IDs[i] = *item.ID
	}

	return &listModel, nil
}

// Resolve resolves anime detail model.
func (e *ShikiAnimeExtractor) Resolve(ctx context.Context, model *models.AnimeDetailModel) error {
	return nil
}

// OnFetchError handles fetch error.
func (e *ShikiAnimeExtractor) OnFetchError(err error) extractor.Action {
	// Handle validation error.
	if validator.IsValidateError(err) {
		return extractor.ActionSkip
	}

	// Handle network error.
	var netErr *extractor.NetworkError

	if eris.As(err, &netErr) {
		if netErr.StatusCode == 429 { // Too Many Requests.
			// Lock for 1 minute.
			time.Sleep(time.Minute)
		}
	}

	return extractor.ActionStopParser
}

// OnResolveError handles resolve error.
func (e *ShikiAnimeExtractor) OnResolveError(err error) extractor.Action {
	return extractor.ActionStopParser
}
