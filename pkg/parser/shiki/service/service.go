package service

import (
	requestbuilder "github.com/ilfey/hikilist-go/internal/request_builder"
	"github.com/ilfey/hikilist-go/pkg/models/anime"
	"github.com/ilfey/hikilist-go/pkg/parser/shiki/repository"
)

type Service interface {
	ParseAnimes(page uint) ([]*anime.ShikiDetailModel, error)
}

type ServiceImpl struct {
	Animes repository.Anime
}

// Service constructor
func New(builder *requestbuilder.RequestBuilder) Service {
	return &ServiceImpl{
		Animes: repository.New(builder),
	}
}

func (s *ServiceImpl) ParseAnimes(page uint) ([]*anime.ShikiDetailModel, error) {
	var models []*anime.ShikiDetailModel

	// Fetch list
	items, err := s.Animes.FetchList(repository.PageOption(page))
	if err != nil {
		return nil, err
	}

	models = make([]*anime.ShikiDetailModel, len(items))

	for i, item := range items {
		models[i], err = s.Animes.FetchByID(*item.ID)
		if err != nil {
			return nil, err
		}
	}

	return models, nil

}
