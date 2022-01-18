package main

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/client/tcp"
	"github.com/josephnormandev/murder/client/world"
	"github.com/josephnormandev/murder/common/engine"
)

var gameWorld *world.World
var gameEngine *engine.Engine
var gameDrawer *drawer.Drawer
var gameInputs *input.Manager
var gamePackets *tcp.Manager
var udpClient *tcp.Client

var logicMS = 50

func main() {
	gameEngine = engine.NewEngine()
	gameDrawer = drawer.NewDrawer()
	gamePackets = tcp.NewManager("Wine_Craft")
	udpClient = tcp.NewClient(gamePackets)

	var sizeable = input.Sizeable(gameDrawer)
	gameInputs = input.NewManager(&sizeable)
	var inputsSystem = tcp.System(gameInputs)
	gamePackets.AddSystem(&inputsSystem)

	gameWorld = world.NewWorld(gameEngine, gameDrawer, gameInputs)

	go udpClient.Send()
	go udpClient.Listen()
	go gameDrawer.Start(updatePhysics)
}

func updatePhysics(ms float64) {
	gameEngine.UpdatePhysics(ms / float64(logicMS))
}
