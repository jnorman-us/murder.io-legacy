package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/engine"
	"github.com/josephnormandev/murder/common/legacy-entities/arrow"
)

func (w *World) AddArrow(a *arrow.Arrow) int {
	var id = a.GetID()

	var spawner = arrow.Spawner(w)
	a.SetSpawner(&spawner)

	var moveable = engine.Moveable(a)
	var drawable = drawer.Drawable(a)

	w.engine.AddMoveable(id, &moveable)
	w.drawer.AddDrawable(id, &drawable)
	w.Arrows[id] = a
	return id
}

func (w *World) RemoveArrow(id int) {
	delete(w.Arrows, id)
	w.drawer.RemoveDrawable(id)
	w.engine.RemoveMoveable(id)
}

func (w *World) RemoveArrowCollidable(id int) {
	return
}
