//go:generate go run github.com/99designs/gqlgen generate

package resolvers

import (
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	animeInterface "github.com/ilfey/hikilist-go/internal/domain/service/anime/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type paginationDeps struct {
	builder   builderInterface.Pagination
	validator validatorInterface.Pagination
}

type animeDeps struct {
	builder   builderInterface.Anime
	validator validatorInterface.Anime
	service   animeInterface.Anime
}

type Resolver struct {
	log        loggerInterface.Logger
	pagination paginationDeps
	anime      animeDeps
}

func NewResolver(container diInterface.ServiceContainer) (*Resolver, error) {
	var (
		pagination paginationDeps
		anime      animeDeps
	)

	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	/* ===== Pagination ===== */

	pagination.builder, err = container.GetPaginationBuilder()
	if err != nil {
		return nil, log.Propagate(err)
	}

	pagination.validator, err = container.GetPaginationValidator()
	if err != nil {
		return nil, log.Propagate(err)
	}

	/* ===== Anime ===== */

	anime.builder, err = container.GetAnimeBuilder()
	if err != nil {
		return nil, log.Propagate(err)
	}

	anime.validator, err = container.GetAnimeValidator()
	if err != nil {
		return nil, log.Propagate(err)
	}

	anime.service, err = container.GetAnimeService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &Resolver{
		log:        log,
		pagination: pagination,
		anime:      anime,
	}, nil
}
