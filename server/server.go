package main

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/wall"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/packet"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var gameWorld *world.World
var gameLogic *logic.Manager
var gameEngine *engine.Engine
var gameNetwork *packet.Manager
var gameCollisions *collisions.Manager

var logicMS = 33

func main() {
	gameLogic = logic.NewManager()
	gameEngine = engine.NewEngine()
	gameCollisions = collisions.NewManager()
	gameNetwork = packet.NewManager("**SERVER**")

	gameWorld = world.NewServerWorld(gameEngine, gameLogic, gameCollisions, gameNetwork)

	var wineCraft = innocent.NewInnocent("Wine_Craft")
	wineCraft.SetPosition(types.NewVector(250, 250))
	wineCraft.SetAngularVelocity(.1)
	//wineCraft.SetVelocity(types.NewVector(10, 0))

	gameWorld.AddInnocent(wineCraft)

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
		player.SetPosition(types.NewRandomVector(0, 0, 600, 600))
		gameWorld.AddInnocent(player)
	}

	for i := 0; i < 5; i++ {
		var border = wall.NewWall(rand.Intn(1000))
		border.SetPosition(types.NewRandomVector(0, 0, 600, 600))
		border.SetAngle(rand.Float64() * math.Pi * 2)
		gameWorld.AddWall(border)
	}

	go tick()

	fs := http.FileServer(http.Dir("./server/static"))
	http.Handle("/", fs)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func tick() {
	for range time.Tick(time.Duration(logicMS) * time.Millisecond) {
		gameEngine.UpdatePhysics(1)
		gameCollisions.ResolveCollisions()
		gameLogic.Tick()
	}
}
