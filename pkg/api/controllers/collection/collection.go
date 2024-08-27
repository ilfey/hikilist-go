package collection

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"

	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/internal/postgres"
	"github.com/ilfey/hikilist-go/internal/validator"

	"github.com/ilfey/hikilist-go/pkg/api/handler"
	"github.com/ilfey/hikilist-go/pkg/api/responses"
	"github.com/ilfey/hikilist-go/pkg/models/anime"
	animecollection "github.com/ilfey/hikilist-go/pkg/models/anime_collection"
	"github.com/ilfey/hikilist-go/pkg/models/collection"
)

type AnimeProvider interface {
	GetFromCollectionListModel(ctx context.Context, p *paginate.Paginator, userId uint, collectionId uint) (*anime.ListModel, error)
}

type CollectionProvider interface {
	Create(ctx context.Context, cm *collection.CreateModel) error
	Get(ctx context.Context, conds any) (*collection.DetailModel, error)
	GetListModel(ctx context.Context, p *paginate.Paginator, conds any) (*collection.ListModel, error)
	Update(ctx context.Context, um *collection.UpdateModel) error

	AddAnimes(ctx context.Context, aam *animecollection.AddAnimesModel) error
	RemoveAnimes(ctx context.Context, ram *animecollection.RemoveAnimesModel) error
}

// Контроллер аниме
type Controller struct {
	Logger logrus.FieldLogger

	Anime      AnimeProvider
	Collection CollectionProvider
}

func (controller *Controller) Create(ctx *handler.Context) {
	// Authorize.
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		controller.Logger.Infof("User is not authorized %v", err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	createModel := collection.NewCreateModelFromRequest(ctx.Request, user.ID)

	// Validate create model.
	err = createModel.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on create model validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate create model %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Create collection.
	err = controller.Collection.Create(ctx, createModel)
	if err != nil {
		controller.Logger.Errorf("Failed to create collection %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(responses.ResponseOK())
}

func (controller *Controller) List(ctx *handler.Context) {
	paginator := collection.NewPaginatorFromQuery(ctx.QueriesMap())

	// Validate paginator.
	err := paginator.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on paginator validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate paginator %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Get list model.
	listModel, err := controller.Collection.GetListModel(ctx, paginator, map[string]any{
		"is_public": true,
	})
	if err != nil {
		controller.Logger.Errorf("Failed to get collections %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(listModel)
}

func (controller *Controller) Update(ctx *handler.Context) {
	// Authorize.
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		controller.Logger.Infof("User is not authorized %v", err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	vars := mux.Vars(ctx.Request)

	// Get id from url vars.
	stringId, ok := vars["id"]
	if !ok {
		controller.Logger.Panic("mux.Vars is not contains id")
	}

	// Parsing id.
	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		controller.Logger.Warnf("Error occurred on parsing uint from string %v", err)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "id must be unsigned integer",
		}))

		return
	}

	updateModel := collection.NewUpdateModelFromRequest(ctx.Request, user.ID, uint(id))

	// Validate update model.
	err = updateModel.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on update model validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate update model %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Update collection.
	err = controller.Collection.Update(ctx, updateModel)
	if err != nil {
		controller.Logger.Errorf("Failed to update collection %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Return success response.
	ctx.SendJSON(responses.ResponseOK())
}

func (controller *Controller) AddAnimes(ctx *handler.Context) {
	// Authorize.
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		controller.Logger.Infof("User is not authorized %v", err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	vars := mux.Vars(ctx.Request)

	// Get id from url vars.
	stringId, ok := vars["id"]
	if !ok {
		controller.Logger.Panic("mux.Vars is not contains id")
	}

	// Parsing id.
	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		controller.Logger.Warnf("Error occurred on parsing uint from string %v", err)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "id must be unsigned integer",
		}))

		return
	}

	addAnimesModel := animecollection.NewAddAnimesModelFromRequest(ctx.Request, user.ID, uint(id))

	// Validate add animes model.
	err = addAnimesModel.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on add animes model validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate add animes model %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Add animes.
	err = controller.Collection.AddAnimes(ctx, addAnimesModel)
	if err != nil {
		// Handle anime already in collection.
		if postgres.PgErrCodeEquals(err, pgerrcode.UniqueViolation) {
			controller.Logger.Infof("Error occurred while adding a added already anime %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": "Anime added already",
			}))

			return
		}

		// Handle collection non-exists or user is not the creator.
		if eris.Is(err, pgx.ErrNoRows) {
			controller.Logger.Infof("Error occurred while getting a non-existent collection %v", err)

			ctx.SendJSON(responses.ResponseUnauthorized(responses.J{
				"error": "Collection does not exist or user is not the creator",
			}))

			return
		}

		// Handle anime non-exists.
		if postgres.PgErrCodeEquals(err, pgerrcode.ForeignKeyViolation) {
			controller.Logger.Infof("Error occurred while adding a non-existent anime %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": "Anime does not exist",
			}))

			return
		}

		controller.Logger.Errorf("Failed to add animes %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Return success response.
	ctx.SendJSON(responses.ResponseOK())
}

func (controller *Controller) RemoveAnimes(ctx *handler.Context) {
	// Authorize.
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		controller.Logger.Infof("User is not authorized %v", err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	vars := mux.Vars(ctx.Request)

	// Get id from url vars.
	stringId, ok := vars["id"]
	if !ok {
		controller.Logger.Panic("mux.Vars is not contains id")
	}

	// Parsing id.
	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		controller.Logger.Warnf("Error occurred on parsing uint from string %v", err)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "id must be unsigned integer",
		}))

		return
	}

	removeAnimesModel := animecollection.NewRemoveAnimesModelFromRequest(ctx.Request, user.ID, uint(id))

	// Validate remove animes model.
	err = removeAnimesModel.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on remove animes model validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate remove animes model %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Remove animes.
	err = controller.Collection.RemoveAnimes(ctx, removeAnimesModel)
	if err != nil {
		// // Handle anime already in collection.
		// if postgres.PgErrCodeEquals(err, pgerrcode.UniqueViolation) {
		// 	controller.Logger.Infof("Error occurred while adding a added already anime %v", err)

		// 	ctx.SendJSON(responses.ResponseBadRequest(responses.J{
		// 		"error": "Anime added already",
		// 	}))

		// 	return
		// }

		// Handle collection non-exists or user is not the creator.
		if eris.Is(err, pgx.ErrNoRows) {
			controller.Logger.Infof("Error occurred while getting a non-existent collection %v", err)

			ctx.SendJSON(responses.ResponseUnauthorized(responses.J{
				"error": "Collection does not exist or user is not the creator",
			}))

			return
		}

		// // Handle anime non-exists.
		// if postgres.PgErrCodeEquals(err, pgerrcode.ForeignKeyViolation) {
		// 	controller.Logger.Infof("Error occurred while adding a non-existent anime %v", err)

		// 	ctx.SendJSON(responses.ResponseBadRequest(responses.J{
		// 		"error": "Anime does not exist",
		// 	}))

		// 	return
		// }

		controller.Logger.Errorf("Failed to remove animes %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Return success response.
	ctx.SendJSON(responses.ResponseOK())
}

func (controller *Controller) Detail(ctx *handler.Context) {
	var userId uint = 0

	// Authorize.
	user := errorsx.Ignore(ctx.GetUser())
	if user != nil {
		userId = user.ID
	}

	vars := mux.Vars(ctx.Request)

	// Get id from url vars.
	stringId, ok := vars["id"]
	if !ok {
		controller.Logger.Panic("mux.Vars is not contains id")
	}

	// Parsing id.
	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		controller.Logger.Warnf("Error occurred on parsing uint from string %v", err)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "id must be unsigned integer",
		}))

		return
	}

	// Get detail model.
	detailModel, err := controller.Collection.Get(ctx, fmt.Sprintf("id = %d AND (is_public = TRUE OR user_id = %d)", id, userId))
	if err != nil {
		// Handle not found.
		if eris.Is(err, pgx.ErrNoRows) {
			controller.Logger.Infof("Error occurred while getting a non-existent collection %v", err)

			ctx.SendJSON(responses.ResponseNotFound())

			return
		}

		controller.Logger.Errorf("Failed to get anime %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(detailModel)
}

func (controller *Controller) Animes(ctx *handler.Context) {
	var userId uint = 0

	// Authorize.
	user := errorsx.Ignore(ctx.GetUser())
	if user != nil {
		userId = user.ID
	}

	vars := mux.Vars(ctx.Request)

	// Get id from url vars.
	stringId, ok := vars["id"]
	if !ok {
		controller.Logger.Panic("mux.Vars is not contains id")
	}

	// Parsing id.
	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		controller.Logger.Warnf("Error occurred on parsing uint from string %v", err)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "id must be unsigned integer",
		}))

		return
	}

	paginator := anime.NewPaginator(ctx.QueriesMap())

	// Validate paginator.
	err = paginator.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			controller.Logger.Infof("Error occurred on paginator validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		controller.Logger.Errorf("Failed to validate paginator %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Get list model.
	listModel, err := controller.Anime.GetFromCollectionListModel(
		ctx,
		paginator,
		userId,
		uint(id),
	)
	if err != nil {
		controller.Logger.Errorf("Failed to get animes %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(&listModel)
}
