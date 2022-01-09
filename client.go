package main

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/dummy"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/wall"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/packet"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
	"math"
	"math/rand"
	"time"
)

var gameWorld *world.World
var gameEngine *engine.Engine
var gameLogic *logic.Manager
var gameCollisions *collisions.Manager
var gameDrawer *drawer.Drawer
var gameInputs *input.Manager
var gameNetwork *packet.Manager

var logicMS = 33

func main() {
	gameEngine = engine.NewEngine()
	gameLogic = logic.NewManager()
	gameDrawer = drawer.NewDrawer()
	gameCollisions = collisions.NewManager()
	gameNetwork = packet.NewManager("Wine_Craft")

	var sizeable = input.Sizeable(gameDrawer)
	gameInputs = input.NewManager(&sizeable)

	gameWorld = world.NewClientWorld(gameEngine, gameLogic, gameCollisions, gameDrawer, gameInputs)

	var wineCraft = innocent.NewInnocent("Wine_Craft")
	var center = drawer.Centerable(wineCraft)
	wineCraft.SetPosition(types.NewVector(250, 250))
	wineCraft.SetAngularVelocity(.1)
	wineCraft.AddInputs(gameInputs)
	//wineCraft.SetVelocity(types.NewVector(10, 0))

	gameWorld.AddInnocent(wineCraft)
	gameDrawer.SetCenterable(&center)
	var inputsSystem = packet.System(gameInputs)
	gameNetwork.AddSystem(inputsSystem.GetChannel(), &inputsSystem)

	var dummyListener = &dummy.Listener{}
	var inputListener = packet.Listener(dummyListener)
	gameNetwork.AddListener(inputListener.GetChannel(), &inputListener)

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
	go gameDrawer.Start(updatePhysics)

	for range time.Tick(time.Second) {
		gameNetwork.EncodeOutputs()
		gameNetwork.CopyOver()
		gameNetwork.DecodeInputs()
	}
}

func updatePhysics(ms float64) {
	gameEngine.UpdatePhysics(ms / float64(logicMS))
}

func tick() {
	for range time.Tick(time.Duration(logicMS) * time.Millisecond) {
		gameCollisions.ResolveCollisions()
		gameLogic.Tick()
	}
}
