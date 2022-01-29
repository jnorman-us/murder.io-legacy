package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/engine"
	"github.com/josephnormandev/murder/common/entities/bow"
)

func (w *World) AddBow(b *bow.Bow) int {
	var id = b.GetID()

	var spawner = bow.Spawner(w)
	b.SetSpawner(&spawner)

	var drawable = drawer.Drawable(b)
	var moveable = engine.Moveable(b)

	w.drawer.AddDrawable(id, &drawable)
	w.engine.AddMoveable(id, &moveable)
	w.Bows[id] = b
	return id
}

func (w *World) RemoveBow(id int) {
	delete(w.Bows, id)
	w.drawer.RemoveDrawable(id)
	w.engine.RemoveMoveable(id)
}

func (w *World) SpawnArrow(h *bow.Holder, charge float64) {
	// do nothing, this is the client!
}
