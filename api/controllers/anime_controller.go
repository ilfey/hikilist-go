package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	animeModels "github.com/ilfey/hikilist-go/data/models/anime"
	"github.com/ilfey/hikilist-go/internal/utils/errorsx"
	"github.com/ilfey/hikilist-go/internal/utils/resx"
	animeService "github.com/ilfey/hikilist-go/services/anime"
	"gorm.io/gorm"
)

// Контроллер аниме
type AnimeController struct {
	animeService *animeService.Service
}

// Конструктор контроллера
func NewAnimeController(animeService *animeService.Service) *AnimeController {
	return &AnimeController{
		animeService,
	}
}

// Привязка контроллера
func (controller *AnimeController) Bind(router *mux.Router) *mux.Router {
	// router.HandleFunc("/api/animes", controller.Create).Methods("POST")
	router.HandleFunc("/api/animes", controller.List).Methods("GET")
	router.HandleFunc("/api/animes/{id:[0-9]+}", controller.Detail).Methods("GET")

	return router
}

// Создание аниме
func (controller *AnimeController) Create(w http.ResponseWriter, r *http.Request) {
	createModel := animeModels.AnimeCreateModelFromRequest(r)

	if e, ok := createModel.Validate(); !ok {
		resx.NewResponse(400, e).JSON(w)
		return
	}

	model, tx := controller.animeService.Create(createModel)
	if tx.Error != nil {
		resx.ResponseInternalServerError.JSON(w)
		return
	}

	model.Response().JSON(w)
}

// Список аниме
func (controller *AnimeController) List(w http.ResponseWriter, r *http.Request) {
	model, tx := controller.animeService.Find()

	if tx.Error != nil {
		resx.ResponseInternalServerError.JSON(w)
		return
	}

	model.Response().JSON(w)
}

// Подробная информация об аниме
func (controller *AnimeController) Detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	model, tx := controller.animeService.GetByID(id)
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
