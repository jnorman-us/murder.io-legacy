package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/arrow"
)

func (w *World) AddArrow(a *arrow.Arrow) int {
	var id = w.NextAvailableID()
	var spawner = arrow.Spawner(w)
	var drawable = drawer.Drawable(a)
	var moveable = engine.Moveable(a)
	var arrowWall = collisions.ArrowWall(a)
	var arrowPlayer = collisions.ArrowPlayer(a)

	a.SetID(id)
	a.SetSpawner(&spawner)

	w.Arrows[id] = a
	w.drawer.AddDrawable(id, &drawable)
	w.engine.AddMoveable(id, &moveable)
	w.collisions.AddArrowWall(id, &arrowWall)
	w.collisions.AddArrowPlayer(id, &arrowPlayer)

	return id
}

func (w *World) RemoveArrow(id int) {
	delete(w.Arrows, id)
	w.drawer.RemoveDrawable(id)
	w.logic.RemoveTickable(id)
	w.engine.RemoveMoveable(id)
	w.collisions.RemoveArrowWall(id)
	w.collisions.RemoveArrowPlayer(id)
}

func (w *World) RemoveArrowCollidable(id int) {
	w.collisions.RemoveArrowWall(id)
	w.collisions.RemoveArrowPlayer(id)
}
