package match

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/game"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
	"github.com/josephnormandev/murder/server/input"
	"github.com/josephnormandev/murder/server/ws"
	"sync"
	"time"
)

const tickTime = time.Millisecond * 1000 / 60
const sendTime = time.Millisecond * 1000 / 5

type Match struct {
	types.ID
	world.World // keeps track of Spawns and Entities
	game.Game   // keeps track of Game state

	worldLock  *sync.Mutex
	entityID   types.ID
	logic      *logic.Manager
	engine     *engine.Engine
	packets    *ws.Lobby
	inputs     *input.Manager
	collisions *collisions.Manager
}

func NewMatch(id types.ID) *Match {
	var m = &Match{
		ID:        id,
		entityID:  1,
		worldLock: &sync.Mutex{},
		Game:      *game.NewGame(),
	}
	var spawner = world.Spawner(m)
	m.World = *world.NewWorld(&spawner)

	var packetsInfo = ws.LobbyInfo(m)

	var gLogic = logic.NewManager()
	var gEngine = engine.NewEngine()
	var packets = ws.NewLobby(&packetsInfo)
	var inputs = input.NewManager()
	var gCollisions = collisions.NewManager()

	//var inputListener = ws.Listener(inputs)
	//packets.AddListener(&inputListener)

	//var gameSystem = ws.System(m.Game)
	//packets.AddSystem(&gameSystem)

	m.logic = gLogic
	m.engine = gEngine
	m.packets = packets
	m.inputs = inputs
	m.collisions = gCollisions
	return m
}

func (m *Match) GetPackets() *ws.Lobby {
	return m.packets
}

func (m *Match) GetWorldLock() *sync.Mutex {
	return m.worldLock
}

func (m *Match) Tick() {
	for range time.Tick(tickTime) {
		m.worldLock.Lock()
		m.logic.Tick()
		m.engine.UpdatePhysics(1)
		m.collisions.ResolveCollisions()
		m.worldLock.Unlock()
	}
}

func (m *Match) Send() {
	for range time.Tick(sendTime) {
		m.worldLock.Lock()
		m.packets.Send()
		m.worldLock.Unlock()
	}
}
