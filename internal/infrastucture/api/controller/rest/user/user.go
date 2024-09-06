package user

//
//import (
//	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
//	actionInterface "github.com/ilfey/hikilist-go/internal/domain/service/action/interface"
//	authInterface "github.com/ilfey/hikilist-go/internal/domain/service/auth/interface"
//	collectionInterface "github.com/ilfey/hikilist-go/internal/domain/service/collection/interface"
//	securityInterface "github.com/ilfey/hikilist-go/internal/domain/service/security/interface"
//	"github.com/ilfey/hikilist-go/internal/domain/service/user/interface"
//	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
//	"github.com/ilfey/hikilist-go/internal/infrastucture/api/handler"
//	"github.com/ilfey/hikilist-go/internal/infrastucture/api/responses"
//	"github.com/ilfey/hikilist-go/pkg/postgres"
//	"github.com/pkg/errors"
//	"strconv"
//
//	"github.com/gorilla/mux"
//	"github.com/jackc/pgerrcode"
//	"github.com/jackc/pgx/v5"
//	"github.com/sirupsen/logrus"
//)
//
//// Контроллер пользователя
//type Controller struct {
//	Logger logrus.FieldLogger
//
//	UserBuilder   builderInterface.User
//	UserValidator validatorInterface.User
//
//	Action actionInterface.Action
//
//	AuthBuilder   builderInterface.Auth
//	AuthValidator validatorInterface.Auth
//	Hasher        securityInterface.Hasher
//	Auth          authInterface.Auth
//	Collection    collectionInterface.Collection
//	User          userInterface.User
//}
//
//// Список пользователей
//func (controller *Controller) List(ctx *handler.Context) {
//	paginator := controller.UserBuilder.BuildPaginator(ctx.Request)
//
//	// Validate paginator.
//	err := controller.UserValidator.ValidatePaginator(paginator)
//	if err != nil {
//		controller.Logger.Infof("Error occurred on paginator validating %v", err)
//
//		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
//			"error": err,
//		}))
//
//		return
//	}
//
//	// Detail list model.
//	listModel, err := controller.User.List(ctx, paginator, nil)
//	if err != nil {
//		controller.Logger.Errorf("Failed to get users %v", err)
//
//		ctx.SendJSON(responses.ResponseInternalServerError())
//
//		return
//	}
//
//	// Send success response.
//	ctx.SendJSON(listModel)
//}
//
//// Подробная информация о пользователе
//func (controller *Controller) Detail(ctx *handler.Context) {
//	vars := mux.Vars(ctx.Request)
//
//	// Detail id from url vars.
//	stringId, ok := vars["id"]
//	if !ok {
//		controller.Logger.Panic("mux.Vars is not contains id")
//	}
//
//	// Parsing id.
//	id, err := strconv.ParseUint(stringId, 10, 64)
//	if err != nil {
//		controller.Logger.Warnf("Error occurred on parsing uint64 from string %v", err)
//
//		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
//			"error": "id must be unsigned integer",
//		}))
//
//		return
//	}
//
//	// Detail detail model.
//	detailModel, err := controller.User.Detail(ctx, map[string]any{
//		"CollectionID": id,
//	})
//	if err != nil {
//		// Handle not found.
//		if errors.Is(err, pgx.ErrNoRows) {
//			controller.Logger.Infof("Error occurred while retrieving non-existent user %v", err)
//
//			ctx.SendJSON(responses.ResponseNotFound())
//
//			return
//		}
//
//		controller.Logger.Errorf("Failed to get user %v", err)
//
//		ctx.SendJSON(responses.ResponseInternalServerError())
//
//		return
//	}
//
//	// Send success response.
//	ctx.SendJSON(detailModel)
//}
//
//func (controller *Controller) Collections(ctx *handler.Context) {
//	vars := mux.Vars(ctx.Request)
//
//	// Detail id from url vars.
//	stringId, ok := vars["id"]
//	if !ok {
//		controller.Logger.Panic("mux.Vars is not contains id")
//	}
//
//	// Parsing id.
//	id, err := strconv.ParseUint(stringId, 10, 64)
//	if err != nil {
//		controller.Logger.Warnf("Error occurred on parsing uint64 from string %v", err)
//
//		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
//			"error": "id must be unsigned integer",
//		}))
//
//		return
//	}
//
//	paginator := controller.UserBuilder.BuildPaginator(ctx.Request)
//
//	// Validate paginator.
//	err = controller.UserValidator.ValidatePaginator(paginator)
//	if err != nil {
//		controller.Logger.Infof("Error occurred on paginator validating %v", err)
//
//		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
//			"error": err,
//		}))
//
//		return
//	}
//
//	// Detail list model.
//	listModel, err := controller.Collection.List(ctx, paginator, map[string]any{
//		"user_id":   id,
//		"is_public": true,
//	})
//	if err != nil {
//		controller.Logger.Errorf("Failed to get user collections %v", err)
//
//		ctx.SendJSON(responses.ResponseInternalServerError())
//
//		return
//	}
//
//	// Send success response.
//	ctx.SendJSON(listModel)
//}
//
//func (controller *Controller) Me(ctx *handler.Context) {
//	// Authorize.
//	user, err := ctx.AuthorizedOnly()
//	if err != nil {
//		controller.Logger.Infof("User is not authorized %v", err)
//
//		ctx.SendJSON(responses.ResponseUnauthorized())
//
//		return
//	}
//
//	// Send success response.
//	ctx.SendJSON(user)
//}
//
//func (controller *Controller) ChangePassword(ctx *handler.Context) {
//	// Authorize
//	user, err := ctx.AuthorizedOnly()
//	if err != nil {
//		controller.Logger.Infof("User is not authorized %v", err)
//
//		ctx.SendJSON(responses.ResponseUnauthorized())
//
//		return
//	}
//
//	changePasswordModel, err := controller.AuthBuilder.BuildChangePasswordModel(ctx.Request)
//	if err != nil {
//		controller.Logger.Infof("Error occurred on change password model building %v", err)
//
//		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
//			"error": err,
//		}))
//
//		return
//	}
//
//	// Validate change password model.
//	err = controller.AuthValidator.ValidateChangePasswordRequestDTO(changePasswordModel)
//	if err != nil {
//		controller.Logger.Infof("Error occurred on change password model validating %v", err)
//
//		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
//			"error": err,
//		}))
//
//		return
//	}
//
//	// Compare password.
//	if ok := controller.Hasher.Verify(user, changePasswordModel.OldPassword); !ok {
//		controller.Logger.Info("Passwords do not match")
//
//		ctx.SendJSON(responses.ResponseUnauthorized())
//
//		return
//	}
//
//	// Change password.
//	err = controller.Auth.ChangePassword(ctx, user.UserID, changePasswordModel.NewPassword)
//	if err != nil {
//		controller.Logger.Error("Failed to change password %v", err)
//
//		ctx.SendJSON(responses.ResponseInternalServerError())
//
//		return
//	}
//
//	// Send success response.
//	ctx.SendJSON(responses.ResponseOK())
//}
//
//func (controller *Controller) ChangeUsername(ctx *handler.Context) {
//	// Authorize.
//	user, err := ctx.AuthorizedOnly()
//	if err != nil {
//		controller.Logger.Infof("User is not authorized %v", err)
//
//		ctx.SendJSON(responses.ResponseUnauthorized())
//
//		return
//	}
//
//	changeUsernameModel, err := controller.AuthBuilder.BuildChangeUsernameModel(ctx.Request)
//	if err != nil {
//		controller.Logger.Infof("Error occurred on change username model building %v", err)
//
//		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
//			"error": err,
//		}))
//
//		return
//	}
//
//	// Validate change username model.
//	err = controller.AuthValidator.ValidateChangeUsernameModel(changeUsernameModel)
//	if err != nil {
//		controller.Logger.Infof("Error occurred on change username model validating %v", err)
//
//		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
//			"error": err,
//		}))
//
//		return
//	}
//
//	// Change username.
//	err = controller.User.ChangeUsername(ctx, user.UserID, user.Username, changeUsernameModel.NewUsername)
//	if err != nil {
//		// Handle username is already taken.
//		if postgres.PgErrCodeEquals(err, pgerrcode.UniqueViolation) {
//			controller.Logger.Infof("Error occurred while user updating username that is already taken %v", err)
//
//			ctx.SendJSON(responses.ResponseBadRequest(responses.J{
//				"error": "NewUsername is already taken",
//			}))
//
//			return
//		}
//
//		controller.Logger.Errorf("Failed to change username %v", err)
//
//		ctx.SendJSON(responses.ResponseInternalServerError())
//
//		return
//	}
//
//	// Send success response.
//	ctx.SendJSON(responses.ResponseOK())
//}
//
//func (controller *Controller) MyActions(ctx *handler.Context) {
//	// Authorize
//	user, err := ctx.AuthorizedOnly()
//	if err != nil {
//		controller.Logger.Infof("User is not authorized %v", err)
//
//		ctx.SendJSON(responses.ResponseUnauthorized())
//
//		return
//	}
//
//	paginator := controller.UserBuilder.BuildPaginator(ctx.Request)
//
//	// Validate paginator.
//	err = controller.UserValidator.ValidatePaginator(paginator)
//	if err != nil {
//		controller.Logger.Infof("Error occurred on paginator validating %v", err)
//
//		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
//			"error": err,
//		}))
//
//		return
//	}
//
//	// Detail list model.
//	lm, err := controller.Action.List(ctx, paginator, map[string]any{
//		"user_id": user.UserID,
//	})
//	if err != nil {
//		controller.Logger.Errorf("Failed to get user actions %v", err)
//
//		ctx.SendJSON(responses.ResponseInternalServerError())
//
//		return
//	}
//
//	// Send success response.
//	ctx.SendJSON(lm)
//}
//
//func (controller *Controller) MyCollections(ctx *handler.Context) {
//	// Authorize.
//	user, err := ctx.AuthorizedOnly()
//	if err != nil {
//		controller.Logger.Infof("User is not authorized %v", err)
//
//		ctx.SendJSON(responses.ResponseUnauthorized())
//
//		return
//	}
//
//	paginator := controller.UserBuilder.BuildPaginator(ctx.Request)
//
//	// Validate paginator.
//	err = controller.UserValidator.ValidatePaginator(paginator)
//	if err != nil {
//		controller.Logger.Infof("Error occurred on paginator validating %v", err)
//
//		ctx.SendJSON(responses.ResponseBadRequest(responses.J{
//			"error": err,
//		}))
//
//		return
//	}
//
//	// Detail list model.
//	listModel, err := controller.Collection.List(ctx, paginator, map[string]any{
//		"user_id": user.UserID,
//	})
//	if err != nil {
//		controller.Logger.Errorf("Failed to get user collections %v", err)
//
//		ctx.SendJSON(responses.ResponseInternalServerError())
//
//		return
//	}
//
//	// Send success response.
//	ctx.SendJSON(listModel)
//}
