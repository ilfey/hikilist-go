package book

import (
	"github.com/gorilla/mux"
)

type ControllerRoute struct {
	Controller *Controller
}

// ControllerRoute constructor
func NewControllerRoute(controller *Controller) *ControllerRoute {
	return &ControllerRoute{Controller: controller}
}

// Register controller routes
func (route *ControllerRoute) Route(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/books", route.Controller.CreateBook).Methods("POST")
	return router
}