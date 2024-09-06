package user

import (
	"github.com/gorilla/mux"
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	responderInterface "github.com/ilfey/hikilist-go/internal/domain/service/responder/interface"
	userInterface "github.com/ilfey/hikilist-go/internal/domain/service/user/interface"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"net/http"
)

const ListControllerPath = "/users"

type ListController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder   builderInterface.User
	validator validatorInterface.User
	user      userInterface.CRUD
}

func NewListController(
	container diInterface.ServiceContainer,
) (*ListController, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	responder, err := container.GetResponderService()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	user, err := container.GetUserService()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	userBuilder, err := container.GetUserBuilder()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	userValidator, err := container.GetUserValidator()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	return &ListController{
		logger:    log,
		responder: responder,

		user:      user,
		builder:   userBuilder,
		validator: userValidator,
	}, nil
}

func (c *ListController) List(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	listRequestDTO, err := c.builder.BuildListRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.LogPropagate(err))
		return
	}

	// Validate dto.
	err = c.validator.ValidateListRequestDTO(listRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.LogPropagate(err))
		return
	}

	// Detail listDTO.
	listDTO, err := c.user.List(r.Context(), listRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.LogPropagate(err))
		return
	}

	c.responder.Respond(w, listDTO)
}

func (c *ListController) AddRoute(router *mux.Router) {
	router.
		Path(ListControllerPath).
		HandlerFunc(c.List).
		Methods(http.MethodGet)
}
