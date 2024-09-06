package builder

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/enum"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	extractorInterface "github.com/ilfey/hikilist-go/internal/domain/service/extractor/interface"
	"github.com/ilfey/hikilist-go/internal/domain/types"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"net/http"
	"strconv"
)

type ActionBuilder struct {
	logger    loggerInterface.Logger
	extractor extractorInterface.RequestParams
}

func NewAction(container diInterface.ServiceContainer) (*ActionBuilder, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	extractor, err := container.GetRequestParametersExtractorService()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	return &ActionBuilder{
		logger:    log,
		extractor: extractor,
	}, nil
}

func (b *ActionBuilder) BuildListRequestDTOFromRequest(r *http.Request) (*dto.ActionListRequestDTO, error) {
	var (
		userID uint64
		page   uint64
		limit  uint64
		order  types.Order
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

	if id, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		userID = id
	}

	return &dto.ActionListRequestDTO{
		UserID: userID,
		Page:   page,
		Limit:  limit,
		Order:  order,
	}, nil
}
