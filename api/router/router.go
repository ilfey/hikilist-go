package router

import (
	"github.com/gorilla/mux"
	"github.com/ilfey/hikilist-go/api/controllers/book"
)

type Router struct {
	BookRoutes *book.ControllerRoute
}

// Router constructor
func NewRouter(bookRoutes *book.ControllerRoute) *Router {
	return &Router{
		BookRoutes: bookRoutes,
	}
}

// Create router and register routes
func (routes *Router) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = routes.BookRoutes.Route(router)
	return router
}