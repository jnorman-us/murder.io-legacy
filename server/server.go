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
	gameInputs = input.NewManager()
	gamePackets = ws.NewManager()

	var listener = ws.Listener(gameInputs)
	gamePackets.AddListener(&listener)

	wsServer = ws.NewServer(names, gamePackets)

	gameWorld = world.NewWorld(gameEngine, gameLogic, gameCollisions, gamePackets, gameInputs)

	for _, name := range names {
		var player = innocent.NewInnocent(name)
		player.SetPosition(types.NewRandomVector(0, 0, 600, 600))
		gameWorld.AddInnocent(player)
	}

	for i := 0; i < 1; i++ {
		var border = wall.NewWall(rand.Intn(1000))
		border.SetPosition(types.NewRandomVector(0, 0, 600, 600))
		border.SetAngle(rand.Float64() * math.Pi * 2)
		gameWorld.AddWall(border)
	}

	go tick()
	go wsServer.Send()

	var staticFiles = http.FileServer(http.Dir("./server/static"))

	var router = mux.NewRouter().StrictSlash(true)
	router.Handle("/ws/{id}", wsServer)
	router.PathPrefix("/").Handler(staticFiles)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func tick() {
	for range time.Tick(50 * time.Millisecond) {
		gameEngine.UpdatePhysics(1)
		gameCollisions.ResolveCollisions()
		gameLogic.Tick()
	}
}
