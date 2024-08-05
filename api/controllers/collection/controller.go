package collectionController

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rotisserie/eris"

	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"

	baseController "github.com/ilfey/hikilist-go/api/controllers/base_controller"

	animeModels "github.com/ilfey/hikilist-go/data/models/anime"
	collectionModels "github.com/ilfey/hikilist-go/data/models/collection"
	authService "github.com/ilfey/hikilist-go/services/auth"
)

// Контроллер аниме
type Controller struct {
	*baseController.Controller
}

// Конструктор контроллера
func New(
	auth authService.Service,
) *Controller {
	return &Controller{
		Controller: &baseController.Controller{
			AuthService: auth,
		},
	}
}

func (c *Controller) Bind(router *mux.Router) *mux.Router {
	c.Controller.Bind(router)

	c.HandleFunc("/api/collections", c.Create).Methods("POST")
	c.HandleFunc("/api/collections", c.List).Methods("GET")
	c.HandleFunc("/api/collections/{id:[0-9]+}", c.Detail).Methods("GET")
	c.HandleFunc("/api/collections/{id:[0-9]+}/animes", c.Animes).Methods("GET")

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

	err = req.Insert(ctx)
	if err != nil {
		logger.Errorf("Failed to create collection: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(responses.ResponseOK())
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

	var lm collectionModels.ListModel

	err := lm.Paginate(ctx, paginate, map[string]any{
		"IsPublic": true,
	})
	if err != nil {
		logger.Errorf("Failed to get collections: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(&lm)
}

func (controller *Controller) Detail(ctx *handler.Context) {
	var userId uint = 0

	user := errorsx.Ignore(ctx.GetUser())
	if user != nil {
		userId = user.ID
	}

	vars := mux.Vars(ctx.Request)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	var dm collectionModels.DetailModel

	err := dm.Get(ctx, fmt.Sprintf("id = %d AND (is_public = true OR user_id = %d)", id, userId))
	if err != nil {
		if eris.Is(err, sql.ErrNoRows) {
			logger.Debug("Collection not found")

			ctx.SendJSON(responses.ResponseNotFound())

			return
		}

		logger.Errorf("Failed to get anime: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(&dm)
}

func (controller *Controller) Animes(ctx *handler.Context) {
	var userId uint = 0

	user := errorsx.Ignore(ctx.GetUser())
	if user != nil {
		userId = user.ID
	}

	vars := mux.Vars(ctx.Request)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	req := animeModels.NewPaginateFromQuery(ctx.QueriesMap())

	vErr := req.Validate()
	if vErr != nil {
		logger.Debugf("Failed to validate paginate: %v", vErr)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": vErr.Error(),
		}))

		return
	}

	var lm animeModels.ListModel

	// TODO: hide animes from private collections

	err := lm.PaginateFromCollection(
		ctx,
		req,
		userId,
		id,
	)
	if err != nil {
		logger.Errorf("Failed to get animes: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(&lm)
}
