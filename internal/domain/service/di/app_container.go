package di

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/config"
	builderInterface "github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	actionInterface "github.com/ilfey/hikilist-go/internal/domain/service/action/interface"
	animeInterface "github.com/ilfey/hikilist-go/internal/domain/service/anime/interface"
	authInterface "github.com/ilfey/hikilist-go/internal/domain/service/auth/interface"
	collectionInterface "github.com/ilfey/hikilist-go/internal/domain/service/collection/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/di/container"
	"github.com/ilfey/hikilist-go/internal/domain/service/di/container/interface"
	extractorInterface "github.com/ilfey/hikilist-go/internal/domain/service/extractor/interface"
	responderInterface "github.com/ilfey/hikilist-go/internal/domain/service/responder/interface"
	securityInterface "github.com/ilfey/hikilist-go/internal/domain/service/security/interface"
	tokenizerInterface "github.com/ilfey/hikilist-go/internal/domain/service/tokenizer/interface"
	userInterface "github.com/ilfey/hikilist-go/internal/domain/service/user/interface"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/postgres"
	"reflect"
)

type AppContainer struct {
	containerInterface.Container
}

func NewServiceContainerManager() *AppContainer {
	return &AppContainer{
		Container: container.NewServiceContainer(),
	}
}

/* ===== Config ===== */

func (s *AppContainer) GetAppConfig() (*config.AppConfig, error) {
	key := (*config.AppConfig)(nil)
	service, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	cfg, ok := service.Interface().(*config.AppConfig)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(service), reflect.TypeOf(key))
	}
	return cfg, nil
}

/* ===== AppContext ===== */

func (s *AppContainer) GetAppContext() (context.Context, error) {
	key := (*context.Context)(nil)
	service, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	ctx, ok := service.Interface().(context.Context)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(service), reflect.TypeOf(key))
	}
	return ctx, nil
}

/* ===== CancelFunc ===== */

func (s *AppContainer) GetCancelFunc() (context.CancelFunc, error) {
	key := (*context.CancelFunc)(nil)
	service, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	cancel, ok := service.Interface().(context.CancelFunc)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(service), reflect.TypeOf(key))
	}
	return cancel, nil
}

/* ===== Logger ===== */

func (s *AppContainer) GetLogger() (loggerInterface.Logger, error) {
	key := (*loggerInterface.Logger)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(loggerInterface.Logger)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

/* ===== Extractor ===== */

func (s *AppContainer) GetRequestParametersExtractorService() (extractorInterface.RequestParams, error) {
	key := (*extractorInterface.RequestParams)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(extractorInterface.RequestParams)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

/* ===== Responder ===== */

func (s *AppContainer) GetResponderService() (responderInterface.Responder, error) {
	key := (*responderInterface.Responder)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(responderInterface.Responder)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

/* ===== Hasher ===== */

func (s *AppContainer) GetHasherService() (securityInterface.Hasher, error) {
	key := (*securityInterface.Hasher)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(securityInterface.Hasher)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

/* ===== Postgres ===== */

func (s *AppContainer) GetPostgresDatabase() (postgres.DB, error) {
	key := (*postgres.DB)(nil)
	service, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	database, ok := service.Interface().(postgres.DB)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(service), reflect.TypeOf(key))
	}
	return database, nil
}

/* ===== Pagination ===== */

func (s *AppContainer) GetPaginationBuilder() (builderInterface.Pagination, error) {
	key := (*builderInterface.Pagination)(nil)
	service, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	database, ok := service.Interface().(builderInterface.Pagination)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(service), reflect.TypeOf(key))
	}
	return database, nil
}

/* ===== Action ===== */

func (s *AppContainer) GetActionRepository() (repositoryInterface.Action, error) {
	key := (*repositoryInterface.Action)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(repositoryInterface.Action)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetActionService() (actionInterface.Action, error) {
	key := (*actionInterface.Action)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(actionInterface.Action)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetActionBuilder() (builderInterface.Action, error) {
	key := (*builderInterface.Action)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(builderInterface.Action)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetActionValidator() (validatorInterface.Action, error) {
	key := (*validatorInterface.Action)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(validatorInterface.Action)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

/* ===== Anime ===== */

func (s *AppContainer) GetAnimeRepository() (repositoryInterface.Anime, error) {
	key := (*repositoryInterface.Anime)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(repositoryInterface.Anime)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetAnimeService() (animeInterface.Anime, error) {
	key := (*animeInterface.Anime)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(animeInterface.Anime)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetAnimeBuilder() (builderInterface.Anime, error) {
	key := (*builderInterface.Anime)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(builderInterface.Anime)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetAnimeValidator() (validatorInterface.Anime, error) {
	key := (*validatorInterface.Anime)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(validatorInterface.Anime)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

/* ===== AnimeCollection ===== */

func (s *AppContainer) GetAnimeCollectionRepository() (repositoryInterface.AnimeCollection, error) {
	key := (*repositoryInterface.AnimeCollection)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(repositoryInterface.AnimeCollection)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

/* ===== Collection ===== */

func (s *AppContainer) GetCollectionRepository() (repositoryInterface.Collection, error) {
	key := (*repositoryInterface.Collection)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(repositoryInterface.Collection)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetCollectionService() (collectionInterface.Collection, error) {
	key := (*collectionInterface.Collection)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(collectionInterface.Collection)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetCollectionBuilder() (builderInterface.Collection, error) {
	key := (*builderInterface.Collection)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(builderInterface.Collection)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetCollectionValidator() (validatorInterface.Collection, error) {
	key := (*validatorInterface.Collection)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(validatorInterface.Collection)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

/* ===== Token ===== */

func (s *AppContainer) GetTokenRepository() (repositoryInterface.Token, error) {
	key := (*repositoryInterface.Token)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(repositoryInterface.Token)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

/* ===== CRUD ===== */

func (s *AppContainer) GetUserRepository() (repositoryInterface.User, error) {
	key := (*repositoryInterface.User)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(repositoryInterface.User)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetUserService() (userInterface.CRUD, error) {
	key := (*userInterface.CRUD)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(userInterface.CRUD)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetUserBuilder() (builderInterface.User, error) {
	key := (*builderInterface.User)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(builderInterface.User)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetUserValidator() (validatorInterface.User, error) {
	key := (*validatorInterface.User)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(validatorInterface.User)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

/* ===== Tokenizer ===== */

func (s *AppContainer) GetTokenizerService() (tokenizerInterface.Tokenizer, error) {
	key := (*tokenizerInterface.Tokenizer)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(tokenizerInterface.Tokenizer)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

/* ===== Auth ===== */

func (s *AppContainer) GetAuthService() (authInterface.Auth, error) {
	key := (*authInterface.Auth)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(authInterface.Auth)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetAuthBuilder() (builderInterface.Auth, error) {
	key := (*builderInterface.Auth)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(builderInterface.Auth)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *AppContainer) GetAuthValidator() (validatorInterface.Auth, error) {
	key := (*validatorInterface.Auth)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(validatorInterface.Auth)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}
