package builder

import (
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
	logger    loggerInterface.Logger
	extractor extractorInterface.RequestParams
}

func NewUser(container diInterface.ServiceContainer) (*UserBuilder, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	extractor, err := container.GetRequestParametersExtractorService()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	return &UserBuilder{
		logger:    log,
		extractor: extractor,
	}, nil
}

//func (b *UserBuilder) BuildCreateRequestDTOFromRequest(r *http.Request) (*user.UserCreateRequestDTO, error) {
//	createRequest := &user.UserCreateRequestDTO{}
//
//	if err := json.NewDecoder(r.Body).Decode(createRequest); err != nil {
//		if errors.Is(err, io.EOF) {
//			return nil, b.logger.LogPropagate(errtype.NewBodyIsEmptyError())
//		}
//
//		return nil, b.logger.LogPropagate(err)
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
		b.logger.Log(err)

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

func (b *UserBuilder) BuildCollectionRequestDTOFromRequest(r *http.Request) (*dto.UserCollectionsRequestDTO, error) {
	var (
		userID uint64
	)

	if id, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		userID = id
	}

	return &dto.UserCollectionsRequestDTO{
		UserID: userID,
	}, nil
}

func (b *UserBuilder) BuildListRequestDTOFromRequest(r *http.Request) (*dto.UserListRequestDTO, error) {
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

	return &dto.UserListRequestDTO{
		Page:  page,
		Limit: limit,
		//Order: order,
	}, nil
}
