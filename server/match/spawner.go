package match

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
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
	var dynamicStatic = collisions.DynamicStatic(d)

	m.packets.AddSpawn(id, &spawn)
	// m.logic.AddTickable(id, &tickable)
	m.logic.AddDriveable(id, &driveable)
	m.logic.AddShootable(id, &shootable)
	m.engine.AddMoveable(id, &moveable)
	m.inputs.AddPlayerListener(id, &inputable)
	m.collisions.AddDynamicStatic(id, &dynamicStatic)

	d.SetSpawner(&spawner)
	d.ID = id
	return id
}

func (m *Match) DespawnDrifter(id types.ID) {
	m.packets.RemoveSpawn(id)
	// m.logic.RemoveTickable(id)
	m.logic.RemoveDriveable(id)
	m.engine.RemoveMoveable(id)
	m.inputs.RemovePlayerListener(id)
	m.collisions.RemoveDynamicStatic(id)
}

func (m *Match) SpawnPole(p *pole.Pole) types.ID {
	var id = m.entityID
	m.entityID++

	var spawn = ws.Spawn(p)
	var staticDynamic = collisions.StaticDynamic(p)

	m.packets.AddSpawn(id, &spawn)
	m.collisions.AddStaticDynamic(id, &staticDynamic)

	p.ID = id
	return id
}

func (m *Match) DespawnPole(id types.ID) {
	m.packets.RemoveSpawn(id)
	m.collisions.RemoveStaticDynamic(id)
}

func (m *Match) SpawnBullet(b *bullet.Bullet) types.ID {
	var id = m.entityID
	m.entityID++

	var spawn = ws.Spawn(b)
	// var fireable = logic.Fireable(b)
	var moveable = engine.Moveable(b)
	// var bulletShootable = collisions.BulletShootable(b)

	m.packets.AddSpawn(id, &spawn)
	// m.logic.AddFireable(id, &fireable)
	m.engine.AddMoveable(id, &moveable)
	// m.collisions.AddBulletShootable(id, &bulletShootable)

	b.ID = id
	return id
}

func (m *Match) DespawnBullet(id types.ID) {
	m.packets.RemoveSpawn(id)
	// m.logic.RemoveFireable(id)
	m.engine.RemoveMoveable(id)
	// m.collisions.RemoveBulletShootable(id)
}
