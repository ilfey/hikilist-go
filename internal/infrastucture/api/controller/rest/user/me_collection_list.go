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

const MeCollectionListControllerPath = "/users/me/collections"

type MeCollectionListController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder    builderInterface.User
	validator  validatorInterface.User
	collection collectionInterface.Collection
}

func NewMeCollectionListController(
	container diInterface.AppContainer,
) (*MeCollectionListController, error) {
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

	return &MeCollectionListController{
		logger:    log,
		responder: responder,

		collection: collection,
		builder:    userBuilder,
		validator:  userValidator,
	}, nil
}

func (c *MeCollectionListController) MeCollectionList(w http.ResponseWriter, r *http.Request) {
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
	listDTO, err := c.collection.GetUserCollectionListDTO(r.Context(), listRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))
		return
	}

	c.responder.Respond(w, listDTO)
}

func (c *MeCollectionListController) AddRoute(router *mux.Router) {
	router.
		Path(MeCollectionListControllerPath).
		HandlerFunc(c.MeCollectionList).
		Methods(http.MethodGet)
}
