package userController

import (
	"database/sql"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rotisserie/eris"

	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/validator"

	baseController "github.com/ilfey/hikilist-go/api/controllers/base_controller"

	collectionModels "github.com/ilfey/hikilist-go/data/models/collection"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	userActionModels "github.com/ilfey/hikilist-go/data/models/user_action"

	authService "github.com/ilfey/hikilist-go/services/auth"
)

// Контроллер пользователя
type UserController struct {
	*baseController.Controller
}

// Конструктор контроллера пользователя
func New(
	auth authService.Service,
) *UserController {
	return &UserController{
		Controller: &baseController.Controller{
			AuthService: auth,
		},
	}
}

// Привязка контроллера
func (c *UserController) Bind(router *mux.Router) *mux.Router {
	c.Controller.Bind(router)

	c.HandleFunc("/api/users", c.List).Methods("GET")
	c.HandleFunc("/api/users/{id:[0-9]+}", c.Detail).Methods("GET")
	c.HandleFunc("/api/users/@{username:[a-zA-Z0-9]+}", c.DetailByUsername).Methods("GET")

	c.HandleFunc("/api/users/me", c.Me).Methods("GET")
	c.HandleFunc("/api/users/me/actions", c.MyActions).Methods("GET")
	c.HandleFunc("/api/users/me/collections", c.MyCollections).Methods("GET")

	return router
}

// Список пользователей
func (controller *UserController) List(ctx *handler.Context) {
	paginate := userModels.NewPaginateFromQuery(ctx.QueriesMap())

	var lm userModels.ListModel

	err := lm.Fill(ctx, paginate, nil)
	if err != nil {
		// Validation error
		var vErr *validator.ValidateError

		if eris.As(err, &vErr) {
			logger.Debug(err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": vErr,
			}))

			return
		}

		logger.Errorf("Failed to get users: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(&lm)
}

// Подробная информация о пользователе
func (controller *UserController) Detail(ctx *handler.Context) {
	vars := mux.Vars(ctx.Request)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	var dm userModels.DetailModel

	err := dm.Get(ctx, map[string]any{
		"ID": id,
	})
	if err != nil {
		if eris.Is(err, sql.ErrNoRows) {
			logger.Debug("User not found")

			ctx.SendJSON(responses.ResponseNotFound())

			return
		}

		logger.Errorf("Failed to get user: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(&dm)
}

func (controller *UserController) DetailByUsername(ctx *handler.Context) {
	vars := mux.Vars(ctx.Request)
	username := vars["username"]

	var dm userModels.DetailModel

	err := dm.Get(ctx, map[string]any{
		"Username": username,
	})
	if err != nil {
		if eris.Is(err, sql.ErrNoRows) {
			logger.Debug("User not found")

			ctx.SendJSON(responses.ResponseNotFound())

			return
		}

		logger.Errorf("Failed to get user: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(&dm)
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

func (c *UserController) MyActions(ctx *handler.Context) {
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		logger.Debug("Not authorized")

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	paginate := userActionModels.NewPaginateFromQuery(ctx.QueriesMap())

	var lm userActionModels.ListModel

	err = lm.Fill(ctx, paginate, map[string]any{
		"user_id": user.ID,
	})
	if err != nil {
		// Validation error
		var vErr *validator.ValidateError

		if eris.As(err, &vErr) {
			logger.Debug(err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": vErr,
			}))

			return
		}

		logger.Errorf("Failed to get user actions: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(&lm)
}

func (c *UserController) MyCollections(ctx *handler.Context) {
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		logger.Debug("Not authorized")

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	paginate := collectionModels.NewPaginateFromQuery(ctx.QueriesMap())

	var lm collectionModels.ListModel

	err = lm.Fill(ctx, paginate, map[string]any{
		"user_id": user.ID,
	})
	if err != nil {
		// Validation error
		var vErr *validator.ValidateError

		if eris.As(err, &vErr) {
			logger.Debug(err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": vErr,
			}))

			return
		}

		logger.Errorf("Failed to get user collections: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(&lm)
}
