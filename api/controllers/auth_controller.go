package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	authService "github.com/ilfey/hikilist-go/services/auth"
	userService "github.com/ilfey/hikilist-go/services/user"
	"github.com/ilfey/hikilist-go/internal/utils/resx"
)

// Контроллер аутентификации
type AuthController struct {
	authService *authService.Service
	userService *userService.Service
}

// Конструктор контроллера
func NewAuthController(authService *authService.Service, userService *userService.Service) *AuthController {
	return &AuthController{
		authService,
		userService,
	}
}

// Привязка контроллера
func (c *AuthController) Bind(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/auth/login", c.Login).Methods("POST")
	router.HandleFunc("/api/auth/register", c.Register).Methods("POST")
	router.HandleFunc("/api/auth/refresh", c.Refresh).Methods("POST")

	return router
}

// Регистрация
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	req := new(authModels.RegisterModel)

	json.NewDecoder(r.Body).Decode(&req)

	e, ok := req.Validate()
	if !ok {
		resx.NewResponse(400, e).JSON(w)
		return
	}

	userModel, tx := c.authService.CreateUser(req)
	if tx.Error != nil {
		resx.JSON(400, resx.J{
			"detail": "User already exists",
		}).JSON(w)
		return
	}

	tokens := c.authService.GenerateTokens(userModel)

	resx.NewResponse(200, tokens.JSON()).JSON(w)
}

// Аутентификация
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	req := authModels.LoginModelFromRequest(r)

	e, ok := req.Validate()
	if !ok {
		resx.NewResponse(400, e).JSON(w)
		return
	}

	userModel, tx := c.userService.GetByUsername(req.Username)
	if tx.Error != nil {
		resx.ResponseUnauthorized.JSON(w)
		return
	}

	if ok := c.authService.CompareUserPassword(userModel, req.Password); !ok {
		resx.ResponseForbidden.JSON(w)
		return
	}

	tokens := c.authService.GenerateTokens(userModel)

	resx.NewResponse(200, tokens.JSON()).JSON(w)
}

// Обновление токенов аутентификации
func (c *AuthController) Refresh(w http.ResponseWriter, r *http.Request) {
	req := authModels.RefreshModelFromRequest(r)

	claims, ok := c.authService.ParseToken(req.Refresh)
	if !ok {
		resx.ResponseBadRequest.JSON(w)
		return
	}

	userModel, tx := c.userService.GetByID(uint64(claims.UserID))
	if tx.Error != nil {
		resx.ResponseUnauthorized.JSON(w)
		return
	}

	tokens := c.authService.GenerateTokens(userModel)

	resx.NewResponse(200, tokens.JSON()).JSON(w)
}
