package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilfey/hikilist-go/api/controllers"
	animeService "github.com/ilfey/hikilist-go/services/anime"
	authService "github.com/ilfey/hikilist-go/services/auth"
	userService "github.com/ilfey/hikilist-go/services/user"
	"github.com/ilfey/hikilist-go/internal/utils/resx"
)

// Роутер
type Router struct {
	AnimeService *animeService.Service
	AuthService  *authService.Service
	UserService  *userService.Service
}

// Привязка роутера
func (r *Router) Bind() http.Handler {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(r.NotFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(r.NotFoundHandler)

	// router.Use(AuthorizationMiddleware(r.jwt))

	router = controllers.NewAnimeController(r.AnimeService).Bind(router)
	router = controllers.NewAuthController(r.AuthService, r.UserService).Bind(router)
	router = controllers.NewUserController(r.AuthService, r.UserService).Bind(router)

	return router
}

func (*Router) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	resx.ResponseNotFound.JSON(w)
}

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
