package main

import (
	"context"
	"fmt"
	"github.com/josephnormandev/murder/client/match"
	"golang.org/x/sync/errgroup"
	"syscall/js"
)

var manager *match.Manager

func main() {
	manager = match.NewManager()

	runGroup, runContext := errgroup.WithContext(context.Background())

	runGroup.Go(func() error {
		return manager.UpdateTick(runContext)
	})
	runGroup.Go(func() error {
		return manager.SteadyTick(runContext)
	})

	manager.ExposeFunctions(js.Global(), runGroup, runContext)

	err := runGroup.Wait()
	fmt.Println("Runtime error!", err)
}
