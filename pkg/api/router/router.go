package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/pkg/api/controllers"
	"github.com/ilfey/hikilist-go/pkg/api/handler"
	"github.com/ilfey/hikilist-go/pkg/api/responses"

	"github.com/ilfey/hikilist-go/pkg/services"
)

type Router struct {
	logger *logrus.Logger

	action     services.Action
	anime      services.Anime
	auth       services.Auth
	collection services.Collection
	user       services.User

	router *mux.Router
}

func New(
	logger *logrus.Logger,

	action services.Action,
	anime services.Anime,
	auth services.Auth,
	collection services.Collection,
	user services.User,
) *Router {
	return &Router{
		logger: logger,

		action:     action,
		anime:      anime,
		auth:       auth,
		collection: collection,
		user:       user,

		router: mux.NewRouter(),
	}
}

// Bind controllers
func (r *Router) Bind() http.Handler {
	r.router.NotFoundHandler = http.HandlerFunc(r.NotFoundHandler)
	r.router.MethodNotAllowedHandler = http.HandlerFunc(r.MethodNotAllowedHandler)

	anime := controllers.Anime{
		Logger: r.createControllerLogger("Anime"),

		Anime: r.anime,
	}

	// c.HandleFunc("/api/animes", c.Create).Methods("POST")
	r.HandleFunc("/api/animes", anime.List).Methods("GET")

	r.HandleFunc("/api/animes/{id:[0-9]+}", anime.Detail).Methods("GET")

	auth := controllers.Auth{
		Logger: r.createControllerLogger("Auth"),

		Auth: r.auth,
		User: r.user,
	}

	r.HandleFunc("/api/auth/login", auth.Login).Methods("POST")
	r.HandleFunc("/api/auth/register", auth.Register).Methods("POST")
	r.HandleFunc("/api/auth/refresh", auth.Refresh).Methods("POST")
	r.HandleFunc("/api/auth/logout", auth.Logout).Methods("POST")
	r.HandleFunc("/api/auth/delete", auth.Delete).Methods("DELETE")

	collection := controllers.Collection{
		Logger: r.createControllerLogger("Collection"),

		Anime:      r.anime,
		Collection: r.collection,
	}

	r.HandleFunc("/api/collections", collection.Create).Methods("POST")
	r.HandleFunc("/api/collections", collection.List).Methods("GET")

	r.HandleFunc("/api/collections/{id:[0-9]+}", collection.Detail).Methods("GET")
	r.HandleFunc("/api/collections/{id:[0-9]+}", collection.Update).Methods("PATCH")
	r.HandleFunc("/api/collections/{id:[0-9]+}/animes", collection.Animes).Methods("GET")
	r.HandleFunc("/api/collections/{id:[0-9]+}/animes/add", collection.AddAnimes).Methods("PATCH")
	r.HandleFunc("/api/collections/{id:[0-9]+}/animes/remove", collection.RemoveAnimes).Methods("PATCH")

	user := controllers.User{
		Logger: r.createControllerLogger("User"),

		Action:     r.action,
		Auth:       r.auth,
		Collection: r.collection,
		User:       r.user,
	}

	r.HandleFunc("/api/users", user.List).Methods("GET")

	r.HandleFunc("/api/users/{id:[0-9]+}", user.Detail).Methods("GET")
	r.HandleFunc("/api/users/{id:[0-9]+}/collections", user.Collections).Methods("GET")

	r.HandleFunc("/api/users/me", user.Me).Methods("GET")
	r.HandleFunc("/api/users/me/password", user.ChangePassword).Methods("PATCH")
	r.HandleFunc("/api/users/me/username", user.ChangeUsername).Methods("PATCH")
	r.HandleFunc("/api/users/me/actions", user.MyActions).Methods("GET")
	r.HandleFunc("/api/users/me/collections", user.MyCollections).Methods("GET")

	return r.router
}

func (r *Router) createControllerLogger(controllerName string) logrus.FieldLogger {
	return r.logger.WithField("controller", controllerName)
}

func (r *Router) HandleFunc(path string, fn func(*handler.Context)) *mux.Route {
	return r.router.HandleFunc(path, r.provideContext(fn))
}

func (r *Router) createContext(rw http.ResponseWriter, req *http.Request) *handler.Context {
	return handler.NewContext(
		r.auth,
		r.user,
		rw,
		req,
	)
}

func (r *Router) provideContext(
	fn func(*handler.Context),
) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := r.createContext(rw, req)

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

			err = r.user.UpdateLastOnline(ctx, user.ID)
			if err != nil {
				r.logger.Errorf("Failed to update user online %v", err)
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
