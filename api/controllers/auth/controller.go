package authController

import (
	"errors"

	"github.com/gorilla/mux"
	baseController "github.com/ilfey/hikilist-go/api/controllers/base_controller"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	"github.com/ilfey/hikilist-go/internal/logger"
	authService "github.com/ilfey/hikilist-go/services/auth"
	userService "github.com/ilfey/hikilist-go/services/user"
	"gorm.io/gorm"
)

// Контроллер аутентификации
type Controller struct {
	*baseController.Controller

	auth authService.Service
	user userService.Service
}

// Конструктор контроллера
func New(
	auth authService.Service,
	user userService.Service,
) *Controller {
	return &Controller{
		Controller: &baseController.Controller{
			AuthService: auth,
		},
		auth: auth,
		user: user,
	}
}

// Привязка контроллера
func (c *Controller) Bind(router *mux.Router) *mux.Router {
	c.Controller.Bind(router)

	c.HandleFunc("/api/auth/login", c.Login).Methods("POST")
	c.HandleFunc("/api/auth/register", c.Register).Methods("POST")
	c.HandleFunc("/api/auth/refresh", c.Refresh).Methods("POST")
	c.HandleFunc("/api/auth/logout", c.Logout).Methods("POST")

	return router
}

// Регистрация
func (c *Controller) Register(ctx *handler.Context) {
	req := authModels.RegisterModelFromRequest(ctx.Request)

	vErr := req.Validate()
	if vErr != nil {
		logger.Debugf("Failed to validate request: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr,
		}))

		return
	}

	userModel, err := c.user.Create(req)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			logger.Debug("User already exists")

			ctx.SendJSON(responses.J{
				"detail": "User already exists",
			}, 400)

			return
		}

		logger.Errorf("Failed to create user: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	tokensModel, err := c.auth.GenerateTokens(userModel)
	if err != nil {
		logger.Debugf("Failed to generate tokens: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError(responses.J{
			"error": "Error generating tokens",
		}))

		return
	}

	ctx.SendJSON(tokensModel)
}

// Аутентификация
func (c *Controller) Login(ctx *handler.Context) {
	req := authModels.LoginModelFromRequest(ctx.Request)

	vErr := req.Validate()
	if vErr != nil {
		logger.Debugf("Failed to validate request: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr,
		}))

		return
	}

	userModel, err := c.user.Get(map[string]any{
		"Username": req.Username,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Debug("User not found")

			ctx.SendJSON(responses.ResponseUnauthorized())

			return
		}

		logger.Errorf("Failed to get user: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	if ok := c.auth.CompareUserPassword(userModel, req.Password); !ok {
		ctx.SendJSON(responses.ResponseForbidden())

		return
	}

	tokensModel, err := c.auth.GenerateTokens(userModel)
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
func (c *Controller) Refresh(ctx *handler.Context) {
	req := authModels.RefreshModelFromRequest(ctx.Request)

	vErr := req.Validate()
	if vErr != nil {
		logger.Debugf("Failed to validate request: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr,
		}))

		return
	}

	claims, err := c.auth.ParseToken(req.Refresh)
	if err != nil {
		logger.Debugf("Failed to parse token: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "Invalid refresh token",
		}))

		return
	}

	userModel, err := c.user.Get(claims.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Debug("User not found")

			ctx.SendJSON(responses.ResponseUnauthorized())
		}

		logger.Errorf("Failed to get user: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Delete old token
	if err = c.auth.DeleteToken(req.Refresh); err != nil {
		logger.Errorf("Failed to delete token: %v", err)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "Token already revoked",
		}))

		return
	}

	// Generate new tokens
	tokensModel, err := c.auth.GenerateTokens(userModel)
	if err != nil {
		logger.Debugf("Failed to generate tokens: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError(responses.J{
			"error": "Error generating tokens",
		}))

		return
	}

	ctx.SendJSON(tokensModel)
}

func (c *Controller) Logout(ctx *handler.Context) {
	req := authModels.RefreshModelFromRequest(ctx.Request)

	vErr := req.Validate()
	if vErr != nil {
		logger.Debugf("Failed to validate request: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr,
		}))

		return
	}

	_, err := c.auth.ParseToken(req.Refresh)
	if err != nil {
		logger.Debugf("Failed to parse token: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "Invalid refresh token",
		}))

		return
	}

	// Delete old token
	if err = c.auth.DeleteToken(req.Refresh); err != nil {
		logger.Debugf("Failed to delete token: %v", err)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "Token already revoked",
		}))

		return
	}

	ctx.SendJSON(responses.ResponseOK())
}
