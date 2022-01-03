package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/bow"
	"github.com/josephnormandev/murder/common/logic"
)

func (w *World) AddBow(b *bow.Bow) int {
	var id = w.NextAvailableID()
	var tickable = logic.Tickable(b)
	var drawable = drawer.Drawable(b)
	var moveable = engine.Moveable(b)
	var collidable = collisions.Collidable(b)

	b.SetID(id)

	w.Bows[id] = b
	w.drawer.AddDrawable(id, &drawable)
	w.logic.AddTickable(id, &tickable)
	w.engine.AddMoveable(id, &moveable)
	w.collisions.AddCollidable(id, &collidable)

	return id
}

func (w *World) RemoveBow(id int) {
	delete(w.Bows, id)
	w.drawer.RemoveDrawable(id)
	w.logic.RemoveTickable(id)
	w.engine.RemoveMoveable(id)
	w.collisions.RemoveCollidable(id)
}
