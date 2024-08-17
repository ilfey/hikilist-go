package repository

import (
	"fmt"

	requestbuilder "github.com/ilfey/hikilist-go/internal/request_builder"
	"github.com/ilfey/hikilist-go/pkg/models/anime"
)

type Anime interface {
	FetchList(opts ...ListOption) ([]*anime.ShikiListItemModel, error)
	FetchByID(id uint64) (*anime.ShikiDetailModel, error)
}

type AnimeImpl struct {
	builder *requestbuilder.RequestBuilder
}

// https://shikimori.one/api/doc/1.0/animes/index
func (a *AnimeImpl) FetchList(opts ...ListOption) ([]*anime.ShikiListItemModel, error) {
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
	var data []*anime.ShikiListItemModel

	err = res.JSON(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// https://shikimori.one/api/doc/1.0/animes/show
func (a *AnimeImpl) FetchByID(id uint64) (*anime.ShikiDetailModel, error) {
	res, err := a.builder.
		Get(fmt.Sprintf("/%d", id)).
		Async().
		Await()
	if err != nil {
		return nil, err
	}

	var data anime.ShikiDetailModel

	err = res.JSON(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
