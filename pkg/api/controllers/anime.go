package controllers

import (
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/rotisserie/eris"

	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/validator"

	"github.com/ilfey/hikilist-go/pkg/api/handler"
	"github.com/ilfey/hikilist-go/pkg/api/responses"
	"github.com/ilfey/hikilist-go/pkg/models/anime"
	"github.com/ilfey/hikilist-go/pkg/services"
)

type Anime struct {
	Anime services.Anime
}

// WARN: Not use this method in prodaction mode!
func (controller *Anime) Create(ctx *handler.Context) {
	createModel := anime.CreateModelFromRequest(ctx.Request)

	// Validate create model.
	err := createModel.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			logger.Infof("Error occurred on create model validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		logger.Errorf("Failed to validate create model %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Create anime.
	err = controller.Anime.Create(ctx, createModel)
	if err != nil {
		logger.Errorf("Failed to create anime %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(responses.ResponseOK())
}

func (controller *Anime) List(ctx *handler.Context) {
	paginator := anime.NewPaginator(ctx.QueriesMap())

	// Validate paginator.
	err := paginator.Validate()
	if err != nil {
		// Handle validation error.
		if validator.IsValidateError(err) {
			logger.Infof("Error occurred on paginator validating %v", err)

			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
				"error": err,
			}))

			return
		}

		logger.Errorf("Failed to validate paginator %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Get list model.
	listModel, err := controller.Anime.GetListModel(ctx, paginator, nil)
	if err != nil {
		logger.Errorf("Failed to find animes %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(listModel)
}

func (controller *Anime) Detail(ctx *handler.Context) {
	vars := mux.Vars(ctx.Request)

	// Get id from url vars.
	stringId, ok := vars["id"]
	if !ok {
		logger.Panic("mux.Vars is not contains id")
	}

	// Parsing id.
	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		logger.Warnf("Error occurred on parsing uint from string %v", err)

		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
			"error": "id must be unsigned integer",
		}))

		return
	}

	// Get detail model.
	detailModel, err := controller.Anime.Get(ctx, map[string]any{
		"id": id,
	})
	if err != nil {
		// Handle not found error.
		if eris.Is(err, pgx.ErrNoRows) {
			logger.Infof("Anime not found %v", err)

			ctx.SendJSON(responses.ResponseNotFound())

			return
		}

		logger.Errorf("Failed to get anime %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	// Send success response.
	ctx.SendJSON(&detailModel)
}
