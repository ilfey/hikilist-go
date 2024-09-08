package builder

import (
	"encoding/json"
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	extractorInterface "github.com/ilfey/hikilist-go/internal/domain/service/extractor/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
)

type AnimeBuilder struct {
	log loggerInterface.Logger

	extractor  extractorInterface.RequestParams
	pagination builderInterface.Pagination
}

func NewAnime(container diInterface.AppContainer) (*AnimeBuilder, error) {
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

	return &AnimeBuilder{
		log:        log,
		extractor:  extractor,
		pagination: pagination,
	}, nil
}

func (b *AnimeBuilder) BuildCreateRequestDTOFromRequest(r *http.Request) (*dto.AnimeCreateRequestDTO, error) {
	dto := new(dto.AnimeCreateRequestDTO)

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, b.log.Propagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.log.Propagate(err)
	}

	return dto, nil
}

func (b *AnimeBuilder) BuildDetailRequestDTOFromRequest(r *http.Request) (*dto.AnimeDetailRequestDTO, error) {
	dto := new(dto.AnimeDetailRequestDTO)

	stringId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, b.log.Propagate(err)
	}

	animeId, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		b.log.Error(err)

		return nil, errtype.NewFieldMustBeIntegerError("id")
	}

	dto.ID = animeId

	return dto, nil
}

func (b *AnimeBuilder) BuildListRequestDTOFromRequest(r *http.Request) (*dto.AnimeListRequestDTO, error) {
	pagination, err := b.pagination.BuildPaginationRequestDROFromRequest(r)
	if err != nil {
		return nil, b.log.Propagate(err)
	}

	return &dto.AnimeListRequestDTO{
		PaginationRequestDTO: pagination,
	}, nil
}
