package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/arrow"
	"github.com/josephnormandev/murder/common/logic"
)

func (w *World) AddArrow(a *arrow.Arrow) int {
	var id = w.NextAvailableID()
	var tickable = logic.Tickable(a)
	var drawable = drawer.Drawable(a)
	var moveable = engine.Moveable(a)
	var collidable = collisions.Collidable(a)

	a.SetID(id)

	w.Arrows[id] = a
	w.drawer.AddDrawable(id, &drawable)
	w.logic.AddTickable(id, &tickable)
	w.engine.AddMoveable(id, &moveable)
	w.collisions.AddCollidable(id, &collidable)

	return id
}

func (w *World) RemoveArrow(id int) {
	delete(w.Arrows, id)
	w.drawer.RemoveDrawable(id)
	w.logic.RemoveTickable(id)
	w.engine.RemoveMoveable(id)
	w.collisions.RemoveCollidable(id)
}
