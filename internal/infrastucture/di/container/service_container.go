package container

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
	extractorInterface "github.com/ilfey/hikilist-go/internal/domain/service/extractor/interface"
	responderInterface "github.com/ilfey/hikilist-go/internal/domain/service/responder/interface"
	securityInterface "github.com/ilfey/hikilist-go/internal/domain/service/security/interface"
	tokenizerInterface "github.com/ilfey/hikilist-go/internal/domain/service/tokenizer/interface"
	userInterface "github.com/ilfey/hikilist-go/internal/domain/service/user/interface"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/infrastucture/di/container/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/postgres"
	"reflect"
)

type ServiceContainer struct {
	diInterface.Container
}

func NewServiceContainerManager() *ServiceContainer {
	return &ServiceContainer{
		Container: NewServiceContainer(),
	}
}

/* ===== Config ===== */

func (s *ServiceContainer) GetAppConfig() (*config.AppConfig, error) {
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

func (s *ServiceContainer) GetAppContext() (context.Context, error) {
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

func (s *ServiceContainer) GetCancelFunc() (context.CancelFunc, error) {
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

func (s *ServiceContainer) GetLogger() (loggerInterface.Logger, error) {
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

func (s *ServiceContainer) GetRequestParametersExtractorService() (extractorInterface.RequestParams, error) {
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

func (s *ServiceContainer) GetResponderService() (responderInterface.Responder, error) {
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

func (s *ServiceContainer) GetHasherService() (securityInterface.Hasher, error) {
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

func (s *ServiceContainer) GetPostgresDatabase() (postgres.DB, error) {
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

func (s *ServiceContainer) GetPaginationBuilder() (builderInterface.Pagination, error) {
	key := (*builderInterface.Pagination)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(builderInterface.Pagination)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

func (s *ServiceContainer) GetPaginationValidator() (validatorInterface.Pagination, error) {
	key := (*validatorInterface.Pagination)(nil)
	reflectService, err := s.Get(reflect.TypeOf(key))
	if err != nil {
		return nil, errtype.NewServiceWasNotFoundIntoContainerError(reflect.TypeOf(key))
	}
	service, ok := reflectService.Interface().(validatorInterface.Pagination)
	if !ok {
		return nil, errtype.NewTypesMismatchedServiceContainerError(reflect.TypeOf(reflectService), reflect.TypeOf(key))
	}
	return service, nil
}

/* ===== Action ===== */

func (s *ServiceContainer) GetActionRepository() (repositoryInterface.Action, error) {
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

func (s *ServiceContainer) GetActionService() (actionInterface.Action, error) {
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

func (s *ServiceContainer) GetActionBuilder() (builderInterface.Action, error) {
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

func (s *ServiceContainer) GetActionValidator() (validatorInterface.Action, error) {
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

func (s *ServiceContainer) GetAnimeRepository() (repositoryInterface.Anime, error) {
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

func (s *ServiceContainer) GetAnimeService() (animeInterface.Anime, error) {
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

func (s *ServiceContainer) GetAnimeBuilder() (builderInterface.Anime, error) {
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

func (s *ServiceContainer) GetAnimeValidator() (validatorInterface.Anime, error) {
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

func (s *ServiceContainer) GetAnimeCollectionRepository() (repositoryInterface.AnimeCollection, error) {
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

func (s *ServiceContainer) GetCollectionRepository() (repositoryInterface.Collection, error) {
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

func (s *ServiceContainer) GetCollectionService() (collectionInterface.Collection, error) {
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

func (s *ServiceContainer) GetCollectionBuilder() (builderInterface.Collection, error) {
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

func (s *ServiceContainer) GetCollectionValidator() (validatorInterface.Collection, error) {
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

func (s *ServiceContainer) GetTokenRepository() (repositoryInterface.Token, error) {
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

func (s *ServiceContainer) GetUserRepository() (repositoryInterface.User, error) {
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

func (s *ServiceContainer) GetUserService() (userInterface.CRUD, error) {
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

func (s *ServiceContainer) GetUserBuilder() (builderInterface.User, error) {
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

func (s *ServiceContainer) GetUserValidator() (validatorInterface.User, error) {
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

func (s *ServiceContainer) GetTokenizerService() (tokenizerInterface.Tokenizer, error) {
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

func (s *ServiceContainer) GetAuthService() (authInterface.Auth, error) {
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

func (s *ServiceContainer) GetAuthBuilder() (builderInterface.Auth, error) {
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

func (s *ServiceContainer) GetAuthValidator() (validatorInterface.Auth, error) {
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
