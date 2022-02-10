package main

import (
	"fmt"
	"github.com/josephnormandev/murder/client/match"
	"github.com/josephnormandev/murder/client/ws"
	"github.com/josephnormandev/murder/common/types"
	"syscall/js"
	"time"
)

var username string

var manager *match.Manager
var wsClient *ws.Client

func main() {
	js.Global().Set("connectToServer", js.FuncOf(connectToServer))

	manager = match.NewManager()
	go manager.Start()
	go manager.SteadyTick()

	for range time.Tick(1 * time.Second) {
	}
}

func connectToServer(this js.Value, values []js.Value) interface{} {
	var hostname = values[0].String()
	var port = values[1].Int()
	manager.Username = types.UserID(values[2].String())

	wsClient = ws.NewClient(manager.GetPackets(), hostname, port, manager.Username)
	go (func() {
		err := wsClient.Connect()
		if err != nil {
			fmt.Printf("Error with WS! %v\n", err)
		}
	})()
	return nil
}
