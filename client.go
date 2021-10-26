package main

import (
	"fmt"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/world"
	"time"
)

var done chan struct{}

func main() {
	var gameWorld = world.NewWorld()
	var _ = engine.NewEngine(gameWorld)

	for _ = range time.Tick(50 * time.Millisecond) { // 20 tps
		fmt.Println("bro")
	}
	fmt.Println("yikles")
}

func run
