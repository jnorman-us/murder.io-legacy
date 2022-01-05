package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/entities/wall"
)

func (w *World) AddWall(wl *wall.Wall) int {
	var id = w.NextAvailableID()
	var drawable = drawer.Drawable(wl)
	var wallArrow = collisions.WallArrow(wl)
	var wallPlayer = collisions.WallPlayer(wl)

	wl.SetID(id)

	w.Walls[id] = wl
	w.drawer.AddDrawable(id, &drawable)
	w.collisions.AddWallArrow(id, &wallArrow)
	w.collisions.AddWallPlayer(id, &wallPlayer)

	return id
}

func (w *World) RemoveWall(id int) {
	delete(w.Walls, id)
	w.drawer.RemoveDrawable(id)
	w.collisions.RemoveWallArrow(id)
	w.collisions.RemoveWallPlayer(id)
}
