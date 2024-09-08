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

const AddAnimeControllerPath = "/collections/{id:[0-9]+}/animes/add"

type AddAnimeController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder    builderInterface.Collection
	validator  validatorInterface.Collection
	collection collectionInterface.Collection
}

func NewAddAnimeController(
	container diInterface.AppContainer,
) (*AddAnimeController, error) {
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

	return &AddAnimeController{
		logger:    log,
		responder: responder,

		collection: collection,
		builder:    collectionBuilder,
		validator:  collectionValidator,
	}, nil
}

func (c *AddAnimeController) GetAddAnime(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	addAnimeRequestDTO, err := c.builder.BuildAddAnimeRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Validate dto.
	err = c.validator.ValidateAddAnimeRequestDTO(addAnimeRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// AddAnime collection.
	err = c.collection.AddAnimes(r.Context(), addAnimeRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	c.responder.Respond(w, map[string]any{
		"detail": "ok",
	})
}

func (c *AddAnimeController) AddRoute(router *mux.Router) {
	router.
		Path(AddAnimeControllerPath).
		HandlerFunc(c.GetAddAnime).
		Methods(http.MethodPost)
}
