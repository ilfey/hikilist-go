package router

import (
	"net/http"

	"github.com/gorilla/mux"

	animeController "github.com/ilfey/hikilist-go/api/controllers/anime"
	authController "github.com/ilfey/hikilist-go/api/controllers/auth"
	collectionController "github.com/ilfey/hikilist-go/api/controllers/collection"
	userController "github.com/ilfey/hikilist-go/api/controllers/user"

	animeService "github.com/ilfey/hikilist-go/services/anime"
	authService "github.com/ilfey/hikilist-go/services/auth"
	collectionService "github.com/ilfey/hikilist-go/services/collection"
	userService "github.com/ilfey/hikilist-go/services/user"
	userActionService "github.com/ilfey/hikilist-go/services/user_action"
)

// Роутер
type Router struct {
	AnimeService      animeService.Service
	AuthService       authService.Service
	CollectionService collectionService.Service
	UserService       userService.Service
	UserActionService userActionService.Service
}

// Привязка роутера
func (r *Router) Bind() http.Handler {
	router := mux.NewRouter()

	// router.NotFoundHandler = http.HandlerFunc(r.NotFoundHandler)
	// router.MethodNotAllowedHandler = http.HandlerFunc(r.NotFoundHandler)

	// router.Use(AuthorizationMiddleware(r.jwt))

	router = animeController.New(
		r.AuthService,
		r.AnimeService,
	).Bind(router)

	router = authController.New(
		r.AuthService,
		r.UserService,
	).Bind(router)

	router = userController.New(
		r.AuthService,
		r.CollectionService,
		r.UserService,
		r.UserActionService,
	).Bind(router)

	router = collectionController.New(
		r.AuthService,
		// r.AnimeService,
		r.CollectionService,
	).Bind(router)

	return router
}

// func (*Router) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
// 	resx.ResponseNotFound.JSON(w)
// }

// func AuthorizationMiddleware(j *jwt.JWT) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			token := r.Header.Get("Authorization")
// 			if token == "" {
// 				next.ServeHTTP(w, r)
// 				return
// 			}

// 			token = strings.TrimPrefix(token, "Bearer ")
// 			claims, ok := j.ParseToken(token)
// 			if !ok {
// 				next.ServeHTTP(w, r)
// 				return
// 			}

// 			r.Header.Set("user_id", strconv.FormatUint(uint64(claims.UserID), 10))

// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }
