package collection

import (
	"fmt"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/rotisserie/eris"

	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/postgres"
	"github.com/ilfey/hikilist-go/internal/validator"

	"github.com/ilfey/hikilist-go/data/database"
	animeModels "github.com/ilfey/hikilist-go/data/models/anime"
	collectionModels "github.com/ilfey/hikilist-go/data/models/collection"
	userActionModels "github.com/ilfey/hikilist-go/data/models/userAction"
)

// Контроллер аниме
type CollectionController struct{}

func (c *CollectionController) Create(ctx *handler.Context) {
	user, err := ctx.AuthorizedOnly()
	if err != nil {
		logger.Debug(err)

		ctx.SendJSON(responses.ResponseUnauthorized())

		return
	}

	req := collectionModels.NewCreateModelFromRequest(ctx.Request)

	req.UserID = user.ID

	err = req.Validate()
	if err != nil {
		logger.Debug(err)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": err,
		}))

		return
	}

	err = database.Instance().RunTx(ctx, func(tx *postgres.Transaction) error {
		// Create collection
		sql, args, err := req.InsertSQL()
		if err != nil {
			return err
		}

		err = tx.QueryRow(ctx, sql, args...).Scan(&req.ID)
		if err != nil {
			return eris.Wrap(err, "failed to create collection")
		}

		// Create action
		actionCm := userActionModels.NewCreateCollectionAction(user.ID, req.Title)

		sql, args, err = actionCm.InsertSQL()
		if err != nil {
			return err
		}

		err = tx.QueryRow(ctx, sql, args...).Scan(&actionCm.ID)
		if err != nil {
			return eris.Wrap(err, "failed to create action")
		}

		return nil
	})

	if err != nil {
		logger.Errorf("Failed to create collection: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(responses.ResponseOK())
}

func (controller *CollectionController) List(ctx *handler.Context) {
	paginate := collectionModels.NewPaginateFromQuery(ctx.QueriesMap())

	var lm collectionModels.ListModel

	err := lm.Fill(ctx, paginate, map[string]any{
		"is_public": true,
	})
	if err != nil {
		// Validation error
		var vErr *validator.ValidateError

		if eris.As(err, &vErr) {
			logger.Debug(err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": vErr,
			}))

			return
		}

		logger.Errorf("Failed to get collections: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(&lm)
}

func (controller *CollectionController) Detail(ctx *handler.Context) {
	var userId uint = 0

	user := errorsx.Ignore(ctx.GetUser())
	if user != nil {
		userId = user.ID
	}

	vars := mux.Vars(ctx.Request)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	var dm collectionModels.DetailModel

	err := dm.Get(ctx, fmt.Sprintf("id = %d AND (is_public = TRUE OR user_id = %d)", id, userId))
	if err != nil {
		if eris.Is(err, pgx.ErrNoRows) {
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

func (controller *CollectionController) Animes(ctx *handler.Context) {
	var userId uint = 0

	user := errorsx.Ignore(ctx.GetUser())
	if user != nil {
		userId = user.ID
	}

	vars := mux.Vars(ctx.Request)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	req := animeModels.NewPaginateFromQuery(ctx.QueriesMap())

	var lm animeModels.ListModel

	err := lm.FillFromCollection(
		ctx,
		req,
		userId,
		uint(id),
	)
	if err != nil {
		// Validation error
		var vErr *validator.ValidateError

		if eris.As(err, &vErr) {
			logger.Debug(err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": vErr,
			}))

			return
		}

		logger.Errorf("Failed to get animes: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(&lm)
}
