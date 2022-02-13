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
	"time"
)

type Match struct {
	types.ID
	world.World // keeps track of Spawns and Entities
	game.Game   // keeps track of Game state
	entityID    types.ID
	logic       *logic.Manager
	engine      *engine.Engine
	packets     *ws.Lobby
	inputs      *input.Manager
	collisions  *collisions.Manager
}

func NewMatch(id types.ID) *Match {
	var match = &Match{
		ID:       id,
		entityID: 1,
		Game:     *game.NewGame(),
	}
	var spawner = world.Spawner(match)
	match.World = *world.NewWorld(&spawner)

	var packetsInfo = ws.LobbyInfo(match)

	var gLogic = logic.NewManager()
	var gEngine = engine.NewEngine()
	var packets = ws.NewLobby(&packetsInfo)
	var inputs = input.NewManager()
	var gCollisions = collisions.NewManager()

	var inputListener = ws.Listener(inputs)
	packets.AddListener(&inputListener)

	var positionsSystem = ws.System(gEngine)
	var deletionsSystem = ws.System(match.Deletions())
	var gameSystem = ws.System(match)
	packets.AddSystem(&positionsSystem)
	packets.AddSystem(&deletionsSystem)
	packets.AddSystem(&gameSystem)

	match.logic = gLogic
	match.engine = gEngine
	match.packets = packets
	match.inputs = inputs
	match.collisions = gCollisions
	return match
}

func (m *Match) GetPackets() *ws.Lobby {
	return m.packets
}

func (m *Match) Tick() {
	for range time.Tick(25 * time.Millisecond) {
		m.logic.Tick()
		m.engine.UpdatePhysics(1)
		m.collisions.ResolveCollisions()
	}
}
