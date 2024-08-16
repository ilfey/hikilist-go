package shikiService

import (
	"net/http"
	"time"

	"github.com/ilfey/hikilist-go/internal/httpx"
	"github.com/ilfey/hikilist-go/internal/logger"
	requestbuilder "github.com/ilfey/hikilist-go/internal/request_builder"
	animeModels "github.com/ilfey/hikilist-go/pkg/models/anime"
	"github.com/ilfey/hikilist-go/pkg/parser/shikimori/api/anime"
	shikiConfig "github.com/ilfey/hikilist-go/pkg/parser/shikimori/shiki"
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

func (s *Service) ParseAnimes(opts ...func(*anime.ListOptions)) ([]*animeModels.ShikiDetailModel, error) {

	var models []*animeModels.ShikiDetailModel

	var err error

	// Get list
	s.Animes.GetList(opts...).
		Then(func(list []*animeModels.ShikiListItemModel) {
			models = make([]*animeModels.ShikiDetailModel, len(list))

			for i, item := range list {
				model, _err := s.Animes.GetByID(*item.ID).Await()
				if _err != nil {
					err = _err
					return
				}

				models[i] = model
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
