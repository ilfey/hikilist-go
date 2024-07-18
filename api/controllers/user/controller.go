package userController

import (
	"errors"
	"strconv"

	"github.com/gorilla/mux"
	baseController "github.com/ilfey/hikilist-go/api/controllers/base_controller"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	userActionModels "github.com/ilfey/hikilist-go/data/models/user_action"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"
	authService "github.com/ilfey/hikilist-go/services/auth"
	userService "github.com/ilfey/hikilist-go/services/user"
	userActionService "github.com/ilfey/hikilist-go/services/user_action"
	"gorm.io/gorm"
)

// Контроллер пользователя
type UserController struct {
	*baseController.Controller

	*Dependencies
}

type Dependencies struct {
	Auth       authService.Service
	UserAction userActionService.Service
	User       userService.Service
}

// Конструктор контроллера пользователя
func NewController(deps *Dependencies) *UserController {
	return &UserController{
		Controller: &baseController.Controller{
			AuthService:       deps.Auth,
			UserActionService: deps.UserAction,
		},
		Dependencies: deps,
	}
}

// Привязка контроллера
func (c *UserController) Bind(router *mux.Router) *mux.Router {
	c.Controller.Bind(router)

	c.HandleFunc("/api/users", c.List).Methods("GET")
	c.HandleFunc("/api/users/{id:[0-9]+}", c.Detail).Methods("GET")
	c.HandleFunc("/api/users/me", c.Me).Methods("GET")
	c.HandleFunc("/api/users/me/actions", c.Actions).Methods("GET")

	return router
}

// Список пользователей
func (controller *UserController) List(ctx *handler.Context) {
	model, err := controller.User.Find()
	if err != nil {
		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(model)
}

// Подробная информация о пользователе
func (controller *UserController) Detail(ctx *handler.Context) {
	vars := mux.Vars(ctx.Request)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	model, err := controller.User.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Debug("User not found")

			ctx.SendJSON(responses.ResponseNotFound())

			return
		}

		logger.Errorf("Failed to get user: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(model)
}

func (c *UserController) Me(ctx *handler.Context) {
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		logger.Debug("Not authorized")

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	ctx.SendJSON(user)
}

func (c *UserController) Actions(ctx *handler.Context) {
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		logger.Debug("Not authorized")

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	paginate := userActionModels.NewPaginateFromQuery(ctx.QueriesMap())

	vErr := paginate.Validate()
	if vErr != nil {
		logger.Debugf("Failed to validate paginate: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr,
		}))

		return
	}

	model, err := c.UserAction.Paginate(paginate, "user_id = ?", user.ID)
	if err != nil {
		logger.Errorf("Failed to get user actions: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(model)
}
