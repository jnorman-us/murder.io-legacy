package main

import (
	"github.com/josephnormandev/murder/client/match"
	"github.com/josephnormandev/murder/common/types"
	"syscall/js"
	"time"
)

var manager *match.Manager

func main() {
	js.Global().Set("connectToServer", js.FuncOf(connectToServer))
	js.Global().Set("setInputs", js.FuncOf(setInputs))
	js.Global().Set("getObjectData", js.FuncOf(getObjectData))

	manager = match.NewManager()
	go manager.Update()
	go manager.SteadyTick()

	for range time.Tick(1 * time.Second) {
	}
}

func connectToServer(this js.Value, values []js.Value) interface{} {
	var hostname = values[0].String()
	var port = values[1].Int()
	var username = types.UserID(values[2].String())

	manager.Connect(hostname, port, username)
	return nil
}

func setInputs(this js.Value, values []js.Value) interface{} {
	return nil
}

func getObjectData(this js.Value, values []js.Value) interface{} {
	return "test"
}
