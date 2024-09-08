package builder

import (
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/enum"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	extractorInterface "github.com/ilfey/hikilist-go/internal/domain/service/extractor/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"net/http"
)

type ActionBuilder struct {
	log        loggerInterface.Logger
	extractor  extractorInterface.RequestParams
	pagination builderInterface.Pagination
}

func NewAction(container diInterface.AppContainer) (*ActionBuilder, error) {
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

	return &ActionBuilder{
		log:        log,
		extractor:  extractor,
		pagination: pagination,
	}, nil
}

func (b *ActionBuilder) BuildListRequestDTOFromRequest(r *http.Request) (*dto.ActionListRequestDTO, error) {
	var (
		userID uint64
	)

	pagination, err := b.pagination.BuildPaginationRequestDROFromRequest(r)
	if err != nil {
		return nil, b.log.Propagate(err)
	}

	if id, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		userID = id
	}

	return &dto.ActionListRequestDTO{
		UserID:               userID,
		PaginationRequestDTO: pagination,
	}, nil
}
