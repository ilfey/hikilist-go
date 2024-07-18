package animeController

import (
	"errors"
	"strconv"

	"github.com/gorilla/mux"
	baseController "github.com/ilfey/hikilist-go/api/controllers/base_controller"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	animeModels "github.com/ilfey/hikilist-go/data/models/anime"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"
	animeService "github.com/ilfey/hikilist-go/services/anime"
	authService "github.com/ilfey/hikilist-go/services/auth"
	userActionService "github.com/ilfey/hikilist-go/services/user_action"
	"gorm.io/gorm"
)

// Контроллер аниме
type Controller struct {
	*baseController.Controller

	*Dependencies
}

type Dependencies struct {
	Auth       authService.Service
	UserAction userActionService.Service
	Anime      animeService.Service
}

// Конструктор контроллера
func NewController(deps *Dependencies) *Controller {
	return &Controller{
		Controller: &baseController.Controller{
			AuthService:       deps.Auth,
			UserActionService: deps.UserAction,
		},
		Dependencies: deps,
	}
}

// Привязка контроллера
func (c *Controller) Bind(router *mux.Router) *mux.Router {
	c.Controller.Bind(router)

	c.HandleFunc("/api/animes", c.List).Methods("GET")
	c.HandleFunc("/api/animes/{id:[0-9]+}", c.Detail).Methods("GET")

	return router
}

// Список аниме
func (controller *Controller) List(ctx *handler.Context) {
	paginate := animeModels.NewPaginateFromQuery(ctx.QueriesMap())

	vErr := paginate.Validate()
	if vErr != nil {
		logger.Debugf("Failed to validate paginate: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr.Error(),
		}))

		return
	}

	model, err := controller.Anime.Paginate(paginate)
	if err != nil {
		logger.Errorf("Failed to get animes: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(model)
}

// Подробная информация об аниме
func (controller *Controller) Detail(ctx *handler.Context) {
	vars := mux.Vars(ctx.Request)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	model, err := controller.Anime.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Debug("Anime not found")

			ctx.SendJSON(responses.ResponseNotFound())

			return
		}

		logger.Errorf("Failed to get anime: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(model)
}
