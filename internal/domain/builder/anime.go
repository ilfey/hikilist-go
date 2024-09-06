package builder

import (
	"encoding/json"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
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

type AnimeBuilder struct {
	logger loggerInterface.Logger

	extractor extractorInterface.RequestParams
}

func NewAnime(container diInterface.ServiceContainer) (*AnimeBuilder, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	extractor, err := container.GetRequestParametersExtractorService()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	return &AnimeBuilder{
		logger:    log,
		extractor: extractor,
	}, nil
}

func (b *AnimeBuilder) BuildCreateRequestDTOFromRequest(r *http.Request) (*dto.AnimeCreateRequestDTO, error) {
	dto := new(dto.AnimeCreateRequestDTO)

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, b.logger.LogPropagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.logger.LogPropagate(err)
	}

	return dto, nil
}

func (b *AnimeBuilder) BuildDetailRequestDTOFromRequest(r *http.Request) (*dto.AnimeDetailRequestDTO, error) {
	dto := new(dto.AnimeDetailRequestDTO)

	stringId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, b.logger.LogPropagate(err)
	}

	animeId, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		b.logger.Log(err)

		return nil, errtype.NewFieldMustBeIntegerError("id")
	}

	dto.ID = animeId

	return dto, nil
}

func (b *AnimeBuilder) BuildListRequestDTOFromRequest(r *http.Request) (*dto.AnimeListRequestDTO, error) {
	var (
		page  uint64
		limit uint64
		order types.Order
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

	return &dto.AnimeListRequestDTO{
		Page:  page,
		Limit: limit,
		Order: order,
	}, nil
}
