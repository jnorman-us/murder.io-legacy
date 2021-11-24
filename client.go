package main

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/events"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
	"time"
)

var done chan struct{}

func main() {
	var eventsManager = events.NewManager()
	var world = world.NewWorld()

	var gameEngine = engine.NewEngine(world, eventsManager)

	var wineCraft = entities.NewPlayer()
	wineCraft.SetPosition(types.NewVector(50, 100))
	wineCraft.Add(world, eventsManager)
	wineCraft.SetAngularVelocity(.1)
	wineCraft.AddInputListener()

	var zhaohang12345 = entities.NewPlayer()
	zhaohang12345.SetPosition(types.NewVector(350, 250))
	zhaohang12345.Add(world, eventsManager)

	var gameDrawer = drawer.NewDrawer(world, gameEngine, 500, 500)
	var _ = input.NewInput(eventsManager, "Wine_Craft")

	go gameDrawer.Start()
	go eventsManager.Start()

	for range time.Tick(time.Second * 1) {

	}
}
