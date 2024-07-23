package collectionController

import (
	"errors"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"

	baseController "github.com/ilfey/hikilist-go/api/controllers/base_controller"

	collectionModels "github.com/ilfey/hikilist-go/data/models/collection"
	animeService "github.com/ilfey/hikilist-go/services/anime"
	authService "github.com/ilfey/hikilist-go/services/auth"
	collectionService "github.com/ilfey/hikilist-go/services/collection"
)

// Контроллер аниме
type Controller struct {
	*baseController.Controller

	anime      animeService.Service
	collection collectionService.Service
}

// Конструктор контроллера
func New(
	auth authService.Service,
	// anime animeService.Service,
	collection collectionService.Service,
) *Controller {
	return &Controller{
		Controller: &baseController.Controller{
			AuthService: auth,
		},
		// anime:      anime,
		collection: collection,
	}
}

func (c *Controller) Bind(router *mux.Router) *mux.Router {
	c.Controller.Bind(router)

	c.HandleFunc("/api/collections", c.Create).Methods("POST")
	c.HandleFunc("/api/collections", c.List).Methods("GET")
	c.HandleFunc("/api/collections/{id:[0-9]+}", c.Detail).Methods("GET")
	// c.HandleFunc("/api/collections/{id:[0-9]+}/animes", c.Animes).Methods("GET")

	return router
}

func (c *Controller) Create(ctx *handler.Context) {
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		logger.Debug("Not authorized")

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	req := collectionModels.NewCreateModelFromRequest(ctx.Request)

	vErr := req.Validate()
	if vErr != nil {
		logger.Debugf("Failed to validate create model: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr,
		}))

		return
	}

	req.UserID = user.ID

	model, err := c.collection.Create(ctx, req)
	if err != nil {
		logger.Errorf("Failed to create collection: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(model)
}

func (controller *Controller) List(ctx *handler.Context) {
	paginate := collectionModels.NewPaginateFromQuery(ctx.QueriesMap())

	vErr := paginate.Validate()
	if vErr != nil {
		logger.Debugf("Failed to validate paginate: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr.Error(),
		}))

		return
	}

	model, err := controller.collection.Paginate(ctx, paginate, "is_public = ?", true)
	if err != nil {
		logger.Errorf("Failed to get collections: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(model)
}

func (controller *Controller) Detail(ctx *handler.Context) {
	vars := mux.Vars(ctx.Request)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	model, err := controller.collection.Get(ctx, map[string]any{
		"id":        id,
		"is_public": true,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Debug("Collection not found")

			ctx.SendJSON(responses.ResponseNotFound())

			return
		}

		logger.Errorf("Failed to get anime: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(model)
}

// func (controller *Controller) Animes(ctx *handler.Context) {
// 	vars := mux.Vars(ctx.Request)

// 	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

// 	req := animeModels.NewPaginateFromQuery(ctx.QueriesMap())

// 	vErr := req.Validate()
// 	if vErr != nil {
// 		logger.Debugf("Failed to validate paginate: %v", vErr)

// 		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
// 			"error": vErr.Error(),
// 		}))

// 		return
// 	}

// 	model, err := controller.anime.Paginate(ctx, req, id)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			logger.Debug("Collection not found")

// 			ctx.SendJSON(responses.ResponseNotFound())

// 			return
// 		}

// 		logger.Errorf("Failed to get animes: %v", err)

// 		ctx.SendJSON(responses.ResponseInternalServerError())

// 		return
// 	}

// 	ctx.SendJSON(model)
// }
