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

const DeleteControllerPath = "/auth/delete"

type DeleteController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder   builderInterface.Auth
	validator validatorInterface.Auth
	auth      authInterface.Auth
}

func NewDeleteController(
	container diInterface.AppContainer,
) (*DeleteController, error) {
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

	return &DeleteController{
		logger:    log,
		responder: responder,

		auth:      anime,
		builder:   animeBuilder,
		validator: animeValidator,
	}, nil
}

func (c *DeleteController) GetDelete(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	deleteDTO, err := c.builder.BuildDeleteRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Validate dto.
	err = c.validator.ValidateDeleteRequestDTO(deleteDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Delete auth.
	err = c.auth.DeleteUser(r.Context(), deleteDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	c.responder.Respond(w, map[string]any{
		"detail": "ok",
	})
}

func (c *DeleteController) AddRoute(router *mux.Router) {
	router.
		Path(DeleteControllerPath).
		HandlerFunc(c.GetDelete).
		Methods(http.MethodPost)
}
