package main

import (
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
	"github.com/josephnormandev/murder/drawer"
	"time"
)

var done chan struct{}

func main() {
	var gameWorld = world.NewWorld()
	var gameEngine = engine.NewEngine(gameWorld)

	var wineCraft = entities.NewPlayer()
	wineCraft.SetPosition(types.NewVector(50, 100))
	wineCraft.SetAngularVelocity(0.01)
	wineCraft.SetVelocity(types.NewVector(.5, 0))
	wineCraft.AddTo(gameWorld)

	var Zhaohang12345 = entities.NewPlayer()
	Zhaohang12345.SetPosition(types.NewVector(450, 100))
	Zhaohang12345.SetAngularVelocity(0.01)
	Zhaohang12345.SetVelocity(types.NewVector(-.5, 0))
	Zhaohang12345.AddTo(gameWorld)

	var gameDrawer = drawer.NewDrawer(gameWorld, gameEngine, 500, 500)

	gameDrawer.Start()

	for range time.Tick(time.Second * 1) {

	}
}
