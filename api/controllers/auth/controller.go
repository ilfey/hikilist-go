package authController

import (
	"encoding/json"

	"github.com/gorilla/mux"
	baseController "github.com/ilfey/hikilist-go/api/controllers/base_controller"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	authService "github.com/ilfey/hikilist-go/services/auth"
	userService "github.com/ilfey/hikilist-go/services/user"
)

// Контроллер аутентификации
type Controller struct {
	*baseController.Controller

	*Dependencies
}

type Dependencies struct {
	Auth authService.Service
	User userService.Service
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

	c.HandleFunc("/api/auth/login", c.Login).Methods("POST")
	c.HandleFunc("/api/auth/register", c.Register).Methods("POST")
	c.HandleFunc("/api/auth/refresh", c.Refresh).Methods("POST")
	c.HandleFunc("/api/auth/logout", c.Logout).Methods("POST")

	return router
}

// Регистрация
func (c *Controller) Register(ctx *handler.Context) {
	registerModel := new(authModels.RegisterModel)

	json.NewDecoder(ctx.Request.Body).Decode(&registerModel)

	vErr := registerModel.Validate()
	if vErr != nil {
		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr,
		}))
		return
	}

	userModel, tx := c.User.Create(registerModel)
	if tx.Error != nil {
		ctx.SendJSON(responses.J{
			"detail": "User already exists",
		}, 400)
		return
	}

	tokensModel, err := c.Auth.GenerateTokens(userModel)
	if err != nil {
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
		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr,
		}))
		return
	}

	userModel, tx := c.User.GetByUsername(req.Username)
	if tx.Error != nil {
		ctx.SendJSON(responses.ResponseUnauthorized())
		return
	}

	if ok := c.Auth.CompareUserPassword(userModel, req.Password); !ok {
		ctx.SendJSON(responses.ResponseForbidden())
		return
	}

	tokensModel, err := c.Auth.GenerateTokens(userModel)
	if err != nil {
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
		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr,
		}))
		return
	}

	claims, err := c.Auth.ParseToken(req.Refresh)
	if err != nil {
		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "Invalid refresh token",
		}))
		return
	}

	userModel, tx := c.User.GetByID(uint64(claims.UserID))
	if tx.Error != nil {
		ctx.SendJSON(responses.ResponseUnauthorized())
		return
	}

	// Delete old token
	if err = c.Auth.DeleteToken(req.Refresh); err != nil {
		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "Token already revoked",
		}))
		return
	}

	// Generate new tokens
	tokensModel, err := c.Auth.GenerateTokens(userModel)
	if err != nil {
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
		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr,
		}))
		return
	}

	_, err := c.Auth.ParseToken(req.Refresh)
	if err != nil {
		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "Invalid refresh token",
		}))
		return
	}

	// Delete old token
	if err = c.Auth.DeleteToken(req.Refresh); err != nil {
		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "Token already revoked",
		}))
		return
	}

	ctx.SendJSON(responses.ResponseOK())
}
