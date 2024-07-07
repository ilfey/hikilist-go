package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	authService "github.com/ilfey/hikilist-go/services/auth"
	userService "github.com/ilfey/hikilist-go/services/user"
	"github.com/ilfey/hikilist-go/internal/utils/errorsx"
	"github.com/ilfey/hikilist-go/internal/utils/resx"
	"gorm.io/gorm"
)

// Контроллер пользователя
type UserController struct {
	authService *authService.Service
	userService *userService.Service
}

// Конструктор контроллера пользователя
func NewUserController(authService *authService.Service, userService *userService.Service) *UserController {
	return &UserController{
		authService,
		userService,
	}
}

// Привязка контроллера
func (controller *UserController) Bind(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/users", controller.List).Methods("GET")
	router.HandleFunc("/api/users/{id:[0-9]+}", controller.Detail).Methods("GET")

	return router
}

// Список пользователей
func (controller *UserController) List(w http.ResponseWriter, r *http.Request) {
	model, tx := controller.userService.Find()

	if tx.Error != nil {
		resx.ResponseInternalServerError.JSON(w)
		return
	}

	model.Response().JSON(w)
}

// Подробная информация о пользователе
func (controller *UserController) Detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	model, tx := controller.userService.GetByID(id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			resx.ResponseNotFound.JSON(w)
			return
		}

		resx.ResponseInternalServerError.JSON(w)
		return
	}

	model.Response().JSON(w)
}
