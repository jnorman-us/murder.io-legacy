package main

import (
	"context"
	"fmt"
	"github.com/josephnormandev/murder/client/match"
	"github.com/josephnormandev/murder/common/types"
	"golang.org/x/sync/errgroup"
	"syscall/js"
	"time"
)

var manager *match.Manager
var runGroup *errgroup.Group
var runContext context.Context

func main() {
	js.Global().Set("connectToServer", js.FuncOf(connectToServer))
	//js.Global().Set("setInputs", js.FuncOf(setInputs))
	// js.Global().Set("getObjectData", js.FuncOf(getObjectData))

	manager = match.NewManager()

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	runGroup, runContext = errgroup.WithContext(ctx)
	runGroup.Go(func() error {
		return manager.UpdateTick(runContext)
	})
	runGroup.Go(func() error {
		return manager.SteadyTick(runContext)
	})

	err := runGroup.Wait()
	fmt.Println("Runtime error!", err)
}

func connectToServer(this js.Value, values []js.Value) interface{} {
	var hostname = values[0].String()
	var port = values[1].Int()
	var username = types.UserID(values[2].String())

	runGroup.Go(func() error {
		return manager.Connect(runContext, hostname, port, username)
	})
	return nil
}
