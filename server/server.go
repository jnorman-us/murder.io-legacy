package main

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/wall"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/server/websocket"
	"github.com/josephnormandev/murder/server/world"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var gameWorld *world.World
var gameLogic *logic.Manager
var gameEngine *engine.Engine
var gamePackets *websocket.Manager
var gameCollisions *collisions.Manager

var logicMS = 33

var names = []string{
	"Wine_Craft",
	"Xiehang",
	"TheStorminNorman",
	"ShadowDragon",
	"Society Member",
	"Envii",
	"Jinseng",
	"Laerir",
	"JoeyD",
}

func main() {
	gameLogic = logic.NewManager()
	gameEngine = engine.NewEngine()
	gameCollisions = collisions.NewManager()

	gamePackets = websocket.NewManager()

	gameWorld = world.NewWorld(gameEngine, gameLogic, gameCollisions, gamePackets)

	for _, name := range names {
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
