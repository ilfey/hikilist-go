package builder

import (
	"encoding/json"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/enum"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type AuthBuilder struct {
	log loggerInterface.Logger
}

func NewAuth(container diInterface.AppContainer) (*AuthBuilder, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	return &AuthBuilder{
		log: log,
	}, nil
}

func (b *AuthBuilder) BuildDeleteRequestDTOFromRequest(r *http.Request) (*dto.UserDeleteRequestDTO, error) {
	deleteDTO := new(dto.UserDeleteRequestDTO)

	if err := json.NewDecoder(r.Body).Decode(deleteDTO); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, b.log.Propagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.log.Propagate(err)
	}

	if userID, ok := r.Context().Value(enum.UserIDContextKey).(uint64); ok {
		deleteDTO.UserID = userID
	}

	return deleteDTO, nil
}

func (b *AuthBuilder) BuildLoginRequestDTOFromRequest(r *http.Request) (*dto.AuthLoginRequestDTO, error) {
	model := new(dto.AuthLoginRequestDTO)

	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, b.log.Propagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.log.Propagate(err)
	}

	return model, nil
}

func (b *AuthBuilder) BuildLogoutRequestDTOFromRequest(r *http.Request) (*dto.AuthLogoutRequestDTO, error) {
	model := new(dto.AuthLogoutRequestDTO)

	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, b.log.Propagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.log.Propagate(err)
	}

	return model, nil
}

func (b *AuthBuilder) BuildRefreshRequestDTOFromRequest(r *http.Request) (*dto.AuthRefreshRequestDTO, error) {
	model := new(dto.AuthRefreshRequestDTO)

	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, b.log.Propagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.log.Propagate(err)
	}

	return model, nil
}

func (b *AuthBuilder) BuildRegisterRequestDTOFromRequest(r *http.Request) (*dto.AuthRegisterRequestDTO, error) {
	model := new(dto.AuthRegisterRequestDTO)

	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, b.log.Propagate(errtype.NewBodyIsEmptyError())
		}

		return nil, b.log.Propagate(err)
	}

	return model, nil
}
