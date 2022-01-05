package main

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/wall"
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

var logicMS = 33

func main() {
	gameEngine = engine.NewEngine()
	gameLogic = logic.NewManager()
	gameDrawer = drawer.NewDrawer()
	gameCollisions = collisions.NewManager()

	var sizeable = input.Sizeable(gameDrawer)
	gameInputs = input.NewManager(&sizeable)

	gameWorld = world.NewClientWorld(gameEngine, gameLogic, gameCollisions, gameDrawer, gameInputs)

	var wineCraft = innocent.NewInnocent("Wine_Craft")
	wineCraft.SetPosition(types.NewVector(250, 250))
	wineCraft.SetAngularVelocity(.1)
	//wineCraft.SetVelocity(types.NewVector(10, 0))

	gameWorld.AddInnocent(wineCraft)

	wineCraft.AddInputs(gameInputs)

	for _, name := range []string{
		"Xiehang",
		"TheStorminNorman",
		"ShadowDragon",
		"Society Member",
		"Envii",
		"Jinseng",
		"Laerir",
		"JoeyD",
	} {
		var player = innocent.NewInnocent(name)
		player.SetPosition(types.NewRandomVector(0, 0, 500, 500))
		gameWorld.AddInnocent(player)
	}

	var border = wall.NewWall(100)
	border.SetPosition(types.NewVector(100, 200))

	var center = drawer.Centerable(wineCraft)
	gameDrawer.SetCenterable(&center)

	gameWorld.AddWall(border)

	go tick()
	go gameDrawer.Start(updatePhysics)

	for range time.Tick(time.Second) {
		// do nothing, just keep this thread alive...
	}
}

func updatePhysics(ms float64) {
	gameEngine.UpdatePhysics(ms / float64(logicMS))
	gameCollisions.ResolveCollisions()
}

func tick() {
	for range time.Tick(time.Duration(logicMS) * time.Millisecond) {
		gameLogic.Tick()
	}
}
