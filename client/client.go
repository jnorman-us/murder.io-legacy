package main

import (
	"fmt"
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/client/world"
	"github.com/josephnormandev/murder/client/ws"
	"github.com/josephnormandev/murder/common/engine"
	"syscall/js"
	"time"
)

var username string

var gameWorld *world.World
var gameEngine *engine.Engine
var gameDrawer *drawer.Drawer
var gameInputs *input.Manager
var gamePackets *ws.Manager
var wsClient *ws.Client

var logicMS = 25

func main() {
	js.Global().Set("connectToServer", js.FuncOf(connectToServer))
	gameEngine = engine.NewEngine()
	gameDrawer = drawer.NewDrawer()
	gamePackets = ws.NewManager()
	gameWorld = world.NewWorld(gameEngine, gameDrawer)

	var sizeable = input.Sizeable(gameDrawer)
	gameInputs = input.NewManager(&sizeable)

	var spawner = ws.Spawner(gameWorld)
	var inputsSystem = ws.System(gameInputs)
	var deletionsListener = ws.Listener(gameWorld)

	gamePackets.SetSpawner(&spawner)
	gamePackets.AddSystem(&inputsSystem)
	gamePackets.AddListener(&deletionsListener)

	go gameDrawer.Start(updatePhysics)

	for range time.Tick(1 * time.Second) {
		if username != "" {
			var centerable = gameWorld.GetCenterable(username)
			if centerable != nil {
				gameDrawer.SetCenterable(centerable)
			}
		}
	}
}

func connectToServer(this js.Value, values []js.Value) interface{} {
	username = values[0].String()

	wsClient = ws.NewClient(gamePackets, username)
	go (func() {
		err := wsClient.Connect()
		if err != nil {
			fmt.Printf("Error with WS! %v\n", err)
		}
	})()
	return nil
}

func updatePhysics(ms float64) {
	// gameEngine.UpdatePhysics(ms / float64(logicMS))
}
