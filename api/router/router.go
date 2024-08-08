package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ilfey/hikilist-go/api/controllers/anime"
	"github.com/ilfey/hikilist-go/api/controllers/auth"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/responses"
	"github.com/ilfey/hikilist-go/api/controllers/collection"
	"github.com/ilfey/hikilist-go/api/controllers/user"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"

	authService "github.com/ilfey/hikilist-go/services/auth"
)

type Router struct {
	AuthService authService.Service

	router *mux.Router
}

func New(
	authService authService.Service,
) *Router {
	return &Router{
		AuthService: authService,

		router: mux.NewRouter(),
	}
}

// Bind controllers
func (r *Router) Bind() http.Handler {
	r.router.NotFoundHandler = http.HandlerFunc(r.NotFoundHandler)
	r.router.MethodNotAllowedHandler = http.HandlerFunc(r.MethodNotAllowedHandler)

	anime := anime.AnimeController{}

	// c.HandleFunc("/api/animes", c.Create).Methods("POST")
	r.HandleFunc("/api/animes", anime.List).Methods("GET")

	r.HandleFunc("/api/animes/{id:[0-9]+}", anime.Detail).Methods("GET")

	auth := auth.AuthController{}

	r.HandleFunc("/api/auth/login", auth.Login).Methods("POST")
	r.HandleFunc("/api/auth/register", auth.Register).Methods("POST")
	r.HandleFunc("/api/auth/refresh", auth.Refresh).Methods("POST")
	r.HandleFunc("/api/auth/logout", auth.Logout).Methods("POST")

	collection := collection.CollectionController{}

	r.HandleFunc("/api/collections", collection.Create).Methods("POST")
	r.HandleFunc("/api/collections", collection.List).Methods("GET")

	r.HandleFunc("/api/collections/{id:[0-9]+}", collection.Detail).Methods("GET")
	r.HandleFunc("/api/collections/{id:[0-9]+}/animes", collection.Animes).Methods("GET")

	user := user.UserController{}

	r.HandleFunc("/api/users", user.List).Methods("GET")

	r.HandleFunc("/api/users/{id:[0-9]+}", user.Detail).Methods("GET")
	r.HandleFunc("/api/users/{id:[0-9]+}/collections", user.Collections).Methods("GET")

	r.HandleFunc("/api/users/@{username:[a-zA-Z0-9]+}", user.DetailByUsername).Methods("GET")

	r.HandleFunc("/api/users/me", user.Me).Methods("GET")
	r.HandleFunc("/api/users/me/delete", user.Delete).Methods("DELETE")
	r.HandleFunc("/api/users/me/actions", user.MyActions).Methods("GET")
	r.HandleFunc("/api/users/me/collections", user.MyCollections).Methods("GET")

	return r.router
}

func (r *Router) HandleFunc(path string, fn func(*handler.Context)) *mux.Route {
	return r.router.HandleFunc(path, r.provideContext(fn))
}

func (r *Router) provideContext(
	fn func(*handler.Context),
) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := handler.NewContext(r.AuthService, rw, req)

		var doneCh = make(chan struct{})

		// Update user last online
		go func() {
			defer func() {
				doneCh <- struct{}{}
			}()

			user, err := ctx.GetUser()
			if err != nil {
				return
			}

			err = ctx.AuthService.UpdateUserOnline(ctx, user)
			if err != nil {
				logger.Errorf("Failed to update user online: %v", err)
			}
		}()

		fn(ctx)

		<-doneCh

		close(doneCh)
	}
}

func (*Router) NotFoundHandler(rw http.ResponseWriter, r *http.Request) {
	data, code := responses.ResponseNotFound()

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	rw.Write(
		errorsx.Must(
			json.Marshal(data),
		),
	)
}

func (*Router) MethodNotAllowedHandler(rw http.ResponseWriter, r *http.Request) {
	data, code := responses.ResponseMethodNotAllowed()

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	rw.Write(
		errorsx.Must(
			json.Marshal(data),
		),
	)
}
