package repository

import (
	requestbuilder "github.com/ilfey/hikilist-go/internal/request_builder"
)

type Shiki interface {
	// TODO: Move to method.
	Anime
}

func New(builder *requestbuilder.RequestBuilder) Shiki {
	return &AnimeImpl{
		builder: builder.Sub("animes"),
	}
}
