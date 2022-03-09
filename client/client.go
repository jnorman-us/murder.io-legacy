package main

import (
	"fmt"
	"github.com/josephnormandev/murder/client/match"
	"syscall/js"
)

var manager *match.Manager

func main() {
	manager = match.NewManager()
	manager.ExposeFunctions(js.Global())
	manager.UpdateTick()
	manager.SteadyTick()

	err := manager.RunGroup.Wait()
	fmt.Println("Runtime error!", err)
}
