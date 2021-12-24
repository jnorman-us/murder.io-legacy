package main

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
	"time"
)

var gameWorld *world.World
var gameEngine *engine.Engine
var gameCollisions *collisions.Manager
var gameDrawer *drawer.Drawer
var gameInputs *input.Manager

func main() {
	gameEngine = engine.NewEngine()
	gameDrawer = drawer.NewDrawer(500, 500)
	gameInputs = input.NewManager("WineCraft")
	gameCollisions = collisions.NewManager()
	gameWorld = world.NewClientWorld(gameEngine, gameCollisions, gameDrawer, gameInputs)

	var wineCraft = innocent.NewInnocent()
	wineCraft.SetPosition(types.NewVector(50, 100))
	wineCraft.SetAngularVelocity(.1)
	wineCraft.AddInputs(gameInputs)

	var xiehang = innocent.NewInnocent()
	xiehang.SetPosition(types.NewVector(80, 100))
	xiehang.SetAngularVelocity(-.175)

	gameWorld.AddInnocent(wineCraft)
	gameWorld.AddInnocent(xiehang)

	go gameDrawer.Start()

	for range time.Tick(50 * time.Millisecond) {
		tick()
	}
}

func tick() {
	gameEngine.UpdatePhysics()
	gameCollisions.ResolveCollisions()
}
