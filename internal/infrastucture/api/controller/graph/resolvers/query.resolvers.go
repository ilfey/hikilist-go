package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/infrastucture/api/controller/graph/exec"
)

// Anime is the resolver for the anime field.
func (r *queryResolver) Anime(ctx context.Context, id uint64) (*agg.Anime, error) {
	model, err := r.anime.service.GetByID(ctx, id)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	return model, nil
}

// Page is the resolver for the page field.
func (r *queryResolver) Page(ctx context.Context, page *uint64, limit *uint64) (*agg.Page, error) {
	pagination, err := r.pagination.builder.BuilderPaginationRequestFromPageAndLimit(page, limit)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	if err := r.pagination.validator.Validate(pagination); err != nil {
		return nil, r.log.Propagate(err)
	}

	return &agg.Page{
		Info: &agg.PageInfo{
			Pagination: pagination,
		},
	}, nil
}

// Query returns exec.QueryResolver implementation.
func (r *Resolver) Query() exec.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
