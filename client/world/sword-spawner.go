package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/engine"
	"github.com/josephnormandev/murder/common/legacy-entities/sword"
)

func (w *World) AddSword(s *sword.Sword) int {
	var id = s.GetID()

	var drawable = drawer.Drawable(s)
	var moveable = engine.Moveable(s)

	w.drawer.AddDrawable(id, &drawable)
	w.engine.AddMoveable(id, &moveable)
	w.Swords[id] = s
	return id
}

func (w *World) RemoveSword(id int) {
	delete(w.Swords, id)
	w.drawer.RemoveDrawable(id)
	w.engine.RemoveMoveable(id)
}
