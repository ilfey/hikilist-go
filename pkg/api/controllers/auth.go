package controllers

import (
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"

	"github.com/ilfey/hikilist-go/internal/postgres"
	"github.com/ilfey/hikilist-go/internal/validator"

	"github.com/ilfey/hikilist-go/pkg/api/handler"
	"github.com/ilfey/hikilist-go/pkg/api/responses"
	"github.com/ilfey/hikilist-go/pkg/models/auth"
	"github.com/ilfey/hikilist-go/pkg/services"
)

type Auth struct {
	Logger logrus.FieldLogger

	Auth services.Auth
	User services.User
}

func (controller *Auth) Register(ctx *handler.Context) {
	registerModel := auth.RegisterModelFromRequest(ctx.Request)

	// Validate register model.
	err := registerModel.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on register model validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate register model %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Register user.
	cm, err := controller.Auth.Register(ctx, registerModel)
	if err != nil {
		// Handle username is already taken error.
		if postgres.PgErrCodeEquals(err, pgerrcode.UniqueViolation) {
			controller.Logger.Infof("Error occurred while user creating with username that is already taken %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": "Username is already taken",
			}))

			return
		}

		controller.Logger.Errorf("Failed to create user %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Generate tokens.
	tokensModel, err := controller.Auth.GenerateTokens(ctx, cm.ID)
	if err != nil {
		controller.Logger.Errorf("Error occurred while generating auth tokens %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError(responses.J{
			"error": "Error generating tokens",
		}))

		return
	}

	// Send success response.
	ctx.SendJSON(tokensModel)
}

func (controller *Auth) Login(ctx *handler.Context) {
	loginModel := auth.LoginModelFromRequest(ctx.Request)

	// Validate login model.
	err := loginModel.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on login model validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate login model %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Get user.
	detailModel, err := controller.User.Get(ctx, map[string]any{
		"Username": loginModel.Username,
	})
	if err != nil {
		// Handle user not found error.
		if eris.Is(err, pgx.ErrNoRows) {
			controller.Logger.Infof("Error occurred while retrieving non-existent user %v", err)

			ctx.SendJSON(responses.ResponseUnauthorized())

			return
		}

		controller.Logger.Errorf("Failed to get user %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Compare passwords.
	if ok := controller.Auth.CompareUserPassword(detailModel, loginModel.Password); !ok {
		ctx.SendJSON(responses.ResponseForbidden())

		return
	}

	// Generate tokens.
	tokensModel, err := controller.Auth.GenerateTokens(ctx, detailModel.ID)
	if err != nil {
		controller.Logger.Errorf("Error occurred while generating auth tokens %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError(responses.J{
			"error": "Error generating tokens",
		}))

		return
	}

	// Send success response.
	ctx.SendJSON(tokensModel)
}

func (controller *Auth) Refresh(ctx *handler.Context) {
	refreshModel := auth.RefreshModelFromRequest(ctx.Request)

	// Validate refresh model.
	err := refreshModel.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on refresh model validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate refresh model %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Refresh tokens.
	tokensModel, err := controller.Auth.RefreshTokens(ctx, refreshModel.Refresh)
	if err != nil {
		controller.Logger.Errorf("Error occurred while generating auth tokens %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError(responses.J{
			"error": "Error generating tokens",
		}))

		return
	}

	// Send success response.
	ctx.SendJSON(tokensModel)
}

func (controller *Auth) Logout(ctx *handler.Context) {
	logoutModel := auth.LogoutModelFromRequest(ctx.Request)

	// Validate logout model.
	err := logoutModel.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on logout model validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate logout model %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Logout.
	err = controller.Auth.Logout(ctx, logoutModel)
	if err != nil {
		// Handle not found token error.
		if eris.Is(err, pgx.ErrNoRows) {
			controller.Logger.Infof("Error occured while retrieving non-existent token %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": "Token already revoked",
			}))
		}

		controller.Logger.Errorf("Failed to delete token %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(responses.ResponseOK())
}

func (controller *Auth) Delete(ctx *handler.Context) {
	// Authorize.
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		controller.Logger.Infof("User is not authorized %v", err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	deleteModel := auth.DeleteModelFromRequest(ctx.Request)

	// Validate delete model.
	err = deleteModel.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on delete model validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate delete model %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Compare password.
	if ok := controller.Auth.CompareUserPassword(user, deleteModel.Password); !ok {
		controller.Logger.Info("Password do not match")

		ctx.SendJSON(responses.ResponseForbidden())

		return
	}

	// Delete user.
	err = controller.Auth.DeleteUser(ctx, user.ID, deleteModel.Refresh)
	if err != nil {
		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(responses.ResponseOK())
}
