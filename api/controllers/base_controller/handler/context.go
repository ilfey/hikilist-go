package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	authService "github.com/ilfey/hikilist-go/services/auth"
	"github.com/rotisserie/eris"
)

type HandleFunc func(ctx *Context)

type Context struct {
	AuthService authService.Service

	Request *http.Request
	Writer  http.ResponseWriter

	store struct {
		token  *string
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

	// TODO: create goroutine
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
		return nil, eris.Wrap(err, "failed to get claims")
	}

	userModel, err := ctx.AuthService.Authorize(ctx, claims)
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

	user, err := ctx.AuthService.Authorize(ctx, claims)
	if err != nil {
		return nil, err
	}

	// Store user
	ctx.store.user = user

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
	// Check if token already exists
	if ctx.store.token != nil {
		return ctx.store.token
	}

	authorizationPtr := ctx.Get("Authorization")
	if authorizationPtr == nil {
		return nil
	}

	// Check Bearer prefix
	if !strings.HasPrefix(*authorizationPtr, "Bearer ") {
		return nil
	}

	token := strings.TrimPrefix(*authorizationPtr, "Bearer ")

	// Save token in context
	ctx.store.token = &token

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
		return nil, eris.New("token is not set")
	}

	claims, err := ctx.AuthService.ParseToken(*token)
	if err != nil {
		return nil, eris.Wrap(err, "failed to parse token")
	}

	// Save claims in context
	ctx.store.claims = claims

	return claims, nil
}

/*
Implementation context.Context
*/

func (c *Context) hasRequestContext() bool {
	return c.Request != nil && c.Request.Context() != nil
}

func (c *Context) Deadline() (time.Time, bool) {
	if !c.hasRequestContext() {
		return time.Time{}, false
	}

	return c.Request.Context().Deadline()
}

func (c *Context) Done() <-chan struct{} {
	if !c.hasRequestContext() {
		return nil
	}

	return c.Request.Context().Done()
}

func (c *Context) Err() error {
	if !c.hasRequestContext() {
		return nil
	}

	return c.Request.Context().Err()
}

func (c *Context) Value(key any) any {
	if !c.hasRequestContext() {
		return nil
	}

	return c.Request.Context().Value(key)
}
