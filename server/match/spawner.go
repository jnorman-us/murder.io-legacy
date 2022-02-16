package match

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/server/input"
	"github.com/josephnormandev/murder/server/ws"
)

func (m *Match) SpawnDrifter(d *drifter.Drifter) types.ID {
	var id = m.entityID
	m.entityID++

	var spawner = drifter.Spawner(m)
	var spawn = ws.Spawn(d)
	// var tickable = logic.Tickable(d)
	var shootable = logic.Shootable(d)
	var inputable = input.Inputable(d)
	var driveable = logic.Driveable(d)
	var moveable = engine.Moveable(d)
	var dynamic = collisions.Dynamic(d)

	d.SetSpawner(&spawner)
	m.packets.AddSpawn(id, &spawn)
	m.logic.AddDriveable(id, &driveable)
	m.logic.AddShootable(id, &shootable)
	m.engine.AddMoveable(id, &moveable)
	m.inputs.AddPlayerListener(id, &inputable)
	m.collisions.AddDynamic(id, &dynamic)

	d.ID = id
	return id
}

func (m *Match) DespawnDrifter(id types.ID) {
	m.packets.RemoveSpawn(id)
	m.logic.RemoveDriveable(id)
	m.logic.RemoveShootable(id)
	m.engine.RemoveMoveable(id)
	m.inputs.RemovePlayerListener(id)
	m.collisions.RemoveDynamic(id)
}

func (m *Match) SpawnDimetrodon(d *dimetrodon.Dimetrodon) types.ID {
	var id = m.entityID
	m.entityID++

	var spawner = dimetrodon.Spawner(m)
	var spawn = ws.Spawn(d)
	var shootable = logic.Shootable(d)
	var inputable = input.Inputable(d)
	var driveable = logic.Driveable(d)
	var moveable = engine.Moveable(d)
	var dynamic = collisions.Dynamic(d)

	d.SetSpawner(&spawner)
	m.packets.AddSpawn(id, &spawn)
	m.logic.AddDriveable(id, &driveable)
	m.logic.AddShootable(id, &shootable)
	m.engine.AddMoveable(id, &moveable)
	m.inputs.AddPlayerListener(id, &inputable)
	m.collisions.AddDynamic(id, &dynamic)

	d.ID = id
	return id
}

func (m *Match) DespawnDimetrodon(id types.ID) {
	m.packets.RemoveSpawn(id)
	m.logic.RemoveDriveable(id)
	m.logic.RemoveShootable(id)
	m.engine.RemoveMoveable(id)
	m.inputs.RemovePlayerListener(id)
	m.collisions.RemoveDynamic(id)
}

func (m *Match) SpawnPole(p *pole.Pole) types.ID {
	var id = m.entityID
	m.entityID++

	var spawn = ws.Spawn(p)
	var static = collisions.Static(p)

	m.packets.AddSpawn(id, &spawn)
	m.collisions.AddStatic(id, &static)

	p.ID = id
	return id
}

func (m *Match) DespawnPole(id types.ID) {
	m.packets.RemoveSpawn(id)
	m.collisions.RemoveStatic(id)
}

func (m *Match) SpawnBullet(b *bullet.Bullet) types.ID {
	var id = m.entityID
	m.entityID++

	var spawner = bullet.Spawner(m)
	var spawn = ws.Spawn(b)
	var dissolvable = logic.Dissolvable(b)
	var moveable = engine.Moveable(b)
	var dynamic = collisions.Dynamic(b)

	b.SetSpawner(&spawner)
	m.packets.AddSpawn(id, &spawn)
	m.logic.AddDissolvable(id, &dissolvable)
	m.engine.AddMoveable(id, &moveable)
	m.collisions.AddDynamic(id, &dynamic)

	b.ID = id
	return id
}

func (m *Match) DespawnBullet(id types.ID) {
	m.packets.RemoveSpawn(id)
	m.logic.RemoveDissolvable(id)
	m.engine.RemoveMoveable(id)
	m.collisions.RemoveDynamic(id)
}
