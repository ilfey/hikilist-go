package collection

import (
	"github.com/gorilla/mux"
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	animeInterface "github.com/ilfey/hikilist-go/internal/domain/service/anime/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	responderInterface "github.com/ilfey/hikilist-go/internal/domain/service/responder/interface"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"net/http"
)

const AnimeListControllerPath = "/collections/{id:[0-9]+}/animes"

type AnimeListController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder   builderInterface.Collection
	validator validatorInterface.Collection
	anime     animeInterface.Anime
}

func NewAnimeListController(
	container diInterface.ServiceContainer,
) (*AnimeListController, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	responder, err := container.GetResponderService()
	if err != nil {
		return nil, log.LogPropagate(err)
	}

	anime, err := container.GetAnimeService()
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

	return &AnimeListController{
		logger:    log,
		responder: responder,

		anime:     anime,
		builder:   collectionBuilder,
		validator: collectionValidator,
	}, nil
}

func (c *AnimeListController) AnimeList(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	animeListFromCollectionRequestDTO, err := c.builder.BuildAnimeListFromCollectionRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.LogPropagate(err))
		return
	}

	// Validate dto.
	err = c.validator.ValidateAnimeListFromCollectionRequestDTO(animeListFromCollectionRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.LogPropagate(err))
		return
	}

	// Detail listDTO.
	listDTO, err := c.anime.GetFromCollectionListDTO(r.Context(), animeListFromCollectionRequestDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.LogPropagate(err))
		return
	}

	c.responder.Respond(w, listDTO)
}

func (c *AnimeListController) AddRoute(router *mux.Router) {
	router.
		Path(AnimeListControllerPath).
		HandlerFunc(c.AnimeList).
		Methods(http.MethodGet)
}
