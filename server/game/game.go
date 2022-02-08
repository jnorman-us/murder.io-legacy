package game

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/game"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/server/input"
	"github.com/josephnormandev/murder/server/ws"
)

type ServerGame struct {
	types.ID
	game.Game
	logic      *logic.Manager
	engine     *engine.Engine
	packets    *ws.Lobby
	inputs     *input.Manager
	collisions *collisions.Manager
}

func NewServerGame(id types.ID) *ServerGame {
	var game = &ServerGame{
		ID:   id,
		Game: *game.NewGame(),
	}
	var packetsInfo = ws.LobbyInfo(game)

	var gLogic = logic.NewManager()
	var gEngine = engine.NewEngine()
	var packets = ws.NewLobby(&packetsInfo)
	var inputs = input.NewManager()
	var gCollisions = collisions.NewManager()

	var inputListener = ws.Listener(inputs)
	packets.AddListener(&inputListener)

	var positionsSystem = ws.System(gEngine)
	var gameSystem = ws.System(game)
	packets.AddSystem(&positionsSystem)
	packets.AddSystem(&gameSystem)

	game.logic = gLogic
	game.engine = gEngine
	game.packets = packets
	game.inputs = inputs
	game.collisions = gCollisions
	return game
}
