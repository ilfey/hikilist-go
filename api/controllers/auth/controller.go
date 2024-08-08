package auth

import (
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/postgres"
	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/rotisserie/eris"
)

// Контроллер аутентификации
type AuthController struct{}

// Регистрация
func (AuthController) Register(ctx *handler.Context) {
	req := authModels.RegisterModelFromRequest(ctx.Request)

	cm, err := ctx.AuthService.Register(ctx, req)
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

		if postgres.PgErrCodeEquals(err, pgerrcode.UniqueViolation) {
			ctx.SendJSON(responses.ResponseBadRequest())
		}

		logger.Errorf("Failed to create user: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	tokensModel, err := ctx.AuthService.GenerateTokens(ctx, cm.ID)
	if err != nil {
		logger.Debug("Failed to generate tokens: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError(responses.J{
			"error": "Error generating tokens",
		}))

		return
	}

	ctx.SendJSON(tokensModel)
}

// Аутентификация
func (AuthController) Login(ctx *handler.Context) {
	req := authModels.LoginModelFromRequest(ctx.Request)

	vErr := req.Validate()
	if vErr != nil {
		logger.Debugf("Failed to validate request: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr,
		}))

		return
	}

	var dm userModels.DetailModel

	err := dm.Get(ctx, map[string]any{
		"Username": req.Username,
	})
	if err != nil {
		if eris.Is(err, pgx.ErrNoRows) {
			logger.Debug("User not found")

			ctx.SendJSON(responses.ResponseUnauthorized())

			return
		}

		logger.Errorf("Failed to get user: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	if ok := ctx.AuthService.CompareUserPassword(&dm, req.Password); !ok {
		ctx.SendJSON(responses.ResponseForbidden())

		return
	}

	tokensModel, err := ctx.AuthService.GenerateTokens(ctx, dm.ID)
	if err != nil {
		logger.Errorf("Failed to generate tokens: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError(responses.J{
			"error": "Error generating tokens",
		}))

		return
	}

	ctx.SendJSON(tokensModel)
}

// Обновление токенов аутентификации
func (AuthController) Refresh(ctx *handler.Context) {
	req := authModels.RefreshModelFromRequest(ctx.Request)

	vErr := req.Validate()
	if vErr != nil {
		logger.Debugf("Failed to validate request: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr,
		}))

		return
	}

	claims, err := ctx.AuthService.ParseToken(req.Refresh)
	if err != nil {
		logger.Debugf("Failed to parse token: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "Invalid refresh token",
		}))

		return
	}

	// Delete old token
	if err = ctx.AuthService.DeleteToken(ctx, req.Refresh); err != nil {
		logger.Errorf("Failed to delete token: %v", err)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "Token already revoked",
		}))

		return
	}

	// Generate new tokens
	tokensModel, err := ctx.AuthService.GenerateTokens(ctx, claims.UserID)
	if err != nil {
		logger.Debugf("Failed to generate tokens: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError(responses.J{
			"error": "Error generating tokens",
		}))

		return
	}

	ctx.SendJSON(tokensModel)
}

func (AuthController) Logout(ctx *handler.Context) {
	req := authModels.RefreshModelFromRequest(ctx.Request)

	err := ctx.AuthService.Logout(ctx, req)
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

		if eris.Is(err, pgx.ErrNoRows) {
			logger.Debug("Failed to delete token: %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": "Token already revoked",
			}))
		}

		logger.Debugf("Failed to delete token: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(responses.ResponseOK())
}
