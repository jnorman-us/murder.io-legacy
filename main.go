package main

import (
	"fmt"
	"github.com/josephnormandev/murder/drawer"
	"github.com/josephnormandev/murder/engine"
	"github.com/josephnormandev/murder/entities"
	"github.com/josephnormandev/murder/types"
	"github.com/josephnormandev/murder/world"
	"time"
)

func main() {
	var gameWorld = world.NewWorld()

	var wineCraft = entities.NewActor("Wine_Craft")
	var jurgmania = entities.NewActor("Jurgmania")
	var walls = entities.NewWalls([][]bool{
		{true, true, true, true, true, true, true, true, true, true},
		{true, false, false, false, false, false, false, false, false, true},
		{true, false, false, false, false, false, false, false, false, true},
		{true, false, false, false, false, false, false, false, false, true},
		{true, false, false, false, false, false, false, false, false, true},
		{true, false, false, false, false, false, false, false, false, true},
		{true, false, false, false, false, false, false, false, false, true},
		{true, false, false, false, false, false, false, false, false, true},
		{true, false, false, false, false, false, false, false, false, true},
		{true, false, false, false, false, false, false, false, false, true},
		{true, true, true, true, true, true, true, true, true, true},
	})

	walls.AddTo(gameWorld)
	walls.SetPosition(types.NewVector(15, 35))
	walls.SetVelocity(types.NewVector(.1, -.1))

	wineCraft.AddTo(gameWorld)
	wineCraft.SetPosition(types.NewVector(20, 20))
	//wineCraft.SetAngle(30.0 / 180.0 * math.Pi)
	wineCraft.SetVelocity(types.NewVector(.1, .0))
	//wineCraft.SetAngularVelocity(.02)

	jurgmania.AddTo(gameWorld)
	jurgmania.SetPosition(types.NewVector(22, 20))
	//jurgmania.SetAngle(-100.0 / 180.0 * math.Pi)
	//jurgmania.SetVelocity(types.NewVector(0, 0))
	//jurgmania.SetAngularVelocity(-.03)

	var gameEngine = engine.NewEngine(gameWorld)
	var gameDrawer = drawer.NewDrawer(gameWorld)

	go gameEngine.Start()
	go gameDrawer.Start()

	time.Sleep(5 * time.Second)

	//wineCraft.RemoveFrom()

	time.Sleep(50 * time.Second)

	gameEngine.Stop()
	gameDrawer.Stop()

	fmt.Println("finished")
}
