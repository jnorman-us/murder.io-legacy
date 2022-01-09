package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/arrow"
	"github.com/josephnormandev/murder/common/packet"
)

func (w *World) AddArrow(a *arrow.Arrow) int {
	var id int

	var spawner = arrow.Spawner(w)
	a.SetSpawner(&spawner)

	if w.environment.IsServer() {
		id = w.NextAvailableID()
		a.SetID(id)

		var spawn = packet.Spawn(a)
		var moveable = engine.Moveable(a)
		var arrowWall = collisions.ArrowWall(a)
		var arrowPlayer = collisions.ArrowPlayer(a)

		w.network.AddSpawn(id, &spawn)
		w.engine.AddMoveable(id, &moveable)
		w.collisions.AddArrowWall(id, &arrowWall)
		w.collisions.AddArrowPlayer(id, &arrowPlayer)
	} else {
		id = a.GetID()

		var moveable = engine.Moveable(a)
		var drawable = drawer.Drawable(a)

		w.engine.AddMoveable(id, &moveable)
		w.drawer.AddDrawable(id, &drawable)
	}
	w.Arrows[id] = a
	return id
}

func (w *World) RemoveArrow(id int) {
	delete(w.Arrows, id)
	if w.environment.IsServer() {
		w.network.RemoveSpawn(id)
		w.engine.RemoveMoveable(id)
		w.collisions.RemoveArrowWall(id)
		w.collisions.RemoveArrowPlayer(id)
	} else {
		w.drawer.RemoveDrawable(id)
		w.engine.RemoveMoveable(id)
	}
}

func (w *World) RemoveArrowCollidable(id int) {
	w.collisions.RemoveArrowWall(id)
	w.collisions.RemoveArrowPlayer(id)
}
