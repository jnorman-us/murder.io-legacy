package main

import (
	"fmt"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/world"
	"time"
)

func main() {
	var gameWorld = world.NewWorld()
	var gameEngine = engine.NewEngine(gameWorld)

	gameEngine.Start()

	time.Sleep(50000)

	gameEngine.Stop()

	fmt.Println("finished")
}
