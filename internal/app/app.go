package app

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/config"
	"github.com/ilfey/hikilist-go/internal/domain/builder"
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	actionInterface "github.com/ilfey/hikilist-go/internal/domain/service/action/interface"
	animeInterface "github.com/ilfey/hikilist-go/internal/domain/service/anime/interface"
	authInterface "github.com/ilfey/hikilist-go/internal/domain/service/auth/interface"
	collectionInterface "github.com/ilfey/hikilist-go/internal/domain/service/collection/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/extractor"
	extractorInterface "github.com/ilfey/hikilist-go/internal/domain/service/extractor/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/responder"
	responderInterface "github.com/ilfey/hikilist-go/internal/domain/service/responder/interface"
	securityInterface "github.com/ilfey/hikilist-go/internal/domain/service/security/interface"
	tokenizerInterface "github.com/ilfey/hikilist-go/internal/domain/service/tokenizer/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/user"
	userInterface "github.com/ilfey/hikilist-go/internal/domain/service/user/interface"
	"github.com/ilfey/hikilist-go/internal/domain/validator"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	"github.com/ilfey/hikilist-go/internal/infrastucture/api/controller"
	animeController "github.com/ilfey/hikilist-go/internal/infrastucture/api/controller/rest/anime"
	authController "github.com/ilfey/hikilist-go/internal/infrastucture/api/controller/rest/auth"
	collectionController "github.com/ilfey/hikilist-go/internal/infrastucture/api/controller/rest/collection"
	userController "github.com/ilfey/hikilist-go/internal/infrastucture/api/controller/rest/user"
	"github.com/ilfey/hikilist-go/internal/infrastucture/repositories"
	"github.com/ilfey/hikilist-go/internal/infrastucture/server"
	"github.com/ilfey/hikilist-go/internal/infrastucture/service/action"
	"github.com/ilfey/hikilist-go/internal/infrastucture/service/anime"
	"github.com/ilfey/hikilist-go/internal/infrastucture/service/auth"
	"github.com/ilfey/hikilist-go/internal/infrastucture/service/collection"
	"github.com/ilfey/hikilist-go/internal/infrastucture/service/security"
	"github.com/ilfey/hikilist-go/internal/infrastucture/service/tokenizer"
	"github.com/ilfey/hikilist-go/internal/providers/database"
	"github.com/ilfey/hikilist-go/pkg/logger"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/postgres"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type App struct {
	config    *config.AppConfig
	container diInterface.AppContainer
}

func NewApp(
	config *config.AppConfig,
	container diInterface.AppContainer,
) *App {
	return &App{
		config:    config,
		container: container,
	}
}

func (a *App) Run(mWg *sync.WaitGroup) {
	defer mWg.Done()

	// AppContext.
	cancel := a.InitAppContext()

	// AppConfig.
	if err := a.InitAppConfig(); err != nil {
		panic(err)
	}

	// Logger.
	log, logCancel, err := a.InitLogger()
	if err != nil {
		panic(err)
	}
	defer logCancel()

	wg := &sync.WaitGroup{}
	defer func() {
		cancel()
		wg.Wait()
		time.Sleep(time.Second)
	}()

	// Postgres database.
	postgresCancel, err := a.InitPostgres()
	if err != nil {
		log.Critical(err)
		return
	}
	defer postgresCancel()

	// ReqRes.
	if err := a.InitReqRes(); err != nil {
		log.Critical(err)
		return
	}

	// Pagination.
	if err := a.InitPagination(); err != nil {
		log.Critical(err)
		return
	}

	// Action.
	if err := a.InitAction(); err != nil {
		log.Critical(err)
		return
	}

	// Anime.
	if err := a.InitAnime(); err != nil {
		log.Critical(err)
		return
	}

	// Collection.
	if err := a.InitCollection(); err != nil {
		log.Critical(err)
		return
	}

	// Token.
	if err := a.InitToken(); err != nil {
		log.Critical(err)
		return
	}

	// User.
	if err := a.InitUser(); err != nil {
		log.Critical(err)
		return
	}

	// Auth.
	if err := a.InitAuth(); err != nil {
		log.Critical(err)
		return
	}

	if err := a.InitHttpServer(wg); err != nil {
		log.Critical(err)
		return
	}

	<-a.shutdown()
}

func (a *App) shutdown() chan os.Signal {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	return stopCh
}

func (a *App) InitAuthedControllers() ([]controller.Controller, error) {
	log, err := a.container.GetLogger()
	if err != nil {
		return nil, err
	}

	// Auth.

	logoutController, err := authController.NewLogoutController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	deleteUserController, err := authController.NewDeleteController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	// Collection.

	createCollectionController, err := collectionController.NewCreateController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	updateCollectionController, err := collectionController.NewUpdateController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	addAnimeCollectionController, err := collectionController.NewAddAnimeController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	removeAnimeCollectionController, err := collectionController.NewRemoveAnimeController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	deleteCollectionController, err := collectionController.NewDeleteController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	// User.

	meController, err := userController.NewMeController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	actionListController, err := userController.NewActionListController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	collectionListController, err := userController.NewMeCollectionListController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	return []controller.Controller{
		// Auth.
		logoutController,
		deleteUserController,

		// Collection.
		createCollectionController,
		updateCollectionController,
		addAnimeCollectionController,
		removeAnimeCollectionController,
		deleteCollectionController,

		// User.
		meController,
		actionListController,
		collectionListController,
	}, nil
}

func (a *App) InitUnauthedControllers() ([]controller.Controller, error) {
	log, err := a.container.GetLogger()
	if err != nil {
		return nil, err
	}

	appConfig, err := a.container.GetAppConfig()
	if err != nil {
		return nil, log.Propagate(err)
	}

	// Anime.

	animeListController, err := animeController.NewListController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	animeDetailController, err := animeController.NewDetailController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	// Auth.

	loginController, err := authController.NewLoginController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	registerController, err := authController.NewRegisterController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	refreshController, err := authController.NewRefreshController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	// Users.

	userListController, err := userController.NewListController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	userDetailController, err := userController.NewDetailController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	userCollectionListController, err := userController.NewCollectionListController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	// Collection.

	listCollectionController, err := collectionController.NewListController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	detailCollectionController, err := collectionController.NewDetailController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	animeListCollectionController, err := collectionController.NewAnimeListController(a.container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	controllers := []controller.Controller{
		// Anime.
		animeListController,
		animeDetailController,

		// Auth.
		loginController,
		registerController,
		refreshController,

		// Users.
		userListController,
		userDetailController,
		userCollectionListController,

		// Collection.
		listCollectionController,
		detailCollectionController,
		animeListCollectionController,
	}

	if appConfig.GetEnv().IsDevelopment() {
		animeCreateController, err := animeController.NewCreateController(a.container)
		if err != nil {
			return nil, log.Propagate(err)
		}

		controllers = append(controllers, animeCreateController)
	}

	return controllers, nil
}

/* ===== HttpServer ===== */

func (a *App) InitHttpServer(wg *sync.WaitGroup) error {
	loggerService, err := a.container.GetLogger()
	if err != nil {
		return err
	}

	ctx, err := a.container.GetAppContext()
	if err != nil {
		return loggerService.Propagate(err)
	}

	// RestAPI.
	authedAPIControllers, err := a.InitAuthedControllers()
	if err != nil {
		return loggerService.Propagate(err)
	}

	unauthedAPIController, err := a.InitUnauthedControllers()
	if err != nil {
		return loggerService.Propagate(err)
	}

	srv, err := server.NewServer(
		a.container,
		authedAPIControllers,
		unauthedAPIController,
	)
	if err != nil {
		return loggerService.Propagate(err)
	}

	wg.Add(1)
	go srv.Listen(ctx, wg)

	return nil
}

/* ===== AppContext ===== */

func (a *App) InitAppContext() context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())

	a.container.
		Set(ctx, reflectTypeOfNil[context.Context]()).
		Set(cancel, reflectTypeOfNil[context.CancelFunc]())

	return cancel
}

/* ===== AppConfig ===== */

func (a *App) InitAppConfig() error {
	a.container.Set(a.config, reflectTypeOfNil[config.AppConfig]())

	return nil
}

/* ===== Logger ===== */

func (a *App) InitLogger() (
	log loggerInterface.Logger,
	cancelFn func(),
	err error,
) {
	ctx, err := a.container.GetAppContext()
	if err != nil {
		return nil, nil, err
	}

	cfg, err := a.container.GetAppConfig()
	if err != nil {
		return nil, nil, err
	}

	if cfg.GetEnv().IsProduction() {
		log, cancelFn, err = logger.NewFile(ctx, cfg.Logger.ErrorBufferCap, cfg.Logger.RequestBufferCap)
		if err != nil {
			return nil, nil, err
		}
	} else {
		log, cancelFn = logger.NewStdErr(ctx, cfg.Logger.ErrorBufferCap, cfg.Logger.RequestBufferCap)
	}

	a.container.Set(log, reflectTypeOfNil[loggerInterface.Logger]())

	return log, cancelFn, nil
}

/* ===== Postgres ===== */

func (a *App) InitPostgres() (func(), error) {
	cfg, err := a.container.GetAppConfig()
	if err != nil {
		return nil, err
	}

	db, err := database.New(cfg.Database)
	if err != nil {
		return nil, err
	}

	a.container.Set(db, reflectTypeOfNil[postgres.DB]())
	a.container.Set(db, reflectTypeOfNil[postgres.RW]())

	return db.Close, nil
}

/* ===== ReqRes ===== */

func (a *App) InitReqRes() error {
	log, err := a.container.GetLogger()
	if err != nil {
		return err
	}

	// ParametersExtractor.
	parametersExtractor := extractor.NewParametersExtractor()

	a.container.Set(parametersExtractor, reflectTypeOfNil[extractorInterface.RequestParams]())

	// Responder.
	resp := responder.NewResponder(log)

	a.container.Set(resp, reflectTypeOfNil[responderInterface.Responder]())

	return nil
}

/* ===== Pagination ===== */

func (a *App) InitPagination() error {
	log, err := a.container.GetLogger()
	if err != nil {
		return err
	}

	// Builder.
	builder, err := builder.NewPagination(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(builder, reflectTypeOfNil[builderInterface.Pagination]())

	return nil
}

/* ===== Action ===== */

func (a *App) InitAction() error {
	log, err := a.container.GetLogger()
	if err != nil {
		return err
	}

	// Repository.
	repo, err := repositories.NewAction(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(repo, reflectTypeOfNil[repositoryInterface.Action]())

	// Service.
	service := action.NewAction(log, repo)

	a.container.Set(service, reflectTypeOfNil[actionInterface.Action]())

	// Builder.
	build, err := builder.NewAction(a.container)
	if err != nil {
		return err
	}

	a.container.Set(build, reflectTypeOfNil[builderInterface.Action]())

	// Validator.
	valid, err := validator.NewAction(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(valid, reflectTypeOfNil[validatorInterface.Action]())

	return nil
}

/* ===== Anime ===== */

func (a *App) InitAnime() error {
	log, err := a.container.GetLogger()
	if err != nil {
		return err
	}

	// Repository.
	repo, err := repositories.NewAnime(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(repo, reflectTypeOfNil[repositoryInterface.Anime]())

	// Service.
	service := anime.NewAnime(log, repo)

	a.container.Set(service, reflectTypeOfNil[animeInterface.Anime]())

	// Builder.
	build, err := builder.NewAnime(a.container)
	if err != nil {
		return err
	}

	a.container.Set(build, reflectTypeOfNil[builderInterface.Anime]())

	// Validator.
	valid, err := validator.NewAnime(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(valid, reflectTypeOfNil[validatorInterface.Anime]())

	return nil
}

/* ===== Collection ===== */

func (a *App) InitCollection() error {
	log, err := a.container.GetLogger()
	if err != nil {
		return err
	}

	animeCollectionRepo, err := repositories.NewAnimeCollection(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	// Repository.
	collectionRepo, err := repositories.NewCollection(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(collectionRepo, reflectTypeOfNil[repositoryInterface.Collection]())

	// Service.
	service := collection.NewCollection(log, animeCollectionRepo, collectionRepo)

	a.container.Set(service, reflectTypeOfNil[collectionInterface.Collection]())

	// Builder.
	build, err := builder.NewCollection(a.container)
	if err != nil {
		return err
	}

	a.container.Set(build, reflectTypeOfNil[builderInterface.Collection]())

	// Validator.
	valid, err := validator.NewCollection(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(valid, reflectTypeOfNil[validatorInterface.Collection]())

	return nil
}

/* ===== Token ===== */

func (a *App) InitToken() error {
	log, err := a.container.GetLogger()
	if err != nil {
		return err
	}

	// Repository.
	repo, err := repositories.NewToken(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(repo, reflectTypeOfNil[repositoryInterface.Token]())

	return nil
}

/* ===== CRUDService ===== */

func (a *App) InitUser() error {
	log, err := a.container.GetLogger()
	if err != nil {
		return err
	}

	// Repository.
	collectionRepo, err := repositories.NewUser(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(collectionRepo, reflectTypeOfNil[repositoryInterface.User]())

	// Validator.
	valid, err := validator.NewUser(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(valid, reflectTypeOfNil[validatorInterface.User]())

	// Service.
	service, err := user.NewCRUDService(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(service, reflectTypeOfNil[userInterface.CRUD]())

	// Builder.
	build, err := builder.NewUser(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(build, reflectTypeOfNil[builderInterface.User]())

	return nil
}

/* ===== Auth ===== */

func (a *App) InitAuth() error {
	log, err := a.container.GetLogger()
	if err != nil {
		return err
	}

	userRepo, err := a.container.GetUserRepository()
	if err != nil {
		return err
	}

	cfg, err := a.container.GetAppConfig()
	if err != nil {
		return err
	}

	token, err := a.container.GetTokenRepository()
	if err != nil {
		return err
	}

	// Tokenizer.
	jwtService := tokenizer.NewJwtService(log, cfg.Tokenizer, token)

	a.container.Set(jwtService, reflectTypeOfNil[tokenizerInterface.Tokenizer]())

	// Hasher.
	hasher := security.NewBcryptService(log, cfg.Hasher)

	a.container.Set(hasher, reflectTypeOfNil[securityInterface.Hasher]())

	// Service.
	service := auth.NewAuth(log, hasher, jwtService, userRepo)

	a.container.Set(service, reflectTypeOfNil[authInterface.Auth]())

	// Builder.
	build, err := builder.NewAuth(a.container)
	if err != nil {
		return err
	}

	a.container.Set(build, reflectTypeOfNil[builderInterface.Auth]())

	// Validator.
	valid, err := validator.NewAuth(a.container)
	if err != nil {
		return log.Propagate(err)
	}

	a.container.Set(valid, reflectTypeOfNil[validatorInterface.Auth]())

	return nil
}
