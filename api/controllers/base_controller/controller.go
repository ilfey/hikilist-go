package baseController

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilfey/hikilist-go/api/controllers/base_controller/handler"
	authService "github.com/ilfey/hikilist-go/services/auth"
)

type Controller struct {
	AuthService authService.Service

	router *mux.Router
}

func (c *Controller) HandleFunc(url string, fn handler.HandleFunc) *mux.Route {
	return c.router.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		ctx := handler.NewContext(c.AuthService, w, r)

		fn(ctx)
	})
}

func (c *Controller) Bind(router *mux.Router) {
	c.router = router
}
