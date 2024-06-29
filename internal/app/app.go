package app

import (
	"go.uber.org/dig"
	"github.com/ilfey/hikilist-go/api/controllers/book"
	"github.com/ilfey/hikilist-go/api/router"
	"github.com/ilfey/hikilist-go/internal/config"
	bookRepository "github.com/ilfey/hikilist-go/internal/repositories/book"
	bookService "github.com/ilfey/hikilist-go/internal/services/book"
	"github.com/ilfey/hikilist-go/server"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.NewConfig)
	_ = container.Provide(server.NewServer)
	_ = container.Provide(router.NewRouter)

	buildBook(container)

	return container
}

func buildBook(container *dig.Container) {
	_ = container.Provide(book.NewController)
	_ = container.Provide(book.NewControllerRoute)
	_ = container.Provide(bookService.NewService)
	_ = container.Provide(bookRepository.NewRepository)
}