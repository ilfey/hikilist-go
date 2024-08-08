package user

import (
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/rotisserie/eris"

	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/validator"

	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	collectionModels "github.com/ilfey/hikilist-go/data/models/collection"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	userActionModels "github.com/ilfey/hikilist-go/data/models/userAction"
)

// Контроллер пользователя
type UserController struct{}

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
		if eris.Is(err, pgx.ErrNoRows) {
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
		if eris.Is(err, pgx.ErrNoRows) {
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

func (UserController) Collections(ctx *handler.Context) {
	vars := mux.Vars(ctx.Request)

	userId := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	paginate := collectionModels.NewPaginateFromQuery(ctx.QueriesMap())

	var lm collectionModels.ListModel

	err := lm.Fill(ctx, paginate, map[string]any{
		"user_id":   userId,
		"is_public": true,
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

func (UserController) Me(ctx *handler.Context) {
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		logger.Debug(err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	ctx.SendJSON(user)
}

func (UserController) Delete(ctx *handler.Context) {
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		logger.Debug(err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	req := authModels.RefreshModelFromRequest(ctx.Request)

	err = req.Validate()
	if err != nil {
		// Validation error
		var vErr *validator.ValidateError

		if eris.As(err, &vErr) {
			logger.Debug(vErr)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": vErr,
			}))

			return
		}

		logger.Errorf("Failed to validate refresh model: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	err = user.Delete(ctx)
	if err != nil {
		logger.Error("Failed to delete user: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	err = ctx.AuthService.Logout(ctx, req)
	if err != nil {
		logger.Errorf("Failed to logout: %v", err)
	}

	ctx.SendJSON(responses.ResponseOK())
}

func (UserController) MyActions(ctx *handler.Context) {
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		logger.Debug(err)

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

func (UserController) MyCollections(ctx *handler.Context) {
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		logger.Debug(err)

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
