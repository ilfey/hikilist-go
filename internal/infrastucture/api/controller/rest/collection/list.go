package collection

import (
	"github.com/gorilla/mux"
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	collectionInterface "github.com/ilfey/hikilist-go/internal/domain/service/collection/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	responderInterface "github.com/ilfey/hikilist-go/internal/domain/service/responder/interface"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"net/http"
)

const ListControllerPath = "/collections"

type ListController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder    builderInterface.Collection
	validator  validatorInterface.Collection
	collection collectionInterface.Collection
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

	collection, err := container.GetCollectionService()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	collectionBuilder, err := container.GetCollectionBuilder()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	collectionValidator, err := container.GetCollectionValidator()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	return &ListController{
		logger:    log,
		responder: responder,

		collection: collection,
		builder:    collectionBuilder,
		validator:  collectionValidator,
	}, nil
}

func (c *ListController) List(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	getListDTO, err := c.builder.BuildListRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.LogPropagate(err))
		return
	}

	// Validate dto.
	err = c.validator.ValidateListRequestDTO(getListDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.LogPropagate(err))
		return
	}

	// Detail listDTO.
	listDTO, err := c.collection.GetListDTO(r.Context(), getListDTO, nil)
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
