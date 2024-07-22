package userController

import (
	"errors"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"

	baseController "github.com/ilfey/hikilist-go/api/controllers/base_controller"

	userActionModels "github.com/ilfey/hikilist-go/data/models/user_action"

	authService "github.com/ilfey/hikilist-go/services/auth"
	collectionService "github.com/ilfey/hikilist-go/services/collection"
	userService "github.com/ilfey/hikilist-go/services/user"
	userActionService "github.com/ilfey/hikilist-go/services/user_action"
)

// Контроллер пользователя
type UserController struct {
	*baseController.Controller

	collection collectionService.Service
	user       userService.Service
	userAction userActionService.Service
}

// Конструктор контроллера пользователя
func New(
	auth authService.Service,
	collection collectionService.Service,
	user userService.Service,
	userAction userActionService.Service,
) *UserController {
	return &UserController{
		Controller: &baseController.Controller{
			AuthService: auth,
		},

		collection: collection,
		user:       user,
		userAction: userAction,
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
	model, err := controller.user.Find()
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

	model, err := controller.user.Get(id)
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

	model, err := c.userAction.Paginate(paginate, "user_id = ?", user.ID)
	if err != nil {
		logger.Errorf("Failed to get user actions: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(model)
}
