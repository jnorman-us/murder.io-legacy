package main

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/client/websocket"
	"github.com/josephnormandev/murder/client/world"
	"github.com/josephnormandev/murder/common/engine"
	"time"
)

var gameWorld *world.World
var gameEngine *engine.Engine
var gameDrawer *drawer.Drawer
var gameInputs *input.Manager
var gamePackets *websocket.Manager

var logicMS = 50

func main() {
	gameEngine = engine.NewEngine()
	gameDrawer = drawer.NewDrawer()
	gamePackets = websocket.NewManager("Wine_Craft")

	var sizeable = input.Sizeable(gameDrawer)
	gameInputs = input.NewManager(&sizeable)
	var inputsSystem = websocket.System(gameInputs)
	gamePackets.AddSystem(&inputsSystem)

	gameWorld = world.NewWorld(gameEngine, gameDrawer, gameInputs)

	go gameDrawer.Start(updatePhysics)

	for range time.Tick(1 * time.Second) {

	}
}

func updatePhysics(ms float64) {
	gameEngine.UpdatePhysics(ms / float64(logicMS))
}
