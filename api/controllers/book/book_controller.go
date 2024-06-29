package book

import (
	"encoding/json"
	"net/http"
	bookModels "github.com/ilfey/hikilist-go/internal/models/book"
	"github.com/ilfey/hikilist-go/internal/services/book"
)

type Controller struct {
	service *book.Service
}

// Controller constructor
func NewController(service *book.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create book handler
func (controller *Controller) CreateBook(w http.ResponseWriter, r *http.Request) {
	request := new(bookModels.CreateModel)
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&request)

	controller.service.Create(request)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}