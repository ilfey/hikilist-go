package user

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

const CollectionListControllerPath = "/users/{id:[0-9]+}/collections"

type CollectionListController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder    builderInterface.User
	validator  validatorInterface.User
	collection collectionInterface.Collection
}

func NewCollectionListController(
	container diInterface.AppContainer,
) (*CollectionListController, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	responder, err := container.GetResponderService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	collection, err := container.GetCollectionService()
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

	return &CollectionListController{
		logger:    log,
		responder: responder,

		collection: collection,
		builder:    userBuilder,
		validator:  userValidator,
	}, nil
}

func (c *CollectionListController) CollectionList(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	listRequestDTO, err := c.builder.BuildCollectionListRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))
		return
	}

	// Validate dto.
	err = c.validator.ValidateCollectionListRequestDTO(listRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))
		return
	}

	// Detail listDTO.
	listDTO, err := c.collection.GetUserPublicCollectionList(r.Context(), listRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))
		return
	}

	c.responder.Respond(w, listDTO)
}

func (c *CollectionListController) AddRoute(router *mux.Router) {
	router.
		Path(CollectionListControllerPath).
		HandlerFunc(c.CollectionList).
		Methods(http.MethodGet)
}
