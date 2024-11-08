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

const UpdateControllerPath = "/collections/{id:[0-9]+}"

type UpdateController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder    builderInterface.Collection
	validator  validatorInterface.Collection
	collection collectionInterface.Collection
}

func NewUpdateController(
	container diInterface.AppContainer,
) (*UpdateController, error) {
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

	collectionBuilder, err := container.GetCollectionBuilder()
	if err != nil {
		return nil, log.Propagate(err)
	}

	collectionValidator, err := container.GetCollectionValidator()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &UpdateController{
		logger:    log,
		responder: responder,

		collection: collection,
		builder:    collectionBuilder,
		validator:  collectionValidator,
	}, nil
}

func (c *UpdateController) GetUpdate(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	updateRequestDTO, err := c.builder.BuildUpdateRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Validate dto.
	err = c.validator.ValidateUpdateRequestDTO(updateRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Build agg from dto.
	agg, err := c.builder.BuildAggFromUpdateRequestDTO(r.Context(), updateRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))
		return
	}

	// Update collection.
	err = c.collection.Update(r.Context(), agg)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	c.responder.Respond(w, map[string]any{
		"detail": "ok",
	})
}

func (c *UpdateController) AddRoute(router *mux.Router) {
	router.
		Path(UpdateControllerPath).
		HandlerFunc(c.GetUpdate).
		Methods(http.MethodPatch)
}
