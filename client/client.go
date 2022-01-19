package main

import (
	"fmt"
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/client/world"
	"github.com/josephnormandev/murder/client/ws"
	"github.com/josephnormandev/murder/common/engine"
	"time"
)

var gameWorld *world.World
var gameEngine *engine.Engine
var gameDrawer *drawer.Drawer
var gameInputs *input.Manager
var gamePackets *ws.Manager
var wsClient *ws.Client

var logicMS = 50

func main() {
	gameEngine = engine.NewEngine()
	gameDrawer = drawer.NewDrawer()
	gamePackets = ws.NewManager()

	wsClient = ws.NewClient(gamePackets, "Wine_Craft")

	var sizeable = input.Sizeable(gameDrawer)
	gameInputs = input.NewManager(&sizeable)
	var inputsSystem = ws.System(gameInputs)
	gamePackets.AddSystem(&inputsSystem)

	gameWorld = world.NewWorld(gameEngine, gameDrawer, gameInputs)

	go gameDrawer.Start(updatePhysics)
	go (func() {
		err := wsClient.Connect()
		if err != nil {
			fmt.Printf("Error with WS! %v\n", err)
		}
	})()

	for range time.Tick(1 * time.Second) {

	}
}

func updatePhysics(ms float64) {
	gameEngine.UpdatePhysics(ms / float64(logicMS))
}
