package userController

import (
	"errors"
	"strconv"

	"github.com/gorilla/mux"
	baseController "github.com/ilfey/hikilist-go/api/controllers/base_controller"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	authService "github.com/ilfey/hikilist-go/services/auth"
	userService "github.com/ilfey/hikilist-go/services/user"
	"gorm.io/gorm"
)

// Контроллер пользователя
type UserController struct {
	*baseController.Controller

	*Dependencies
}

type Dependencies struct {
	Auth authService.Service
	User userService.Service
}

// Конструктор контроллера пользователя
func NewController(deps *Dependencies) *UserController {
	return &UserController{
		Controller: &baseController.Controller{
			AuthService: deps.Auth,
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

	return router
}

// Список пользователей
func (controller *UserController) List(ctx *handler.Context) {
	model, tx := controller.User.Find()

	if tx.Error != nil {
		ctx.SendJSON(responses.ResponseInternalServerError())
		return
	}

	ctx.SendJSON(model)
}

// Подробная информация о пользователе
func (controller *UserController) Detail(ctx *handler.Context) {
	vars := mux.Vars(ctx.Request)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	model, tx := controller.User.GetByID(id)
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

func (c *UserController) Me(ctx *handler.Context) {
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		ctx.SendJSON(responses.ResponseUnauthorized())
		return
	}

	ctx.SendJSON(user)
}
