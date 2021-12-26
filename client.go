package main

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/sword"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
	"time"
)

var gameWorld *world.World
var gameEngine *engine.Engine
var gameLogic *logic.Manager
var gameCollisions *collisions.Manager
var gameDrawer *drawer.Drawer
var gameInputs *input.Manager

func main() {
	gameEngine = engine.NewEngine()
	gameLogic = logic.NewManager()
	gameDrawer = drawer.NewDrawer(500, 500)
	gameInputs = input.NewManager("WineCraft")
	gameCollisions = collisions.NewManager()
	gameWorld = world.NewClientWorld(gameEngine, gameLogic, gameCollisions, gameDrawer, gameInputs)

	var wineCraft = innocent.NewInnocent()
	wineCraft.SetPosition(types.NewVector(50, 100))
	wineCraft.SetAngularVelocity(.1)
	wineCraft.SetVelocity(types.NewVector(10, 0))
	wineCraft.SetAngularFriction(.01)
	wineCraft.SetFriction(.5)

	wineCraft.AddInputs(gameInputs)

	var xiehang = innocent.NewInnocent()
	xiehang.SetPosition(types.NewVector(80, 100))
	xiehang.SetVelocity(types.NewVector(0, 10))
	xiehang.SetAngularVelocity(-.175)
	xiehang.SetAngularFriction(.1)
	xiehang.SetFriction(.5)

	var sword = sword.NewSword()
	sword.SetPosition(types.NewVector(300, 100))
	sword.SetAngularVelocity(.1)

	gameWorld.AddInnocent(wineCraft)
	gameWorld.AddInnocent(xiehang)
	gameWorld.AddSword(sword)

	go gameDrawer.Start()

	for range time.Tick(50 * time.Millisecond) {
		tick()
	}
}

func tick() {
	gameLogic.Tick()
	gameEngine.UpdatePhysics()
	gameCollisions.ResolveCollisions()
}
