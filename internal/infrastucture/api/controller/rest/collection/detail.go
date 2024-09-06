package collection

import (
	"fmt"
	"github.com/gorilla/mux"
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	collectionInterface "github.com/ilfey/hikilist-go/internal/domain/service/collection/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	responderInterface "github.com/ilfey/hikilist-go/internal/domain/service/responder/interface"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"net/http"
)

const DetailControllerPath = "/collections/{id:[0-9]+}"

type DetailController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder    builderInterface.Collection
	validator  validatorInterface.Collection
	collection collectionInterface.Collection
}

func NewDetailController(
	container diInterface.ServiceContainer,
) (*DetailController, error) {
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

	return &DetailController{
		logger:    log,
		responder: responder,

		collection: collection,
		builder:    collectionBuilder,
		validator:  collectionValidator,
	}, nil
}

func (c *DetailController) Detail(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	detailRequestDTO, err := c.builder.BuildDetailRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.LogPropagate(err))

		return
	}

	// Validate dto.
	err = c.validator.ValidateDetailRequestDTO(detailRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.LogPropagate(err))

		return
	}

	// Detail detailDTO.
	detailDTO, err := c.collection.Get(r.Context(),
		fmt.Sprintf(
			"id = %d AND (is_public = TRUE OR user_id = %d)",
			detailRequestDTO.CollectionID, detailRequestDTO.UserID,
		),
	)
	if err != nil {
		c.responder.Respond(w, c.logger.LogPropagate(err))

		return
	}

	c.responder.Respond(w, detailDTO)
}

func (c *DetailController) AddRoute(router *mux.Router) {
	router.
		Path(DetailControllerPath).
		HandlerFunc(c.Detail).
		Methods(http.MethodGet)
}
