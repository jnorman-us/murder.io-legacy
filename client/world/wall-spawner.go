package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/entities/wall"
)

func (w *World) AddWall(wl *wall.Wall) int {
	var id = wl.GetID()

	var drawable = drawer.Drawable(wl)

	w.drawer.AddDrawable(id, &drawable)
	w.Walls[id] = wl
	return id
}

func (w *World) RemoveWall(id int) {
	delete(w.Walls, id)
	w.drawer.RemoveDrawable(id)
}
