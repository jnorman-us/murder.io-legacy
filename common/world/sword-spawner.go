package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/sword"
	"github.com/josephnormandev/murder/common/logic"
)

func (w *World) AddSword(s *sword.Sword) int {
	var id = w.NextAvailableID()
	var tickable = logic.Tickable(s)
	var drawable = drawer.Drawable(s)
	var moveable = engine.Moveable(s)
	var collidable = collisions.Collidable(s)

	s.SetID(id)

	w.Swords[id] = s
	w.drawer.AddDrawable(id, &drawable)
	w.logic.AddTickable(id, &tickable)
	w.engine.AddMoveable(id, &moveable)
	w.collisions.AddCollidable(id, &collidable)

	return id
}

func (w *World) RemoveSword(id int) {
	delete(w.Swords, id)
	w.drawer.RemoveDrawable(id)
	w.logic.RemoveTickable(id)
	w.engine.RemoveMoveable(id)
	w.collisions.RemoveCollidable(id)
}
