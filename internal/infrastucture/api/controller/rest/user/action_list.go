package user

import (
	"github.com/gorilla/mux"
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	actionInterface "github.com/ilfey/hikilist-go/internal/domain/service/action/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	responderInterface "github.com/ilfey/hikilist-go/internal/domain/service/responder/interface"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"net/http"
)

const ActionListControllerPath = "/users/me/actions"

type ActionListController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder   builderInterface.Action
	validator validatorInterface.Action
	action    actionInterface.Action
}

func NewActionListController(
	container diInterface.ServiceContainer,
) (*ActionListController, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	responder, err := container.GetResponderService()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	action, err := container.GetActionService()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	actionBuilder, err := container.GetActionBuilder()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	actionValidator, err := container.GetActionValidator()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	return &ActionListController{
		logger:    log,
		responder: responder,

		action:    action,
		builder:   actionBuilder,
		validator: actionValidator,
	}, nil
}

func (c *ActionListController) ActionList(w http.ResponseWriter, r *http.Request) {
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
	listDTO, err := c.action.GetListDTO(r.Context(), listRequestDTO, nil)
	if err != nil {
		c.responder.Respond(w, c.logger.LogPropagate(err))
		return
	}

	c.responder.Respond(w, listDTO)
}

func (c *ActionListController) AddRoute(router *mux.Router) {
	router.
		Path(ActionListControllerPath).
		HandlerFunc(c.ActionList).
		Methods(http.MethodGet)
}
