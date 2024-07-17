package animeController

import (
	"errors"
	"strconv"

	"github.com/gorilla/mux"
	baseController "github.com/ilfey/hikilist-go/api/controllers/base_controller"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	animeService "github.com/ilfey/hikilist-go/services/anime"
	authService "github.com/ilfey/hikilist-go/services/auth"
	"gorm.io/gorm"
)

// Контроллер аниме
type Controller struct {
	*baseController.Controller

	*Dependencies
}

type Dependencies struct {
	Auth  authService.Service
	Anime animeService.Service
}

// Конструктор контроллера
func NewController(deps *Dependencies) *Controller {
	return &Controller{
		Controller: &baseController.Controller{
			AuthService: deps.Auth,
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
	model, tx := controller.Anime.Find()

	if tx.Error != nil {
		ctx.SendJSON(responses.ResponseInternalServerError())
		return
	}

	ctx.SendJSON(model)
}

// Подробная информация об аниме
func (controller *Controller) Detail(ctx *handler.Context) {
	vars := mux.Vars(ctx.Request)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	model, tx := controller.Anime.GetByID(id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			ctx.SendJSON(responses.ResponseNotFound())
			return
		}

		ctx.SendJSON(responses.ResponseInternalServerError())
		return
	}

	ctx.SendJSON(model)
}
