package main

import (
	"<package_name>/internal/app"
	"<package_name>/internal/config"
	"<package_name>/server"
)

func main() {
	// Load config
	config.LoadEnvironment()

	// Build container
	container := app.BuildContainer()

	// Run server
	err := container.Invoke(func(server *server.Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}