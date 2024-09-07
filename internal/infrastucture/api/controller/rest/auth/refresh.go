package auth

import (
	"github.com/gorilla/mux"
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	authInterface "github.com/ilfey/hikilist-go/internal/domain/service/auth/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	responderInterface "github.com/ilfey/hikilist-go/internal/domain/service/responder/interface"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"net/http"
)

const RefreshControllerPath = "/auth/refresh"

type RefreshController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder   builderInterface.Auth
	validator validatorInterface.Auth
	auth      authInterface.Auth
}

func NewRefreshController(
	container diInterface.ServiceContainer,
) (*RefreshController, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	responder, err := container.GetResponderService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	anime, err := container.GetAuthService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	animeBuilder, err := container.GetAuthBuilder()
	if err != nil {
		return nil, log.Propagate(err)
	}

	animeValidator, err := container.GetAuthValidator()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &RefreshController{
		logger:    log,
		responder: responder,

		auth:      anime,
		builder:   animeBuilder,
		validator: animeValidator,
	}, nil
}

func (c *RefreshController) GetRefresh(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	refreshDTO, err := c.builder.BuildRefreshRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Validate dto.
	err = c.validator.ValidateRefreshRequestDTO(refreshDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Refresh auth.
	tokens, err := c.auth.Refresh(r.Context(), refreshDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	c.responder.Respond(w, tokens)
}

func (c *RefreshController) AddRoute(router *mux.Router) {
	router.
		Path(RefreshControllerPath).
		HandlerFunc(c.GetRefresh).
		Methods(http.MethodPost)
}
