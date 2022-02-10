package match

import (
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/server/input"
	"github.com/josephnormandev/murder/server/ws"
)

func (m *Match) SpawnDrifter(d *drifter.Drifter) types.ID {
	var id = m.entityID
	m.entityID++

	var spawn = ws.Spawn(d)
	var tickable = logic.Tickable(d)
	var inputable = input.Inputable(d)
	var driveable = logic.Driveable(d)
	var moveable = engine.Moveable(d)

	m.packets.AddSpawn(id, &spawn)
	m.logic.AddTickable(id, &tickable)
	m.logic.AddDriveable(id, &driveable)
	m.engine.AddMoveable(id, &moveable)
	m.inputs.AddPlayerListener(id, &inputable)

	d.ID = id
	return id
}

func (m *Match) DespawnDrifter(id types.ID) {
	m.packets.RemoveSpawn(id)
	m.logic.RemoveTickable(id)
	m.logic.RemoveDriveable(id)
	m.engine.RemoveMoveable(id)
	m.inputs.RemovePlayerListener(id)
}
