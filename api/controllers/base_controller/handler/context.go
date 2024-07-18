package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	authService "github.com/ilfey/hikilist-go/services/auth"
)

type HandleFunc func(ctx *Context)

type Context struct {
	AuthService authService.Service

	Request *http.Request
	Writer  http.ResponseWriter

	store struct {
		claims *authService.Claims
		user   *userModels.DetailModel
	}
}

func NewContext(authService authService.Service, w http.ResponseWriter, r *http.Request) *Context {
	ctx := &Context{
		AuthService: authService,

		Request: r,
		Writer:  w,
	}

	ctx.GetUser()

	return ctx
}

// Get возвращает заголовок запроса.
//
// Возвращает nil, если заголовок не был установлен.
// Если значение заголовка пустое, то возвращается nil.
func (ctx *Context) Get(key string) *string {
	value := ctx.Request.Header.Get(key)
	if value == "" {
		return nil
	}

	return &value
}

func (ctx *Context) Queries(key string) string {
	return ctx.Request.URL.Query().Encode()
}

func (ctx *Context) QueriesMap() map[string][]string {
	return ctx.Request.URL.Query()
}

// Set устанавливает заголовок ответа.
func (ctx *Context) Set(key, value string) {
	ctx.Writer.Header().Set(key, value)
}

// SetType устанавливает тип ответа.
func (ctx *Context) SetType(contentType string) {
	ctx.Writer.Header().Set("Content-Type", contentType)
}

// SetStatus устанавливает статус.
func (ctx *Context) SetStatus(status int) {
	ctx.Writer.WriteHeader(status)
}

// SendJSON отправляет данные в JSON.
func (ctx *Context) SendJSON(data interface{}, code ...int) {

	if len(code) > 0 {
		ctx.SetStatus(code[0])
	} else {
		ctx.SetStatus(http.StatusOK)
	}

	ctx.SetType("application/json")

	ctx.Writer.Write(errorsx.Must(json.Marshal(data)))
}

// AuthorizedOnly проверяет, авторизован ли пользователь.
//
// Возвращает модель пользователя и состояние.
// Возвращает nil и false если пользователь не был авторизован.
func (ctx *Context) AuthorizedOnly() (*userModels.DetailModel, error) {
	claims, err := ctx.GetClaims()
	if err != nil {
		return nil, err
	}

	userModel, err := ctx.AuthService.GetUser(claims)
	if err != nil {
		return nil, err
	}

	return userModel, nil
}

func (ctx *Context) GetUser() (*userModels.DetailModel, error) {
	// Check if user already exists
	if ctx.store.user != nil {
		return ctx.store.user, nil
	}

	claims, err := ctx.GetClaims()
	if err != nil {
		return nil, err
	}

	user, err := ctx.AuthService.GetUser(claims)
	if err != nil {
		return nil, err
	}

	// Store user
	ctx.store.user = user

	err = ctx.AuthService.UpdateUserOnline(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ctx *Context) Authorized() bool {
	_, err := ctx.GetClaims()

	return err == nil
}

// GetToken возвращает токен пользователя без проверки на валидность.
//
// Возвращает nil, если токен не был установлен.
func (ctx *Context) GetToken() *string {
	authorizationPtr := ctx.Get("Authorization")
	if authorizationPtr == nil {
		return nil
	}

	// Check Bearer prefix
	if !strings.HasPrefix(*authorizationPtr, "Bearer ") {
		return nil
	}

	token := strings.TrimPrefix(*authorizationPtr, "Bearer ")

	return &token
}

// GetClaims возвращает Claims пользователя.
//
// Возвращает nil, если токен не был установлен или невалиден.
func (ctx *Context) GetClaims() (*authService.Claims, error) {
	// Check if claims already exists
	if ctx.store.claims != nil {
		return ctx.store.claims, nil
	}

	token := ctx.GetToken()
	if token == nil {
		return nil, fmt.Errorf("token not found")
	}

	claims, err := ctx.AuthService.ParseToken(*token)
	if err != nil {
		return nil, err
	}

	// Save claims in context
	ctx.store.claims = claims

	return claims, nil
}
