package main

import (
	"github.com/ilfey/hikilist-go/internal/app"
	config2 "github.com/ilfey/hikilist-go/internal/config"
	"github.com/ilfey/hikilist-go/internal/domain/service/di"
	"sync"
)

func main() {
	config2.MustLoadEnvironment()

	cfg := config2.New()

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go app.NewApp(
		cfg,
		di.NewServiceContainerManager(),
	).Run(wg)

	wg.Wait()
}
