package user

import (
	"github.com/gorilla/mux"
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	responderInterface "github.com/ilfey/hikilist-go/internal/domain/service/responder/interface"
	userInterface "github.com/ilfey/hikilist-go/internal/domain/service/user/interface"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"net/http"
)

const MeControllerPath = "/users/me"

type MeController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder   builderInterface.User
	validator validatorInterface.User
	user      userInterface.CRUD
}

func NewMeController(
	container diInterface.ServiceContainer,
) (*MeController, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	responder, err := container.GetResponderService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	user, err := container.GetUserService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	userBuilder, err := container.GetUserBuilder()
	if err != nil {
		return nil, log.Propagate(err)
	}

	userValidator, err := container.GetUserValidator()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &MeController{
		logger:    log,
		responder: responder,

		user:      user,
		builder:   userBuilder,
		validator: userValidator,
	}, nil
}

func (c *MeController) Me(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	meRequestDTO, err := c.builder.BuildMeRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Validate dto.
	err = c.validator.ValidateMeRequestDTO(meRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Detail detailDTO.
	detailDTO, err := c.user.Detail(r.Context(), &dto.UserDetailRequestDTO{
		UserID: meRequestDTO.UserID,
	})
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	c.responder.Respond(w, detailDTO)
}

func (c *MeController) AddRoute(router *mux.Router) {
	router.
		Path(MeControllerPath).
		HandlerFunc(c.Me).
		Methods(http.MethodGet)
}
