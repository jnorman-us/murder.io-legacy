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
	events.InitializeEvents()
	world.InitializeWorld()

	var gameEngine = engine.NewEngine()

	var wineCraft = entities.NewPlayer()
	wineCraft.SetPosition(types.NewVector(50, 100))
	wineCraft.Add()
	wineCraft.SetAngularVelocity(.3)
	wineCraft.AddInputListener()

	var zhaohang12345 = entities.NewPlayer()
	zhaohang12345.SetPosition(types.NewVector(450, 100))
	zhaohang12345.Add()

	var gameDrawer = drawer.NewDrawer(gameEngine, 500, 500)
	var _ = input.NewInput("Wine_Craft")

	go gameDrawer.Start()

	for range time.Tick(time.Second * 1) {

	}
}
