package main

import (
	"github.com/gorilla/mux"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/wall"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/server/input"
	"github.com/josephnormandev/murder/server/world"
	"github.com/josephnormandev/murder/server/ws"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var gameWorld *world.World
var gameLogic *logic.Manager
var gameEngine *engine.Engine
var gamePackets *ws.Manager
var gameInputs *input.Manager
var gameCollisions *collisions.Manager
var wsServer *ws.Server

var names = []string{
	"Jellotinous",
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
	gamePackets = ws.NewManager()
	gameLogic = logic.NewManager()
	gameEngine = engine.NewEngine()
	gameInputs = input.NewManager()
	gameCollisions = collisions.NewManager()
	gameWorld = world.NewWorld(gameEngine, gameLogic, gameCollisions, gamePackets, gameInputs)

	var listener = ws.Listener(gameInputs)
	gamePackets.AddListener(&listener)

	var deletionsSystem = ws.System(gameWorld.Deletions)
	var positionsSystem = ws.System(gameEngine)
	gamePackets.AddSystem(&deletionsSystem)
	gamePackets.AddSystem(&positionsSystem)

	wsServer = ws.NewServer(names, gamePackets)

	for i := 0; i < 4; i++ {
		var border = wall.NewWall(rand.Intn(1000))
		border.SetPosition(types.NewRandomVector(0, 0, 600, 600))
		border.SetAngle(rand.Float64() * math.Pi * 2)
		gameWorld.AddWall(border)
	}

	go tick()
	go wsServer.Send()
	go playerReset()

	var staticFiles = http.FileServer(http.Dir("./server/static"))

	var router = mux.NewRouter().StrictSlash(true)
	router.Handle("/ws/{id}", wsServer)
	router.PathPrefix("/").Handler(staticFiles)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func playerReset() {
	for range time.Tick(500 * time.Millisecond) {
		if len(gameWorld.Innocents) <= 1 {
			gameWorld.ResetInnocents()
			for _, name := range names {
				var player = innocent.NewInnocent(name)
				player.SetPosition(types.NewRandomVector(0, 0, 600, 600))
				gameWorld.AddInnocent(player)
			}
		}
	}
}

func tick() {
	for range time.Tick(1000 / 40 * time.Millisecond) {
		gameEngine.UpdatePhysics(1)
		gameCollisions.ResolveCollisions()
		gameLogic.Tick()
	}
}
