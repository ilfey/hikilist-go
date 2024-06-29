package main

import (
	"github.com/ilfey/hikilist-go/internal/app"
	"github.com/ilfey/hikilist-go/internal/config"
	"github.com/ilfey/hikilist-go/server"
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