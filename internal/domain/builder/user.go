package builder

import (
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/enum"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	extractorInterface "github.com/ilfey/hikilist-go/internal/domain/service/extractor/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"net/http"
	"strconv"
)

type UserBuilder struct {
	log        loggerInterface.Logger
	extractor  extractorInterface.RequestParams
	pagination builderInterface.Pagination
}

func NewUser(container diInterface.AppContainer) (*UserBuilder, error) {
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

	return &UserBuilder{
		log:        log,
		extractor:  extractor,
		pagination: pagination,
	}, nil
}

//func (b *UserBuilder) BuildCreateRequestDTOFromRequest(r *http.Request) (*user.UserCreateRequestDTO, error) {
//	createRequest := &user.UserCreateRequestDTO{}
//
//	if err := json.NewDecoder(r.Body).Decode(createRequest); err != nil {
//		if errors.Is(err, io.EOF) {
//			return nil, b.log.Propagate(errtype.NewBodyIsEmptyError())
//		}
//
//		return nil, b.log.Propagate(err)
//	}
//
//	return createRequest, nil
//}

func (b *UserBuilder) BuildDetailRequestDTOFromRequest(r *http.Request) (*dto.UserDetailRequestDTO, error) {
	stringId, err := b.extractor.GetParameter(r, "id")
	if err != nil {
		return nil, err
	}

	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		b.log.Error(err)

		return nil, errtype.NewFieldMustBeIntegerError("id")
	}

	return &dto.UserDetailRequestDTO{
		UserID: id,
	}, nil
}

func (b *UserBuilder) BuildMeRequestDTOFromRequest(r *http.Request) (*dto.UserMeRequestDTO, error) {
	var (
		userID uint64
	)

	if id, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		userID = id
	}

	return &dto.UserMeRequestDTO{
		UserID: userID,
	}, nil
}

func (b *UserBuilder) BuildCollectionListRequestDTOFromRequest(r *http.Request) (*dto.UserCollectionListRequestDTO, error) {
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

	return &dto.UserCollectionListRequestDTO{
		UserID:               userID,
		PaginationRequestDTO: pagination,
	}, nil
}

func (b *UserBuilder) BuildListRequestDTOFromRequest(r *http.Request) (*dto.UserListRequestDTO, error) {
	pagination, err := b.pagination.BuildPaginationRequestDROFromRequest(r)
	if err != nil {
		return nil, b.log.Propagate(err)
	}

	return &dto.UserListRequestDTO{
		PaginationRequestDTO: pagination,
	}, nil
}
