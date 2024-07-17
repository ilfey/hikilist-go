package shikiService

import (
	"net/http"
	"time"

	animeModels "github.com/ilfey/hikilist-go/data/models/anime"
	"github.com/ilfey/hikilist-go/internal/httpx"
	"github.com/ilfey/hikilist-go/internal/logger"
	requestbuilder "github.com/ilfey/hikilist-go/internal/request_builder"
	"github.com/ilfey/hikilist-go/parser/shikimori/api/anime"
	shikiConfig "github.com/ilfey/hikilist-go/parser/shikimori/config"
)

type Service struct {
	config *shikiConfig.Config

	Animes anime.IAnimeAPI
}

// Service constructor
func NewShikimoriService(config *shikiConfig.Config) *Service {
	baseBuilder := requestbuilder.NewRequestBuilder(
		config.BaseUrl+"/api/",
		&http.Client{
			Transport: http.DefaultTransport,
			Timeout:   2000 * time.Millisecond,
		},
	)

	baseBuilder.AddResponseHook(responseLoggingHook)
	baseBuilder.AddRequestHook(requestLoggingHook)

	return &Service{
		config: config,
		Animes: anime.NewAnimeAPI(baseBuilder.Sub("animes")),
	}
}

func (s *Service) ParseAnimes(opts ...func(*anime.ListOptions)) ([]*animeModels.CreateModel, error) {

	var models []*animeModels.CreateModel

	var err error

	// Get list
	s.Animes.GetList(opts...).
		Then(func(list []*anime.ListItemModel) {
			models = make([]*animeModels.CreateModel, len(list))

			for i, item := range list {
				s.Animes.GetByID(*item.ID).
					Then(func(anime *anime.AnimeDetailModel) {
						model := animeModels.CreateModel{
							Title:            *anime.Russian,
							Description:      anime.Description,
							Poster:           anime.Image.Original,
							Episodes:         anime.Episodes,
							EpisodesReleased: *anime.EpisodesAired,
							MalID:            anime.MyanimelistID,
							ShikiID:          anime.ID,
						}

						if anime.CompareStatus("released") {
							model.EpisodesReleased = *anime.Episodes
						}

						models[i] = &model
					}).
					Catch(func(_err error) {
						err = _err
					})
			}
		}).
		Catch(func(_err error) {
			err = _err
		})

	if err != nil {
		return nil, err
	}

	return models, nil

}

func requestLoggingHook(rb *httpx.RequestBuilder) {
	logger.Debugf("New request: %s", rb)
}

func responseLoggingHook(r *httpx.Response) {
	logger.Info(r)
}
