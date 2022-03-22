package match

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/server/input"
	"github.com/josephnormandev/murder/server/worldout"
)

func (m *Match) SpawnDimetrodon(d *dimetrodon.Dimetrodon) types.ID {
	var id = m.entityID
	m.entityID++

	var spawner = dimetrodon.Spawner(m)
	var shootable = logic.Shootable(d)
	var inputable = input.Inputable(d)
	var driveable = logic.Driveable(d)
	var moveable = engine.Moveable(&d.Collider)
	var dynamic = collisions.Dynamic(d)
	var datafiable = worldout.Datafiable(d)

	d.SetSpawner(&spawner)
	m.logic.AddDriveable(id, &driveable)
	m.logic.AddShootable(id, &shootable)
	m.engine.AddMoveable(id, &moveable)
	m.inputs.AddPlayerListener(id, &inputable)
	m.collisions.AddDynamic(id, &dynamic)
	m.worldOut.AddDatafiable(id, &datafiable)

	d.ID = id
	return id
}

func (m *Match) DespawnDimetrodon(id types.ID) {
	m.logic.RemoveDriveable(id)
	m.logic.RemoveShootable(id)
	m.engine.RemoveMoveable(id)
	m.inputs.RemovePlayerListener(id)
	m.collisions.RemoveDynamic(id)
	m.worldOut.RemoveDatafiable(id)
}

func (m *Match) SpawnPole(p *pole.Pole) types.ID {
	var id = m.entityID
	m.entityID++

	var static = collisions.Static(p)
	var datafiable = worldout.Datafiable(p)

	m.collisions.AddStatic(id, &static)
	m.worldOut.AddDatafiable(id, &datafiable)

	p.ID = id
	return id
}

func (m *Match) DespawnPole(id types.ID) {
	m.collisions.RemoveStatic(id)
	m.worldOut.RemoveDatafiable(id)
}

func (m *Match) SpawnBullet(b *bullet.Bullet) types.ID {
	var id = m.entityID
	m.entityID++

	var spawner = bullet.Spawner(m)
	var dissolvable = logic.Dissolvable(b)
	var moveable = engine.Moveable(&b.Collider)
	var dynamic = collisions.Dynamic(b)
	var datafiable = worldout.Datafiable(b)

	b.SetSpawner(&spawner)
	m.logic.AddDissolvable(id, &dissolvable)
	m.engine.AddMoveable(id, &moveable)
	m.collisions.AddDynamic(id, &dynamic)
	m.worldOut.AddDatafiable(id, &datafiable)

	b.ID = id
	return id
}

func (m *Match) DespawnBullet(id types.ID) {
	m.logic.RemoveDissolvable(id)
	m.engine.RemoveMoveable(id)
	m.collisions.RemoveDynamic(id)
	m.worldOut.RemoveDatafiable(id)
}

func (m *Match) DisableBullet(id types.ID) {
	m.logic.RemoveDissolvable(id)
	m.engine.RemoveMoveable(id)
	m.collisions.RemoveDynamic(id)
}
