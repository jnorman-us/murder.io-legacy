package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/arrow"
	"github.com/josephnormandev/murder/common/entities/bow"
	"github.com/josephnormandev/murder/common/logic"
)

func (w *World) AddBow(b *bow.Bow) int {
	var id = w.NextAvailableID()
	var spawner = bow.Spawner(w)
	var tickable = logic.Tickable(b)
	var drawable = drawer.Drawable(b)
	var moveable = engine.Moveable(b)

	b.SetID(id)
	b.SetSpawner(&spawner)

	w.Bows[id] = b
	w.drawer.AddDrawable(id, &drawable)
	w.logic.AddTickable(id, &tickable)
	w.engine.AddMoveable(id, &moveable)

	return id
}

func (w *World) RemoveBow(id int) {
	delete(w.Bows, id)
	w.drawer.RemoveDrawable(id)
	w.logic.RemoveTickable(id)
	w.engine.RemoveMoveable(id)
}

func (w *World) SpawnArrow(h *bow.Holder, charge float64) {
	var shooter = (*h).(arrow.Shooter)
	var a = arrow.NewArrow(&shooter, charge)

	w.AddArrow(a)
}
