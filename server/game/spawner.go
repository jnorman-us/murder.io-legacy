package game

import (
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/server/input"
	"github.com/josephnormandev/murder/server/ws"
)

func (g *ServerGame) SpawnDrifter(d *drifter.Drifter) types.ID {
	var id = g.entityID
	g.entityID++

	var spawn = ws.Spawn(d)
	var tickable = logic.Tickable(d)
	var inputable = input.Inputable(d)
	var driveable = logic.Driveable(d)
	var moveable = engine.Moveable(d)

	g.Game.Drifters[id] = d
	g.packets.AddSpawn(id, &spawn)
	g.logic.AddTickable(id, &tickable)
	g.logic.AddDriveable(id, &driveable)
	g.engine.AddMoveable(id, &moveable)
	g.inputs.AddPlayerListener(id, &inputable)

	d.ID = id
	return id
}

func (g *ServerGame) DespawnDrifter(id types.ID) {
	g.packets.RemoveSpawn(id)
	g.logic.RemoveTickable(id)
	g.logic.RemoveDriveable(id)
	g.engine.RemoveMoveable(id)
	g.inputs.RemovePlayerListener(id)
	delete(g.Game.Drifters, id)
}
