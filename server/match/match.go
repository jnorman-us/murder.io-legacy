package match

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/game"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
	"github.com/josephnormandev/murder/server/input"
	"github.com/josephnormandev/murder/server/worldout"
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

	sync.Mutex
	entityID   types.ID
	logic      *logic.Manager
	engine     *engine.Engine
	lobby      *ws.Lobby
	inputs     *input.Manager
	collisions *collisions.Manager
	worldOut   *worldout.Manager
}

func NewMatch(id types.ID) *Match {
	var m = &Match{
		ID:       id,
		entityID: 1,
		Game:     *game.NewGame(),
	}
	var spawner = world.Spawner(m)
	m.World = *world.NewWorld(&spawner)

	var packetsInfo = ws.LobbyInfo(m)

	var gLogic = logic.NewManager()
	var gEngine = engine.NewEngine()
	var lobby = ws.NewLobby(&packetsInfo, packets.NewManager())
	var inputs = input.NewManager()
	var gCollisions = collisions.NewManager()
	var worldOut = worldout.NewManager(&lobby.Manager)

	m.logic = gLogic
	m.engine = gEngine
	m.lobby = lobby
	m.inputs = inputs
	m.collisions = gCollisions
	m.worldOut = worldOut
	return m
}

func (m *Match) GetLobby() *ws.Lobby {
	return m.lobby
}

func (m *Match) Tick() {
	for range time.Tick(tickTime) {
		m.tick()
	}
}

func (m *Match) tick() {
	m.Lock()
	defer m.Unlock()
	m.logic.Tick()
	m.engine.UpdatePhysics(1)
	m.collisions.ResolveCollisions()
	m.worldOut.PollData()
}

func (m *Match) Send() {
	for range time.Tick(sendTime) {
		m.send()
	}
}

func (m *Match) send() {
	m.Lock()
	defer m.Unlock()
	m.lobby.Send()
}
