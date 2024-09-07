package server

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/ilfey/hikilist-go/internal/config/server"
	"github.com/ilfey/hikilist-go/internal/domain/enum"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	authInterface "github.com/ilfey/hikilist-go/internal/domain/service/auth/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/extractor"
	extractorInterface "github.com/ilfey/hikilist-go/internal/domain/service/extractor/interface"
	reponderInterface "github.com/ilfey/hikilist-go/internal/domain/service/responder/interface"
	userInterface "github.com/ilfey/hikilist-go/internal/domain/service/user/interface"
	"github.com/ilfey/hikilist-go/internal/infrastucture/api/controller"
	"github.com/ilfey/hikilist-go/internal/infrastucture/helper/ruid"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/pkg/errors"
	"net"
	"sync"
	"time"

	"net/http"
)

type Server struct {
	ctx context.Context

	logger             loggerInterface.Logger
	auth               authInterface.Auth
	responder          reponderInterface.Responder
	reqParamsExtractor extractorInterface.RequestParams
	config             *server.Config

	user userInterface.CRUD

	authedControllers   []controller.Controller
	unauthedControllers []controller.Controller
}

type Router interface {
	Bind() http.Handler
}

func NewServer(
	container diInterface.ServiceContainer,

	authedControllers []controller.Controller,
	unauthedControllers []controller.Controller,
) (*Server, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	appCtx, err := container.GetAppContext()
	if err != nil {
		return nil, log.Propagate(err)
	}

	auth, err := container.GetAuthService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	responder, err := container.GetResponderService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	reqParamsExtractor, err := container.GetRequestParametersExtractorService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	appConfig, err := container.GetAppConfig()
	if err != nil {
		return nil, log.Propagate(err)
	}

	user, err := container.GetUserService()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &Server{
		ctx: appCtx,

		logger:             log,
		auth:               auth,
		responder:          responder,
		reqParamsExtractor: reqParamsExtractor,
		config:             appConfig.Server,

		user: user,

		authedControllers:   authedControllers,
		unauthedControllers: unauthedControllers,
	}, nil
}

func (server *Server) Listen(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	addr, err := net.ResolveTCPAddr("tcp", server.config.Address())
	if err != nil {
		server.logger.Critical(err)
		return
	}

	httpServer := &http.Server{
		ReadTimeout:       server.config.ReadTimeout * time.Millisecond,
		WriteTimeout:      server.config.WriteTimeout * time.Millisecond,
		IdleTimeout:       server.config.IdleTimeout * time.Millisecond,
		ReadHeaderTimeout: server.config.ReadHeaderTimeout * time.Millisecond,

		Addr:    addr.String(),
		Handler: server.addRoutes(),
	}

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			server.logger.Info("stopped")
		}()
		if lsErr := httpServer.ListenAndServe(); lsErr != nil && !errors.Is(lsErr, http.ErrServerClosed) {
			server.logger.Critical(lsErr)
			return
		}
	}()

	server.logger.Info("running...")

	<-ctx.Done()

	server.logger.Info("shutting down...")

	time.Sleep(time.Second)

	serverCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if shErr := httpServer.Shutdown(serverCtx); shErr != nil && !errors.Is(shErr, context.Canceled) {
		server.logger.Critical(err)
		return
	}
}

func (server *Server) addRoutes() *mux.Router {
	router := mux.NewRouter()

	server.addAuthedRoutes(router)
	server.addUnauthedRoutes(router)

	return router
}

func (server *Server) addAuthedRoutes(router *mux.Router) {
	restAuthedRouterV1 := router.
		PathPrefix("/api/v1").
		Subrouter()
	restAuthedRouterV1.
		Use(
			server.apiHeaderMiddleware,
			server.loggingMiddleware,
			server.authorizationMiddleware,
			server.onlineMiddleware,
			server.authorizationRequiredMiddleware,
		)

	for _, c := range server.authedControllers {
		c.AddRoute(restAuthedRouterV1)
	}
}

func (server *Server) addUnauthedRoutes(router *mux.Router) {
	restUnauthedRouterV1 := router.
		PathPrefix("/api/v1").
		Subrouter()
	restUnauthedRouterV1.
		Use(
			server.apiHeaderMiddleware,
			server.loggingMiddleware,
			server.authorizationMiddleware,
			server.onlineMiddleware,
		)

	for _, c := range server.unauthedControllers {
		c.AddRoute(restUnauthedRouterV1)
	}
}

func (server *Server) authorizationMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			userID, err := server.auth.IsAuthed(r)
			if err == nil && userID > 0 {
				// Create a new context with userID value.
				ctx := context.WithValue(r.Context(), enum.UserIDContextKey, userID)

				// Serve the next layer.
				handler.ServeHTTP(w, r.WithContext(ctx))

				return
			}

			// serve the next layer
			handler.ServeHTTP(w, r)
		},
	)
}

func (server *Server) authorizationRequiredMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if _, ok := r.Context().Value(enum.UserIDContextKey).(uint64); !ok {
				server.responder.Respond(w, server.logger.Propagate(errtype.NewAuthFailedError("")))
				return
			}

			// Serve the next layer.
			handler.ServeHTTP(w, r)
		},
	)
}

func (server *Server) apiHeaderMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Adding the rest api header.
			w.Header().Set("Content-Type", "application/json")

			// Serve the next layer.
			handler.ServeHTTP(w, r)
		},
	)
}

func (server *Server) onlineMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var (
				wg *sync.WaitGroup
			)

			userID, ok := r.Context().Value(enum.UserIDContextKey).(uint64)
			if ok {
				wg = &sync.WaitGroup{}

				wg.Add(1)
				go func(ctx context.Context) {
					defer wg.Done()
					err := server.user.UpdateLastOnline(ctx, userID)
					if err != nil {
						server.logger.Error(err)
					}
				}(r.Context())
			}

			// Serve the next layer.
			handler.ServeHTTP(w, r)

			if wg != nil {
				wg.Wait()
			}
		},
	)
}

func (server *Server) loggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			uniqueReqID := ruid.RequestUniqueID(r)
			requestData := &extractor.LoggableData{
				Date:       time.Now(),
				ReqID:      uniqueReqID,
				Type:       extractor.LogType,
				Method:     r.Method,
				URL:        r.URL.String(),
				Header:     r.Header,
				RemoteAddr: r.RemoteAddr,
				Params:     server.reqParamsExtractor.Parameters(r),
			}

			// Log the request.
			server.logger.Object(requestData)

			// Pass a requestID through entire app.
			server.logger.SetContext(context.WithValue(server.ctx, enum.RequestIDContextKey, uniqueReqID)) //nolint:contextcheck

			// Serve the next layer.
			handler.ServeHTTP(w, r)
		},
	)
}
