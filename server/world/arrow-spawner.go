package world

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/arrow"
	"github.com/josephnormandev/murder/server/ws"
)

func (w *World) AddArrow(a *arrow.Arrow) int {
	var id = w.NextAvailableID()
	a.SetID(id)

	var spawner = arrow.Spawner(w)
	a.SetSpawner(&spawner)

	var spawn = ws.Spawn(a)
	var moveable = engine.Moveable(a)
	var arrowWall = collisions.ArrowWall(a)
	var arrowPlayer = collisions.ArrowPlayer(a)

	w.network.AddSpawn(id, &spawn)
	w.engine.AddMoveable(id, &moveable)
	w.collisions.AddArrowWall(id, &arrowWall)
	w.collisions.AddArrowPlayer(id, &arrowPlayer)
	w.Arrows[id] = a
	return id
}

func (w *World) RemoveArrow(id int) {
	delete(w.Arrows, id)
	w.network.RemoveSpawn(id)
	w.engine.RemoveMoveable(id)
	w.collisions.RemoveArrowWall(id)
	w.collisions.RemoveArrowPlayer(id)
	w.Deletions.RemoveID(id)
}

func (w *World) RemoveArrowCollidable(id int) {
	w.collisions.RemoveArrowWall(id)
	w.collisions.RemoveArrowPlayer(id)
}
