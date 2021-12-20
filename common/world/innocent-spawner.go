package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/innocent"
)

func (w *World) AddInnocent(i *innocent.Innocent) int {
	var id = w.NextAvailableID()
	var drawable = drawer.Drawable(i)
	var moveable = engine.Moveable(i)
	var collidable = collisions.Collidable(i)

	i.SetID(id)
	w.Innocents[id] = i
	w.drawer.AddDrawable(id, &drawable)
	w.engine.AddMoveable(id, &moveable)
	w.collisions.AddCollidable(id, &collidable)

	return id
}

func (w *World) RemoveInnocent(id int) {
	delete(w.Innocents, id)
	w.drawer.RemoveDrawable(id)
	w.engine.RemoveMoveable(id)
	w.collisions.RemoveCollidable(id)
}

func (w *World) SpawnSword(i *innocent.Innocent) *innocent.Swingable {
	return nil
}
