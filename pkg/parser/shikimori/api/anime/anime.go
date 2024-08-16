package anime

import (
	"fmt"

	"github.com/ilfey/hikilist-go/internal/async"
	requestbuilder "github.com/ilfey/hikilist-go/internal/request_builder"
	animeModels "github.com/ilfey/hikilist-go/pkg/models/anime"
)

type IAnimeAPI interface {
	GetList(opts ...ListOption) *async.Promise[[]*animeModels.ShikiListItemModel]
	GetByID(id uint64) *async.Promise[*animeModels.ShikiDetailModel]
}

type AnimeAPI struct {
	builder *requestbuilder.RequestBuilder
}

// Api constructor
func NewAnimeAPI(builder *requestbuilder.RequestBuilder) IAnimeAPI {
	return &AnimeAPI{
		builder: builder,
	}
}

// https://shikimori.one/api/doc/1.0/animes/index
func (a *AnimeAPI) GetList(opts ...ListOption) *async.Promise[[]*animeModels.ShikiListItemModel] {
	options := &ListOptions{
		Page:     1,
		Limit:    50,
		Order:    OrderID,
		Kind:     KindNone,
		Status:   StatusNone,
		Censored: false,
	}

	for _, opt := range opts {
		opt(options)
	}

	return async.New(func() ([]*animeModels.ShikiListItemModel, error) {
		// Get list
		res, err := a.builder.
			Get("/").
			QueryMap(options.ToMap()).
			Async().
			Await()
		if err != nil {
			return nil, err
		}

		// Parse response
		var data []*animeModels.ShikiListItemModel

		res.JSON(&data)

		return data, nil
	})
}

// https://shikimori.one/api/doc/1.0/animes/show
func (a *AnimeAPI) GetByID(id uint64) *async.Promise[*animeModels.ShikiDetailModel] {
	return async.New(func() (*animeModels.ShikiDetailModel, error) {
		res, err := a.builder.
			Get(fmt.Sprintf("/%d", id)).
			Async().
			Await()
		if err != nil {
			return nil, err
		}

		var data animeModels.ShikiDetailModel

		res.JSON(&data)

		return &data, nil
	})
}
