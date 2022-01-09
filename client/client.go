package main

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/packet"
	"github.com/josephnormandev/murder/common/world"
	"time"
)

var gameWorld *world.World
var gameEngine *engine.Engine
var gameDrawer *drawer.Drawer
var gameInputs *input.Manager
var gameNetwork *packet.Manager

var logicMS = 33

func main() {
	gameEngine = engine.NewEngine()
	gameDrawer = drawer.NewDrawer()
	gameNetwork = packet.NewManager("Wine_Craft")

	var sizeable = input.Sizeable(gameDrawer)
	gameInputs = input.NewManager(&sizeable)
	var inputsSystem = packet.System(gameInputs)
	gameNetwork.AddSystem(inputsSystem.GetChannel(), &inputsSystem)

	gameWorld = world.NewClientWorld(gameEngine, gameDrawer, gameInputs)

	go gameDrawer.Start(updatePhysics)

	for range time.Tick(time.Second) {
		gameNetwork.EncodeOutputs()
		gameNetwork.CopyOver()
		gameNetwork.DecodeInputs()
	}
}

func updatePhysics(ms float64) {
	gameEngine.UpdatePhysics(ms / float64(logicMS))
}
