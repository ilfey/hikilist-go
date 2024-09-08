package anime

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

const DetailControllerPath = "/animes/{id:[0-9]+}"

type DetailController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder   builderInterface.Anime
	validator validatorInterface.Anime
	anime     animeInterface.Anime
}

func NewDetailController(
	container diInterface.AppContainer,
) (*DetailController, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	responder, err := container.GetResponderService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	anime, err := container.GetAnimeService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	animeBuilder, err := container.GetAnimeBuilder()
	if err != nil {
		return nil, log.Propagate(err)
	}

	animeValidator, err := container.GetAnimeValidator()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &DetailController{
		logger:    log,
		responder: responder,

		anime:     anime,
		builder:   animeBuilder,
		validator: animeValidator,
	}, nil
}

func (c *DetailController) GetDetail(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	getDTO, err := c.builder.BuildDetailRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Validate dto.
	err = c.validator.ValidateDetailRequestDTO(getDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Detail detailDTO.
	detailDTO, err := c.anime.Get(r.Context(), map[string]interface{}{
		"id": getDTO.ID,
	})
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	c.responder.Respond(w, detailDTO)
}

func (c *DetailController) AddRoute(router *mux.Router) {
	router.
		Path(DetailControllerPath).
		HandlerFunc(c.GetDetail).
		Methods(http.MethodGet)
}
