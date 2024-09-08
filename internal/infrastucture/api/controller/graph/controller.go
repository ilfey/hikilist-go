package graph

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	errtypeInterface "github.com/ilfey/hikilist-go/internal/domain/errtype/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	"github.com/ilfey/hikilist-go/internal/infrastucture/api/controller/graph/exec"
	"github.com/ilfey/hikilist-go/internal/infrastucture/api/controller/graph/resolvers"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	Path           = "/graphql"
	PlaygroundPath = "/play"
)

type GQLController struct {
	log          loggerInterface.Logger
	resolver     *resolvers.Resolver
	graphHandler *handler.Server
}

func NewGQLController(container diInterface.ServiceContainer) (*GQLController, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	resolver, err := resolvers.NewResolver(container)
	if err != nil {
		return nil, log.Propagate(err)
	}

	graphHandler := handler.NewDefaultServer(
		exec.NewExecutableSchema(
			exec.Config{
				Schema:     nil,
				Resolvers:  resolver,
				Directives: exec.DirectiveRoot{},
				Complexity: exec.ComplexityRoot{},
			},
		),
	)

	return &GQLController{
		log:          log,
		resolver:     resolver,
		graphHandler: graphHandler,
	}, nil
}

func (c *GQLController) Init() {
	c.graphHandler.SetErrorPresenter(c.errorPresenter)
}

func (c *GQLController) errorPresenter(ctx context.Context, e error) *gqlerror.Error {
	err := graphql.DefaultErrorPresenter(ctx, e)

	c.log.Error(err)

	var publicErr errtypeInterface.PublicError
	if errors.As(e, &publicErr) {

		// Handle validation error.
		var validationErr *errtype.ValidatorError
		if errors.As(e, &validationErr) {
			err.Message = validationErr.ErrorDetail
			err.Extensions = map[string]any{
				"fields": validationErr.Expectations,
			}

			return err
		}

		err.Message = e.Error()

		return err
	}

	err.Message = "internal server error"

	return err
}

func (c *GQLController) AddRoute(router *mux.Router) {
	c.Init()

	r := router.
		Path(Path).
		Handler(c.graphHandler)

	path, err := r.URLPath()
	if err != nil {
		c.log.Critical(err)
		panic(err)
	}

	// Add playground
	router.Path(PlaygroundPath).
		Handler(playground.Handler("Hikilist", path.Path))
}
