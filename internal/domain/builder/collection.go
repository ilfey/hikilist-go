package builder

import (
	"encoding/json"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/enum"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	extractorInterface "github.com/ilfey/hikilist-go/internal/domain/service/extractor/interface"
	"github.com/ilfey/hikilist-go/internal/domain/types"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
)

type CollectionBuilder struct {
	logger    loggerInterface.Logger
	extractor extractorInterface.RequestParams
}

func NewCollection(container diInterface.ServiceContainer) (*CollectionBuilder, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	extractor, err := container.GetRequestParametersExtractorService()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	return &CollectionBuilder{
		logger:    log,
		extractor: extractor,
	}, nil
}

func (b *CollectionBuilder) BuildUpdateRequestDTOFromRequest(r *http.Request) (*dto.CollectionUpdateRequestDTO, error) {
	updateReqDTO := new(dto.CollectionUpdateRequestDTO)

	if err := json.NewDecoder(r.Body).Decode(updateReqDTO); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, b.logger.LogPropagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.logger.LogPropagate(err)
	}

	stringCollectionId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, b.logger.LogPropagate(err)
	}

	collectionId, err := strconv.ParseUint(stringCollectionId, 10, 64)
	if err != nil {
		return nil, b.logger.LogPropagate(errtype.NewFieldMustBeIntegerError("id"))
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
			return nil, b.logger.LogPropagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.logger.LogPropagate(err)
	}

	stringCollectionId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, b.logger.LogPropagate(err)
	}

	collectionId, err := strconv.ParseUint(stringCollectionId, 10, 64)
	if err != nil {
		return nil, b.logger.LogPropagate(errtype.NewFieldMustBeIntegerError("id"))
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
			return nil, b.logger.LogPropagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.logger.LogPropagate(err)
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
			return nil, b.logger.LogPropagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.logger.LogPropagate(err)
	}

	stringCollectionId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, b.logger.LogPropagate(err)
	}

	collectionId, err := strconv.ParseUint(stringCollectionId, 10, 64)
	if err != nil {
		return nil, b.logger.LogPropagate(errtype.NewFieldMustBeIntegerError("id"))
	}

	addAnimeReqDTO.CollectionID = collectionId

	if userId, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		addAnimeReqDTO.UserID = userId
	}

	return addAnimeReqDTO, nil
}

func (b *CollectionBuilder) BuildListRequestDTOFromRequest(r *http.Request) (*dto.CollectionListRequestDTO, error) {
	var (
		page  uint64
		limit uint64
		//order types.Order
	)

	stringPage, err := b.extractor.GetParameter(r, "page")
	if err != nil {
		limit = 10
	} else {
		page, err = strconv.ParseUint(stringPage, 10, 64)
		if err != nil {
			b.logger.Log(err)

			return nil, errtype.NewFieldMustBeIntegerError("page")
		}
	}

	stringLimit, err := b.extractor.GetParameter(r, "limit")
	if err != nil {
		page = 1
	} else {
		limit, err = strconv.ParseUint(stringLimit, 10, 64)
		if err != nil {
			b.logger.Log(err)

			return nil, errtype.NewFieldMustBeIntegerError("limit")
		}
	}

	//stringOrder, err := b.extractor.GetParameter(r, "order")
	//if err != nil {
	//	order = "-id"
	//} else {
	//	order = types.Order(stringOrder)
	//}

	return &dto.CollectionListRequestDTO{
		Page:  page,
		Limit: limit,
		//Order: order,
	}, nil
}

func (b *CollectionBuilder) BuildAnimeListFromCollectionRequestDTOFromRequest(r *http.Request) (*dto.AnimeListFromCollectionRequestDTO, error) {
	var (
		page         uint64
		limit        uint64
		order        types.Order
		collectionId uint64
		userId       uint64
	)

	stringPage, err := b.extractor.GetParameter(r, "page")
	if err != nil {
		limit = 10
	} else {
		page, err = strconv.ParseUint(stringPage, 10, 64)
		if err != nil {
			b.logger.Log(err)

			return nil, errtype.NewFieldMustBeIntegerError("page")
		}
	}

	stringLimit, err := b.extractor.GetParameter(r, "limit")
	if err != nil {
		page = 1
	} else {
		limit, err = strconv.ParseUint(stringLimit, 10, 64)
		if err != nil {
			b.logger.Log(err)

			return nil, errtype.NewFieldMustBeIntegerError("limit")
		}
	}

	stringOrder, err := b.extractor.GetParameter(r, "order")
	if err != nil {
		order = "-id"
	} else {
		order = types.Order(stringOrder)
	}

	stringCollectionId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, b.logger.LogPropagate(err)
	}

	collectionId, err = strconv.ParseUint(stringCollectionId, 10, 64)
	if err != nil {
		return nil, b.logger.LogPropagate(errtype.NewFieldMustBeIntegerError("id"))
	}

	if id, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		userId = id
	}

	return &dto.AnimeListFromCollectionRequestDTO{
		UserID:       userId,
		CollectionID: collectionId,
		Page:         page,
		Limit:        limit,
		Order:        order,
	}, nil
}

func (b *CollectionBuilder) BuildDetailRequestDTOFromRequest(r *http.Request) (*dto.CollectionDetailRequestDTO, error) {
	detailRequestDTO := new(dto.CollectionDetailRequestDTO)

	stringCollectionId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, b.logger.LogPropagate(err)
	}

	collectionId, err := strconv.ParseUint(stringCollectionId, 10, 64)
	if err != nil {
		return nil, b.logger.LogPropagate(errtype.NewFieldMustBeIntegerError("id"))
	}

	detailRequestDTO.CollectionID = collectionId

	if userId, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		detailRequestDTO.UserID = userId
	}

	return detailRequestDTO, nil
}
