package controllers

import (
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"

	"github.com/ilfey/hikilist-go/internal/postgres"
	"github.com/ilfey/hikilist-go/internal/validator"

	"github.com/ilfey/hikilist-go/pkg/api/handler"
	"github.com/ilfey/hikilist-go/pkg/api/responses"
	"github.com/ilfey/hikilist-go/pkg/models/action"
	"github.com/ilfey/hikilist-go/pkg/models/auth"
	"github.com/ilfey/hikilist-go/pkg/models/collection"
	"github.com/ilfey/hikilist-go/pkg/models/user"
	"github.com/ilfey/hikilist-go/pkg/services"
)

// Контроллер пользователя
type User struct {
	Logger logrus.FieldLogger

	Action     services.Action
	Auth       services.Auth
	Collection services.Collection
	User       services.User
}

// Список пользователей
func (controller *User) List(ctx *handler.Context) {
	paginator := user.NewPaginator(ctx.QueriesMap())

	// Validate paginator.
	err := paginator.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on paginator validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate paginator %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Get list model.
	listModel, err := controller.User.GetListModel(ctx, paginator, nil)
	if err != nil {
		controller.Logger.Errorf("Failed to get users %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(listModel)
}

// Подробная информация о пользователе
func (controller *User) Detail(ctx *handler.Context) {
	vars := mux.Vars(ctx.Request)

	// Get id from url vars.
	stringId, ok := vars["id"]
	if !ok {
		controller.Logger.Panic("mux.Vars is not contains id")
	}

	// Parsing id.
	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		controller.Logger.Warnf("Error occurred on parsing uint from string %v", err)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "id must be unsigned integer",
		}))

		return
	}

	// Get detail model.
	detailModel, err := controller.User.Get(ctx, map[string]any{
		"ID": id,
	})
	if err != nil {
		// Handle not found.
		if eris.Is(err, pgx.ErrNoRows) {
			controller.Logger.Infof("Error occurred while retrieving non-existent user %v", err)

			ctx.SendJSON(responses.ResponseNotFound())

			return
		}

		controller.Logger.Errorf("Failed to get user %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(detailModel)
}

func (controller *User) Collections(ctx *handler.Context) {
	vars := mux.Vars(ctx.Request)

	// Get id from url vars.
	stringId, ok := vars["id"]
	if !ok {
		controller.Logger.Panic("mux.Vars is not contains id")
	}

	// Parsing id.
	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		controller.Logger.Warnf("Error occurred on parsing uint from string %v", err)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "id must be unsigned integer",
		}))

		return
	}

	paginator := collection.NewPaginatorFromQuery(ctx.QueriesMap())

	// Validate paginator.
	err = paginator.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on paginator validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate paginator %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Get list model.
	listModel, err := controller.Collection.GetListModel(ctx, paginator, map[string]any{
		"user_id":   id,
		"is_public": true,
	})
	if err != nil {
		controller.Logger.Errorf("Failed to get user collections %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(listModel)
}

func (controller *User) Me(ctx *handler.Context) {
	// Authorize.
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		controller.Logger.Infof("User is not authorized %v", err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	// Send success response.
	ctx.SendJSON(user)
}

func (controller *User) ChangePassword(ctx *handler.Context) {
	// Authorize
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		controller.Logger.Infof("User is not authorized %v", err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	changePasswordModel := auth.ChangePasswordModelFromRequest(ctx.Request)

	// Validate change password model.
	err = changePasswordModel.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on change password model validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate change password model %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Compare password.
	if ok := controller.Auth.CompareUserPassword(user, changePasswordModel.OldPassword); !ok {
		controller.Logger.Info("Passwords do not match")

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	// Change password.
	err = controller.Auth.ChangePassword(ctx, user.ID, changePasswordModel.NewPassword)
	if err != nil {
		controller.Logger.Error("Failed to change password %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(responses.ResponseOK())
}

func (controller *User) ChangeUsername(ctx *handler.Context) {
	// Authorize.
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		controller.Logger.Infof("User is not authorized %v", err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	changeUsernameModel := auth.ChangeUsernameModelFromRequest(ctx.Request)

	// Validate change username model.
	err = changeUsernameModel.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on change username model validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate change username model %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Change username.
	err = controller.User.ChangeUsername(ctx, user.ID, user.Username, changeUsernameModel.Username)
	if err != nil {
		// Handle username is already taken.
		if postgres.PgErrCodeEquals(err, pgerrcode.UniqueViolation) {
			controller.Logger.Infof("Error occurred while user updating username that is already taken %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": "Username is already taken",
			}))

			return
		}

		controller.Logger.Errorf("Failed to change username %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(responses.ResponseOK())
}

func (controller *User) MyActions(ctx *handler.Context) {
	// Authorize
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		controller.Logger.Infof("User is not authorized %v", err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	paginator := action.NewPaginator(ctx.QueriesMap())

	// Validate paginator.
	err = paginator.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on paginator validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate paginator %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Get list model.
	lm, err := controller.Action.GetListModel(ctx, paginator, map[string]any{
		"user_id": user.ID,
	})
	if err != nil {
		controller.Logger.Errorf("Failed to get user actions %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(lm)
}

func (controller *User) MyCollections(ctx *handler.Context) {
	// Authorize.
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		controller.Logger.Infof("User is not authorized %v", err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	paginator := collection.NewPaginatorFromQuery(ctx.QueriesMap())

	// Validate paginator.
	err = paginator.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on paginator validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate paginator %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Get list model.
	listModel, err := controller.Collection.GetListModel(ctx, paginator, map[string]any{
		"user_id": user.ID,
	})
	if err != nil {
		controller.Logger.Errorf("Failed to get user collections %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(listModel)
}
