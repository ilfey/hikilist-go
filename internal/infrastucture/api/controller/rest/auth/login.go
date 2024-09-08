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

const LoginControllerPath = "/auth/login"

type LoginController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder   builderInterface.Auth
	validator validatorInterface.Auth
	auth      authInterface.Auth
}

func NewLoginController(
	container diInterface.AppContainer,
) (*LoginController, error) {
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

	return &LoginController{
		logger:    log,
		responder: responder,

		auth:      anime,
		builder:   animeBuilder,
		validator: animeValidator,
	}, nil
}

func (c *LoginController) GetLogin(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	loginDTO, err := c.builder.BuildLoginRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Validate dto.
	err = c.validator.ValidateLoginRequestDTO(loginDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Login auth.
	tokens, err := c.auth.Login(r.Context(), loginDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	c.responder.Respond(w, tokens)
}

func (c *LoginController) AddRoute(router *mux.Router) {
	router.
		Path(LoginControllerPath).
		HandlerFunc(c.GetLogin).
		Methods(http.MethodPost)
}
