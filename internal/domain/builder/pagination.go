package builder

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	extractorInterface "github.com/ilfey/hikilist-go/internal/domain/service/extractor/interface"
	"github.com/ilfey/hikilist-go/internal/domain/types"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"net/http"
	"strconv"
)

type Pagination struct {
	extractor extractorInterface.RequestParams
	log       loggerInterface.Logger
}

func NewPagination(container diInterface.AppContainer) (*Pagination, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	extractor, err := container.GetRequestParametersExtractorService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &Pagination{
		extractor: extractor,
		log:       log,
	}, nil
}

func (b *Pagination) BuildPaginationRequestDROFromRequest(r *http.Request) (*dto.PaginationRequestDTO, error) {
	var (
		page  uint64      = 1
		limit uint64      = 10
		order types.Order = "-id"
	)

	stringPage, err := b.extractor.GetParameter(r, "page")
	if err == nil {
		page, err = strconv.ParseUint(stringPage, 10, 64)
		if err != nil {
			b.log.Error(err)

			return nil, errtype.NewFieldMustBeIntegerError("page")
		}
	}

	stringLimit, err := b.extractor.GetParameter(r, "limit")
	if err == nil {
		limit, err = strconv.ParseUint(stringLimit, 10, 64)
		if err != nil {
			b.log.Error(err)

			return nil, errtype.NewFieldMustBeIntegerError("limit")
		}
	}

	stringOrder, err := b.extractor.GetParameter(r, "order")
	if err == nil {
		order = types.Order(stringOrder)
	}

	return &dto.PaginationRequestDTO{
		Page:  page,
		Limit: limit,
		Order: order,
	}, nil
}
