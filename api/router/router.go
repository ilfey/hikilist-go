package router

import (
	"net/http"

	"github.com/gorilla/mux"

	animeController "github.com/ilfey/hikilist-go/api/controllers/anime"
	authController "github.com/ilfey/hikilist-go/api/controllers/auth"
	collectionController "github.com/ilfey/hikilist-go/api/controllers/collection"
	userController "github.com/ilfey/hikilist-go/api/controllers/user"

	authService "github.com/ilfey/hikilist-go/services/auth"
)

// Роутер
type Router struct {
	AuthService authService.Service
}

// Привязка роутера
func (r *Router) Bind() http.Handler {
	router := mux.NewRouter()

	// router.NotFoundHandler = http.HandlerFunc(r.NotFoundHandler)
	// router.MethodNotAllowedHandler = http.HandlerFunc(r.NotFoundHandler)

	// router.Use(AuthorizationMiddleware(r.jwt))

	router = animeController.New(
		r.AuthService,
	).Bind(router)

	router = authController.New(
		r.AuthService,
	).Bind(router)

	router = userController.New(
		r.AuthService,
	).Bind(router)

	router = collectionController.New(
		r.AuthService,
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
