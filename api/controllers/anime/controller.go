package anime

import (
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	animeModels "github.com/ilfey/hikilist-go/data/models/anime"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/jackc/pgx/v5"
	"github.com/rotisserie/eris"
)

// Контроллер аниме
type AnimeController struct{}

func (AnimeController) Create(ctx *handler.Context) {
	req := animeModels.CreateModelFromRequest(ctx.Request)

	err := req.Insert(ctx)
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

		logger.Errorf("Failed to create user: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(responses.ResponseOK())
}

func (AnimeController) List(ctx *handler.Context) {
	paginate := animeModels.NewPaginateFromQuery(ctx.QueriesMap())

	var lm animeModels.ListModel

	err := lm.Fill(ctx, paginate, nil)
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

// Подробная информация об аниме
func (AnimeController) Detail(ctx *handler.Context) {
	vars := mux.Vars(ctx.Request)

	id := errorsx.Must(strconv.ParseUint(vars["id"], 10, 64))

	var dm animeModels.DetailModel

	err := dm.Get(ctx, map[string]any{
		"ID": id,
	})
	if err != nil {
		if eris.Is(err, pgx.ErrNoRows) {
			logger.Debug("Anime not found")

			ctx.SendJSON(responses.ResponseNotFound())

			return
		}

		logger.Errorf("Failed to get anime: %v", err)

		ctx.SendJSON(responses.ResponseInternalServerError())

		return
	}

	ctx.SendJSON(&dm)
}
