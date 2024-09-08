package diInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/config"
	"github.com/ilfey/hikilist-go/internal/domain/builder/interface"
	"github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/action/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/anime/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/auth/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/collection/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/extractor/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/responder/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/security/interface"
	tokenizerInterface "github.com/ilfey/hikilist-go/internal/domain/service/tokenizer/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/user/interface"
	"github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/infrastucture/di/container/interface"
	"github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/postgres"
)

type ServiceContainer interface {
	diInterface.Container

	GetAppConfig() (*config.AppConfig, error)
	GetAppContext() (context.Context, error)
	GetCancelFunc() (context.CancelFunc, error)

	GetLogger() (loggerInterface.Logger, error)

	GetRequestParametersExtractorService() (extractorInterface.RequestParams, error)

	GetResponderService() (responderInterface.Responder, error)
	GetTokenizerService() (tokenizerInterface.Tokenizer, error)
	GetHasherService() (securityInterface.Hasher, error)

	GetPostgresDatabase() (postgres.DB, error)

	GetPaginationBuilder() (builderInterface.Pagination, error)
	GetPaginationValidator() (validatorInterface.Pagination, error)

	GetActionRepository() (repositoryInterface.Action, error)
	GetActionService() (actionInterface.Action, error)
	GetActionBuilder() (builderInterface.Action, error)
	GetActionValidator() (validatorInterface.Action, error)

	GetAnimeRepository() (repositoryInterface.Anime, error)
	GetAnimeService() (animeInterface.Anime, error)
	GetAnimeBuilder() (builderInterface.Anime, error)
	GetAnimeValidator() (validatorInterface.Anime, error)
	GetAnimeCollectionRepository() (repositoryInterface.AnimeCollection, error)

	GetCollectionRepository() (repositoryInterface.Collection, error)
	GetCollectionService() (collectionInterface.Collection, error)
	GetCollectionBuilder() (builderInterface.Collection, error)
	GetCollectionValidator() (validatorInterface.Collection, error)

	GetTokenRepository() (repositoryInterface.Token, error)

	GetUserRepository() (repositoryInterface.User, error)
	GetUserService() (userInterface.CRUD, error)
	GetUserBuilder() (builderInterface.User, error)
	GetUserValidator() (validatorInterface.User, error)

	GetAuthService() (authInterface.Auth, error)
	GetAuthBuilder() (builderInterface.Auth, error)
	GetAuthValidator() (validatorInterface.Auth, error)
}
