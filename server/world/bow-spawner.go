package world

import (
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/arrow"
	"github.com/josephnormandev/murder/common/entities/bow"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/server/ws"
)

func (w *World) AddBow(b *bow.Bow) int {
	var id = w.NextAvailableID()
	b.SetID(id)

	var spawner = bow.Spawner(w)
	b.SetSpawner(&spawner)

	var spawn = ws.Spawn(b)
	var tickable = logic.Tickable(b)
	var moveable = engine.Moveable(b)

	w.network.AddSpawn(id, &spawn)
	w.logic.AddTickable(id, &tickable)
	w.engine.AddMoveable(id, &moveable)
	w.Bows[id] = b
	return id
}

func (w *World) RemoveBow(id int) {
	delete(w.Bows, id)
	w.network.RemoveSpawn(id)
	w.logic.RemoveTickable(id)
	w.engine.RemoveMoveable(id)
}

func (w *World) SpawnArrow(h *bow.Holder, charge float64) {
	var shooter = (*h).(arrow.Shooter)
	var a = arrow.NewArrow(&shooter, charge)

	w.AddArrow(a)
}
