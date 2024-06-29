package server

import (
	"<package_name>/api/router"
	"<package_name>/internal/config"

	"github.com/codegangsta/negroni"

	"net/http"
)

type Server struct {
	AppConfig *config.Config
	Router    *router.Router
}

// Server constructor
func NewServer(appConfig *config.Config, router *router.Router) *Server {
	return &Server{
		AppConfig: appConfig,
		Router:    router,
	}
}

// Run server
func (server *Server) Run() {
	ngRouter := server.Router.InitRoutes()

	ngClassic := negroni.Classic()

	ngClassic.UseHandler(ngRouter)

	err := http.ListenAndServe(":5000", ngClassic)
	if err != nil {
		return
	}
}
