package anime

import (
	"fmt"

	"github.com/ilfey/hikilist-go/internal/async"
	requestbuilder "github.com/ilfey/hikilist-go/internal/request_builder"
)

type IAnimeAPI interface {
	GetList(opts ...ListOption) *async.Promise[[]*ListItemModel]
	GetByID(id uint64) *async.Promise[*AnimeDetailModel]
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
func (a *AnimeAPI) GetList(opts ...ListOption) *async.Promise[[]*ListItemModel] {
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

	return async.New(func() ([]*ListItemModel, error) {
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
		var data []*ListItemModel

		res.JSON(&data)

		return data, nil
	})
}

// https://shikimori.one/api/doc/1.0/animes/show
func (a *AnimeAPI) GetByID(id uint64) *async.Promise[*AnimeDetailModel] {
	return async.New(func() (*AnimeDetailModel, error) {
		res, err := a.builder.
			Get(fmt.Sprintf("/%d", id)).
			Async().
			Await()
		if err != nil {
			return nil, err
		}

		var data AnimeDetailModel

		res.JSON(&data)

		return &data, nil
	})
}
