package builder

import (
	"context"
	"encoding/json"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/enum"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	collectionInterface "github.com/ilfey/hikilist-go/internal/domain/service/collection/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	extractorInterface "github.com/ilfey/hikilist-go/internal/domain/service/extractor/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
)

type CollectionBuilder struct {
	log        loggerInterface.Logger
	extractor  extractorInterface.RequestParams
	service    collectionInterface.Collection
	pagination builderInterface.Pagination
}

func NewCollection(container diInterface.AppContainer) (*CollectionBuilder, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	extractor, err := container.GetRequestParametersExtractorService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	pagination, err := container.GetPaginationBuilder()
	if err != nil {
		return nil, log.Propagate(err)
	}

	service, err := container.GetCollectionService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &CollectionBuilder{
		log:        log,
		extractor:  extractor,
		service:    service,
		pagination: pagination,
	}, nil
}

func (b *CollectionBuilder) BuildUpdateRequestDTOFromRequest(r *http.Request) (*dto.CollectionUpdateRequestDTO, error) {
	updateReqDTO := new(dto.CollectionUpdateRequestDTO)

	if err := json.NewDecoder(r.Body).Decode(updateReqDTO); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, b.log.Propagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.log.Propagate(err)
	}

	stringCollectionId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, b.log.Propagate(err)
	}

	collectionId, err := strconv.ParseUint(stringCollectionId, 10, 64)
	if err != nil {
		return nil, b.log.Propagate(errtype.NewFieldMustBeIntegerError("id"))
	}

	updateReqDTO.CollectionID = collectionId

	if userId, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		updateReqDTO.UserID = userId
	}

	return updateReqDTO, nil
}

func (b *CollectionBuilder) BuildRemoveAnimeRequestDTOFromRequest(r *http.Request) (*dto.CollectionRemoveAnimeRequestDTO, error) {
	removeAnimeReqDTO := new(dto.CollectionRemoveAnimeRequestDTO)

	if err := json.NewDecoder(r.Body).Decode(removeAnimeReqDTO); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, b.log.Propagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.log.Propagate(err)
	}

	stringCollectionId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, b.log.Propagate(err)
	}

	collectionId, err := strconv.ParseUint(stringCollectionId, 10, 64)
	if err != nil {
		return nil, b.log.Propagate(errtype.NewFieldMustBeIntegerError("id"))
	}

	removeAnimeReqDTO.CollectionID = collectionId

	if userId, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		removeAnimeReqDTO.UserID = userId
	}

	return removeAnimeReqDTO, nil
}

func (b *CollectionBuilder) BuildCreateRequestDTOFromRequest(r *http.Request) (*dto.CollectionCreateRequestDTO, error) {
	createReqDTO := new(dto.CollectionCreateRequestDTO)

	if err := json.NewDecoder(r.Body).Decode(createReqDTO); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, b.log.Propagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.log.Propagate(err)
	}

	if userId, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		createReqDTO.UserID = userId
	}

	return createReqDTO, nil
}

func (b *CollectionBuilder) BuildAddAnimeRequestDTOFromRequest(r *http.Request) (*dto.CollectionAddAnimeRequestDTO, error) {
	addAnimeReqDTO := new(dto.CollectionAddAnimeRequestDTO)

	if err := json.NewDecoder(r.Body).Decode(addAnimeReqDTO); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, b.log.Propagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.log.Propagate(err)
	}

	stringCollectionId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, b.log.Propagate(err)
	}

	collectionId, err := strconv.ParseUint(stringCollectionId, 10, 64)
	if err != nil {
		return nil, b.log.Propagate(errtype.NewFieldMustBeIntegerError("id"))
	}

	addAnimeReqDTO.CollectionID = collectionId

	if userId, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		addAnimeReqDTO.UserID = userId
	}

	return addAnimeReqDTO, nil
}

func (b *CollectionBuilder) BuildListRequestDTOFromRequest(r *http.Request) (*dto.CollectionListRequestDTO, error) {
	pagination, err := b.pagination.BuildPaginationRequestDROFromRequest(r)
	if err != nil {
		return nil, b.log.Propagate(err)
	}

	return &dto.CollectionListRequestDTO{
		PaginationRequestDTO: pagination,
	}, nil
}

func (b *CollectionBuilder) BuildAnimeListFromCollectionRequestDTOFromRequest(r *http.Request) (*dto.AnimeListFromCollectionRequestDTO, error) {
	var (
		collectionId uint64
		userId       uint64
	)

	pagination, err := b.pagination.BuildPaginationRequestDROFromRequest(r)
	if err != nil {
		return nil, b.log.Propagate(err)
	}

	stringCollectionId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, b.log.Propagate(err)
	}

	collectionId, err = strconv.ParseUint(stringCollectionId, 10, 64)
	if err != nil {
		return nil, b.log.Propagate(errtype.NewFieldMustBeIntegerError("id"))
	}

	if id, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		userId = id
	}

	return &dto.AnimeListFromCollectionRequestDTO{
		UserID:               userId,
		CollectionID:         collectionId,
		PaginationRequestDTO: pagination,
	}, nil
}

func (b *CollectionBuilder) BuildDetailRequestDTOFromRequest(r *http.Request) (*dto.CollectionDetailRequestDTO, error) {
	detailRequestDTO := new(dto.CollectionDetailRequestDTO)

	stringCollectionId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, b.log.Propagate(err)
	}

	collectionId, err := strconv.ParseUint(stringCollectionId, 10, 64)
	if err != nil {
		return nil, b.log.Propagate(errtype.NewFieldMustBeIntegerError("id"))
	}

	detailRequestDTO.CollectionID = collectionId

	if userId, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		detailRequestDTO.UserID = userId
	}

	return detailRequestDTO, nil
}

func (b *CollectionBuilder) BuildAggFromUpdateRequestDTO(ctx context.Context, req *dto.CollectionUpdateRequestDTO) (*agg.CollectionDetail, error) {
	var changes int

	stored, err := b.service.Get(ctx, map[string]any{
		"id": req.CollectionID,
	})
	if err != nil {
		return nil, err
	}

	if stored.UserID != req.UserID {
		return nil, b.log.Propagate(errtype.NewAuthFailedError("you are not the creator of this collection"))
	}

	if req.Title != nil && *req.Title != stored.Title {
		stored.Title = *req.Title
		changes++
	}

	if req.Description != nil && (stored.Description == nil || *req.Description != *stored.Description) {
		stored.Description = req.Description
		changes++
	}

	if req.IsPublic != nil && *req.IsPublic != stored.IsPublic {
		stored.IsPublic = *req.IsPublic
		changes++
	}

	if changes == 0 {
		return nil, b.log.Propagate(errtype.NewBadRequestError("nothing to update"))
	}

	return stored, nil
}
