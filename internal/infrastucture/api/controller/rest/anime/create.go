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

/*
This controller work only in development mode!
*/

const CreateControllerPath = "/animes"

type CreateController struct {
	logger    loggerInterface.Logger
	responder responderInterface.Responder

	builder   builderInterface.Anime
	validator validatorInterface.Anime
	anime     animeInterface.Anime
}

func NewCreateController(
	container diInterface.ServiceContainer,
) (*CreateController, error) {
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

	return &CreateController{
		logger:    log,
		responder: responder,

		anime:     anime,
		builder:   animeBuilder,
		validator: animeValidator,
	}, nil
}

func (c *CreateController) GetCreate(w http.ResponseWriter, r *http.Request) {
	// Build dto.
	createDTO, err := c.builder.BuildCreateRequestDTOFromRequest(r)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Validate dto.
	err = c.validator.ValidateCreateRequestDTO(createDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	// Create anime.
	err = c.anime.Create(r.Context(), createDTO)
	if err != nil {
		c.responder.Respond(w, c.logger.Propagate(err))

		return
	}

	c.responder.Respond(w, map[string]any{
		"detail": "ok",
	})
}

func (c *CreateController) AddRoute(router *mux.Router) {
	router.
		Path(CreateControllerPath).
		HandlerFunc(c.GetCreate).
		Methods(http.MethodPost)
}
