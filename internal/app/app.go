package app

import (
	"go.uber.org/dig"
	"<package_name>/api/controllers/book"
	"<package_name>/api/router"
	"<package_name>/internal/config"
	bookRepository "<package_name>/internal/repositories/book"
	bookService "<package_name>/internal/services/book"
	"<package_name>/server"
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